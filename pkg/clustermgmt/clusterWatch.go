/*
IBM Confidential
OCO Source Materials
5737-E67
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/

package clustermgmt

import (
	"strings"
	"time"

	"github.com/golang/glog"
	mcm "github.ibm.com/IBMPrivateCloud/hcm-api/pkg/apis/mcm/v1alpha1"
	mcmClientset "github.ibm.com/IBMPrivateCloud/hcm-api/pkg/client/clientset_generated/clientset"
	"github.ibm.com/IBMPrivateCloud/search-aggregator/pkg/config"
	db "github.ibm.com/IBMPrivateCloud/search-aggregator/pkg/dbconnector"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	clusterregistry "k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
	clientset "k8s.io/cluster-registry/pkg/client/clientset/versioned"
	informers "k8s.io/cluster-registry/pkg/client/informers/externalversions"
)

// WatchClusters watches k8s cluster and clusterstatus objects and updates the search cache.
func WatchClusters() {
	var clientConfig *rest.Config
	var err error

	if config.Cfg.KubeConfig != "" {
		glog.Infof("Creating k8s client using path: %s", config.Cfg.KubeConfig)
		clientConfig, err = clientcmd.BuildConfigFromFlags("", config.Cfg.KubeConfig)
	} else {
		glog.Info("Creating k8s client using InClusterConfig()")
		clientConfig, err = rest.InClusterConfig()
	}

	if err != nil {
		glog.Fatal("Error Constructing Client From Config: ", err)
	}

	stopper := make(chan struct{})
	defer close(stopper)

	// Initialize the mcm client, used for ClusterStatus resource
	mcmClient, err := mcmClientset.NewForConfig(clientConfig)
	if err != nil {
		glog.Fatal("Cannot Construct MCM Client From Config: ", err)
	}

	// Initialize the cluster client, used for Cluster resource
	clusterClient, err := clientset.NewForConfig(clientConfig)
	if err != nil {
		glog.Fatal("Cannot Construct Cluster Client From Config: ", err)
	}
	clusterFactory := informers.NewSharedInformerFactory(clusterClient, 0)
	clusterInformer := clusterFactory.Clusterregistry().V1alpha1().Clusters().Informer()
	clusterInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			cluster, ok := obj.(*clusterregistry.Cluster)
			if !ok {
				glog.Error("Failed to assert Cluster informer obj to clusterregistry.Cluster")
				return
			}

			clusterStatus, err := mcmClient.McmV1alpha1().
				ClusterStatuses(cluster.GetNamespace()).Get(cluster.GetName(), v1.GetOptions{})
			if err != nil {
				glog.Error("Failed to fetch cluster resource: ", err)
			}

			resource := transformCluster(cluster, clusterStatus)

			glog.Info("Inserting Cluster resource in RedisGraph. ", resource)

			_, _, err = db.Insert([]*db.Resource{&resource})
			if err != nil {
				glog.Error("Error adding Cluster kind with error: ", err)
			}
		},
		UpdateFunc: func(prev interface{}, next interface{}) {
			cluster, ok := next.(*clusterregistry.Cluster)
			if !ok {
				glog.Error("Failed to assert Cluster informer obj to clusterregistry.Cluster")
				return
			}

			clusterStatus, err := mcmClient.McmV1alpha1().
				ClusterStatuses(cluster.GetNamespace()).Get(cluster.GetName(), v1.GetOptions{})
			if err != nil {
				glog.Error("Failed to fetch cluster resource: ", err)
			}

			resource := transformCluster(cluster, clusterStatus)

			glog.Info("Updating Cluster resource in RedisGraph. ", resource)
			_, _, err = db.Update([]*db.Resource{&resource})
			if err != nil {
				glog.Error("Error updating Cluster kind with errors: ", err)
				// If the key is missing from redis we should try to insert it again
				if isGraphMissing(err) {
					glog.Info("Attempting to recreate Cluster graph object")
					_, _, err = db.Insert([]*db.Resource{&resource})
					if err != nil {
						glog.Error("Error adding Cluster kind with error: ", err)
					}
				}
			}

			// If a cluster is offline we should remove the cluster objects
			if resource.Properties["status"] == "offline" {
				_, badNameErr, err := db.DeleteCluster(cluster.GetName())
				if badNameErr != nil {
					glog.Error("Invalid Cluster Name: ", cluster.GetName())
				}
				if err != nil {
					glog.Error("Error deleting current resources for cluster: ", err)
				}
			}
		},
		DeleteFunc: func(obj interface{}) {
			cluster, ok := obj.(*clusterregistry.Cluster)
			if !ok {
				glog.Error("Failed to assert Cluster informer obj to clusterregistry.Cluster")
				return
			}

			glog.Info("Deleting Cluster resource in RedisGraph.")
			uid := string(cluster.GetUID())
			_, err = db.Delete([]string{uid})
			if err != nil {
				glog.Error("Error deleting Cluster kind with error: ", err)
			}

			// When a cluster (ClusterStatus) gets deleted, we must remove all resources for that cluster from RedisGraph.
			_, badNameErr, err := db.DeleteCluster(cluster.GetName())
			if badNameErr != nil {
				glog.Error("Invalid Cluster Name: ", cluster.GetName())
			}
			if err != nil {
				glog.Error("Error deleting current resources for cluster: ", err)
			}
		},
	})

	clusterInformer.Run(stopper)
}

// Test for specific redis graph update error
func isGraphMissing(err error) bool {
	return strings.Contains(err.Error(), "key doesn't contains a graph object")
}

func transformCluster(cluster *clusterregistry.Cluster, clusterStatus *mcm.ClusterStatus) db.Resource {
	props := make(map[string]interface{})

	props["name"] = cluster.GetName()
	props["kind"] = "Cluster"
	props["cluster"] = "local-cluster" // Needed for rbac
	props["selfLink"] = cluster.GetSelfLink()
	props["created"] = cluster.GetCreationTimestamp().UTC().Format(time.RFC3339)

	if cluster.GetLabels() != nil {
		props["label"] = clusterStatus.GetLabels()
	}
	if cluster.GetNamespace() != "" {
		props["namespace"] = clusterStatus.GetNamespace()
	}

	// we are pulling the status from the cluster object and cluster info from the clusterStatus object :(
	if len(cluster.Status.Conditions) > 0 && cluster.Status.Conditions[0].Type != "" {
		props["status"] = string(cluster.Status.Conditions[0].Type)
	} else {
		props["status"] = "offline"
	}

	if clusterStatus != nil {
		props["consoleURL"] = clusterStatus.Spec.ConsoleURL
		props["cpu"], _ = clusterStatus.Spec.Capacity.Cpu().AsInt64()
		props["memory"] = clusterStatus.Spec.Capacity.Memory().String()
		props["klusterletVersion"] = clusterStatus.Spec.KlusterletVersion
		props["kubernetesVersion"] = clusterStatus.Spec.Version

		props["nodes"] = int64(0)
		nodes, ok := clusterStatus.Spec.Capacity["nodes"]
		if ok {
			props["nodes"], _ = nodes.AsInt64()
		}

		props["storage"] = ""
		storage, ok := clusterStatus.Spec.Capacity["storage"]
		if ok {
			props["storage"] = storage.String()
		}
	}

	return db.Resource{
		Kind:       "Cluster",
		UID:        string(cluster.GetUID()),
		Properties: props,
	}
}

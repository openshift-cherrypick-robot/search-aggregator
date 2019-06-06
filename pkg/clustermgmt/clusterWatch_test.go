/*
IBM Confidential
OCO Source Materials
5737-E67
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/

package clustermgmt

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mcmapi "github.ibm.com/IBMPrivateCloud/hcm-api/pkg/apis/mcm/v1alpha1"
	utils "github.ibm.com/IBMPrivateCloud/search-aggregator/pkg/utils"
	clusterregistry "k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
)

func TestTransformCluster(t *testing.T) {
	testcluster := clusterregistry.Cluster{}
	testclusterstatus := mcmapi.ClusterStatus{}
	utils.UnmarshalFile("../../test-data/cluster.json", &testcluster, t)
	utils.UnmarshalFile("../../test-data/clusterstatus.json", &testclusterstatus, t)
	result := transformCluster(&testcluster, &testclusterstatus)
	assert.Equal(t, result.Kind, "Cluster", "Test Kind")
	assert.Equal(t, result.ResourceString, "clusters", "Test ResourceString")
	assert.Equal(t, result.UID, "1baa5f8a-758f-11e9-9527-667a72062d69", "Test UID")
	assert.Equal(t, result.Properties["name"], "xav-cluster", "Test Name")
	assert.Equal(t, result.Properties["namespace"], "xav-cluster-ns", "Test namespace")
	assert.Equal(t, (result.Properties["label"]).(map[string]string)["cloud"], "IBM", "Test label")
	assert.Equal(t, result.Properties["status"], "OK", "Test status")
	assert.Equal(t, result.Properties["created"], "2019-05-13T14:55:11Z", "Test created")
	assert.Equal(t, result.Properties["consoleURL"], "https://222.222.222.222:8443", "Test consoleURL")
	assert.Equal(t, result.Properties["cpu"], int64(24), "Test cpu")
	assert.Equal(t, result.Properties["memory"], "98143Mi", "Test memory")
	assert.Equal(t, result.Properties["nodes"], int64(3), "Test nodes")
	assert.Equal(t, result.Properties["storage"], "60Gi", "Test storage")

}
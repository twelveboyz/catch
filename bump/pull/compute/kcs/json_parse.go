package kcs

import (
	"Catch/bump/internal/utils"
	"github.com/tidwall/gjson"
	"sync"
)

func ResourceInfoALL(body string) []*KubernetesCloudServer {

	var sliceKubernetesCloudServer []*KubernetesCloudServer
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.cluster|@pretty")
	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var kubernetesCloudServer KubernetesCloudServer
		kubernetesCloudServer.ClusterId = value.Get("clusterId").String()
		kubernetesCloudServer.ClusterName = value.Get("clusterName").String()
		kubernetesCloudServer.Pool = value.Get("pool").String()
		kubernetesCloudServer.Region = value.Get("region").String()
		kubernetesCloudServer.NodeQuota = value.Get("nodeQuota").String()
		kubernetesCloudServer.KubeVersion = value.Get("kubeVersion").String()
		kubernetesCloudServer.ClusterRuntime = value.Get("clusterRuntime").String()
		kubernetesCloudServer.ClusterRuntimeVersion = value.Get("clusterRuntimeVersion").String()
		kubernetesCloudServer.VpcName = value.Get("vpcName").String()
		kubernetesCloudServer.Os = value.Get("os").String()
		kubernetesCloudServer.PeriodType = value.Get("periodType").String()
		sliceKubernetesCloudServer = append(sliceKubernetesCloudServer, &kubernetesCloudServer)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceKubernetesCloudServer
}

func GetNodeResourceInfo(kcss []*KubernetesCloudServer) []*KubernetesCloudServer {
	var wg sync.WaitGroup
	wg.Add(len(kcss))
	for _, kcs := range kcss {
		utils.GoPool.Go(func() {
			defer wg.Done()
			url := GetNodeInfoUrl(kcs.ClusterId)
			jsonBody := utils.Get(url, Headers)
			nodes := gjson.Get(jsonBody, "body.nodes")
			nodes.ForEach(func(key, value gjson.Result) bool {
				var node Nodes
				node.ServerId = value.Get("serverId").String()
				node.NodeId = value.Get("nodeID").String()
				node.Name = value.Get("name").String()
				node.Cpu = value.Get("cpu").String()
				node.Memory = value.Get("memory").String()
				node.OsImage = value.Get("osImage").String()
				node.Os = value.Get("os").String()
				node.ServerType = value.Get("serverType").String()
				node.VmType = value.Get("vmType").String()
				node.SpecsName = value.Get("specsName").String()
				node.InstanceId = value.Get("instanceId").String()

				kcs.Nodes = append(kcs.Nodes, node)
				return true
			})
		})

	}

	wg.Wait()
	return kcss
}

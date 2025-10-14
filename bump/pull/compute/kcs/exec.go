package kcs

import (
	"Catch/bump/internal/utils"
	"fmt"
)

func ExecuteALL() {
	Body := utils.Get(Url, Headers)
	//fmt.Println(cloudPortBody)

	infoSet := ResourceInfoALL(Body)

	infoSet = GetNodeResourceInfo(infoSet)

	for _, info := range infoSet {
		fmt.Println(info)
		for _, node := range info.Nodes {
			fmt.Printf("name=%s,nodeId=%s,os=%s,spec=%s,cpu=%s,memory=%s\n", node.Name, node.NodeId, node.Os, node.SpecsName, node.Cpu, node.Memory)
		}
	}

}

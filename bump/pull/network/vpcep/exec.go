package vpcep

import (
	"Catch/bump/internal/utils"
	"fmt"
)

func ExecuteALL() {
	Body := utils.Get(Url, Headers)
	//fmt.Println(cloudPortBody)

	InfoSet := ResourceInfoALL(Body)
	for _, info := range InfoSet {
		fmt.Println(info)
	}
}

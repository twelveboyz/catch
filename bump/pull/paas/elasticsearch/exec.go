package elasticsearch

import (
	"Catch/bump/internal/utils"
	"bytes"
	"fmt"
)

func ExecuteALL() {
	Body := utils.Post(Url, Headers, bytes.NewBufferString(Payload))
	//fmt.Println(Body)

	InfoSet := ResourceInfoALL(Body)
	for _, info := range InfoSet {
		fmt.Println(info.Id, info.Alias, info.ZoneZoneName, info.AppType)

		for _, s := range info.Servers {
			fmt.Println(s)
		}
		fmt.Println()
	}
}

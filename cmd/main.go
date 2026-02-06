package main

import (
	"Catch/cmdb"
	"Catch/config"
	"Catch/daily"
	"Catch/internal"
	"Catch/resourceCount"
)

func main() {
	//tip
	internal.Tip()

	for {
		choose := internal.InputChoose()
		switch choose {
		case "1":
			cmdb.Run()
		case "2":
			daily.Run()
		case "3":
			config.Run()
		case "4":
			resourceCount.Run()
		}
	}
}

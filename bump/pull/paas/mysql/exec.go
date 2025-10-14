package mysql

import (
	"Catch/bump/internal/utils"
	"fmt"
)

func ExecuteALL() {
	/*--------------------------------------mysql--------------------------------------*/
	mysqlBody := utils.Get(Url, Headers)

	mysqlInfoSet := ResourceInfoALL(mysqlBody)
	for _, v := range mysqlInfoSet {
		fmt.Println(v)
	}
}

func Execute() {
	/*--------------------------------------mysql--------------------------------------*/
	mysqlBody := utils.Get(Url, Headers)

	var testData = []string{"c292-gz3-szmp-mysql-01", "c135-gz3-imai-mysql-39f23", "c130_szyg_mysql_001"}
	mysqlInfoSet := ResourceInfo(mysqlBody, testData)
	for _, v := range mysqlInfoSet {
		fmt.Println(v)
	}
}

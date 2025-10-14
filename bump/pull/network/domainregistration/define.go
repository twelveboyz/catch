package domainregistration

import (
	"Catch/bump/internal/bootstrap"
	"fmt"
)

type DomainRegistration struct {
	Id              string
	DomainName      string
	DomainLifeCycle string
	DomainOwner     string
	Days            string
}

var Url = "https://console.ecloud.10086.cn/api/web/domains/v1/service/console/domain/queryAll?page=1&size=10"

var Headers = map[string]string{
	"Accept":          "application/json, text/plain, */*",
	"Accept-Language": "zh-CN",
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
	"Host":            "console.ecloud.10086.cn",
}

var Payload string = `{"domainLifeCycle":["AUTH_PERIOD","HANDLE_PERIOD","USEFUL_PERIOD","RENEWAL_PERIOD","REDEEM_PERIOD","AUTH_FAIL","DOMAIN_INFO_UPDATING","REDEEM_IN","TRANSFER_IN_PERIOD","TRANSFER_OUT_PERIOD","COMMIT"],"sortColumn":"domain_due_time","sortType":"ASC","domainName":""}`

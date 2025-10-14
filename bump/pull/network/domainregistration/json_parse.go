package domainregistration

import "github.com/tidwall/gjson"

func ResourceInfoALL(body string) []DomainRegistration {

	var sliceDomainRegistration []DomainRegistration
	//获取json路径下的所有内容|@pertty 是格式化输出json
	JsonBody := gjson.Get(body, "body.content|@pretty")

	//fmt.Println(JsonBody) //排错使用，打印获取到的json原数据内容
	//循环content下的所有元素
	JsonBody.ForEach(func(key, value gjson.Result) bool {
		var domainRegistration DomainRegistration
		domainRegistration.Id = value.Get("id").String()
		domainRegistration.DomainName = value.Get("domainName").String()
		domainRegistration.DomainLifeCycle = value.Get("domainLifeCycle").String()
		domainRegistration.DomainOwner = value.Get("domainOwner").String()
		domainRegistration.Days = value.Get("days").String()

		sliceDomainRegistration = append(sliceDomainRegistration, domainRegistration)
		//fmt.Println("test", cloudPort)
		return true
	})

	return sliceDomainRegistration
}

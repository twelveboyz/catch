package utils

import (
	"Catch/bump/internal/bootstrap"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func HeadersNoPool() map[string]string {
	return map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN",
		"Connection":      "keep-alive",
		"Content-Type":    "application/json",
		"Cookie":          fmt.Sprintf("CMECLOUDTOKEN=%s", bootstrap.Token),
		"Host":            "console.ecloud.10086.cn",
	}
}

func DetermineTokenValid() error {
	url := "https://ecloud.10086.cn/api/web/portalcenter/mopWeb/getUserInfoOutline"
	requests, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("创建请求失败:%s", err.Error()))
	}

	for k, v := range HeadersNoPool() {
		requests.Header.Set(k, v)
	}

	c := &http.Client{
		//Transport: transport,
		Timeout: 60 * time.Second,
	}

	response, err := c.Do(requests)
	if err != nil {
		return errors.New(fmt.Sprintf("发起请求失败:%s", err.Error()))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("读取body失败:%s", err.Error()))
	}

	if strings.Contains(string(body), `"state":"OK"`) {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Invalid Token!!!"))
	}

}

func Get(url string, header map[string]string) string {
	for i := 0; i < 3; i++ {
		requests, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("创建请求失败", err)
		}

		for k, v := range header {
			requests.Header.Set(k, v)
		}

		c := &http.Client{
			//Transport: transport,
			Timeout: 60 * time.Second,
		}

		response, err := c.Do(requests)
		if err != nil {
			logrus.Errorln("发起请求失败", err)
			time.Sleep(5 * time.Second)
			continue
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			logrus.Errorln("读取body失败", err)
		}

		bodyStr := string(body)

		statusCode := gjson.Get(bodyStr, "errorCode").String()
		if statusCode != "" {
			logrus.Errorln("http返回信息有误：", response.StatusCode, bodyStr)
		}

		if err = response.Body.Close(); err != nil {
			logrus.Errorln("关闭response.body失败", err)
		}

		/*	total := gjson.Get(bodyStr, "body.total").Int()
			log.Println("total:", total)*/

		//fmt.Println(bodyStr)

		return bodyStr
	}
	return ""
}

func Post(url string, header map[string]string, payload io.Reader) string {
	requests, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("创建请求失败", err)
	}

	for k, v := range header {
		requests.Header.Set(k, v)
	}

	c := &http.Client{}
	response, err := c.Do(requests)
	if err != nil {
		fmt.Println("发起请求失败", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取body失败", err)
	}

	bodyStr := string(body)

	statusCode := gjson.Get(bodyStr, "errorCode").String()
	if statusCode != "" {
		log.Println("http返回信息有误：", response.StatusCode, bodyStr)
	}

	//fmt.Println(bodyStr)

	return bodyStr
}

func GetTotal(url string, header map[string]string) int64 {
	requests, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败", err)
	}

	for k, v := range header {
		requests.Header.Set(k, v)
	}

	c := &http.Client{}
	response, err := c.Do(requests)
	if err != nil {
		fmt.Println("发起请求失败", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取body失败", err)
	}

	bodyStr := string(body)

	statusCode := gjson.Get(bodyStr, "errorCode").String()
	if statusCode != "" {
		log.Println("http返回信息有误：", response.StatusCode, bodyStr)
	}

	total := gjson.Get(bodyStr, "body.total").Int()
	//log.Println("total:", total)

	//fmt.Println(bodyStr)

	return total
}

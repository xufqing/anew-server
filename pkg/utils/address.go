package utils

import (
	"anew-server/pkg/common"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IpResp struct {
	Status   string `json:"status"`
	Province string `json:"province"`
	City     string `json:"city"`
}

// 去高德地图申请账号并生成Key
const apiKey string = "58129d7b9b628e8d26978b5714687d69"

// 获取IP真实地址
func GetIpRealLocation(ip string) string {
	address := "未知地址"
	if ip == "127.0.0.1" {
		address = "本地网络"
	} else {
		resp, err := http.Get(fmt.Sprintf("https://restapi.amap.com/v3/ip?ip=%s&key=%s", ip, apiKey))
		if err != nil {
			common.Log.Error(fmt.Sprintf("[GetIpRealLocation]IP地址查询失败: %v", err))
			return address
		}
		defer resp.Body.Close()
		// 读取响应数据
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			common.Log.Error(fmt.Sprintf("[GetIpRealLocation]IP地址查询失败: %v", err))
			return address
		}
		// json数据转结构体
		var result IpResp
		Json2Struct(string(data), &result)
		if result.Status == "1" {
			address = result.Province
			// 城市不为空且城市与省份不重复
			if result.City != "" && result.Province != result.City {
				address += result.City
			}
		}
	}
	return address
}

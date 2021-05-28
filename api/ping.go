package api

import (
	response2 "anew-server/api/response"
	"anew-server/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 检查服务器是否通畅
func Ping(c *gin.Context) {
	response2.Success()
}

func ShowUserAvatar(c *gin.Context) {
	path := c.Query("path")
	// 读取头像
	if path != "" {
		path = strings.Replace(path,"..","",-1) //  防止任意文件读取漏洞
		fmt.Println(path)
		imgPath := common.Conf.Upload.SaveDir + "/avatar/" + path
		c.File(imgPath)
	}
}
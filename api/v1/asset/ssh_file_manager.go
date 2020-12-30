package asset

import (
	"anew-server/dto/response"
	"anew-server/models"
	"anew-server/pkg/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
)

// 获取目录数据
func GetPathFromSsh(c *gin.Context) {
	urlPath := c.Query("path")
	key := c.Query("key")
	if key == ""  {
		response.FailWithCode(response.ParmError)
		return
	}
	conn, err := hub.get(key)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取Ssh连接失败: %s", err))
		return
	}
	if conn.SshClient == nil {
		response.FailWithMsg("主机未连接")
		return
	}
	if urlPath == "" {
		session, _ := conn.SshClient.NewSession()
		res, _ := session.CombinedOutput("echo ${HOME}")
		urlPath = strings.Replace(utils.Bytes2Str(res), "\n", "", -1)
		defer session.Close()
	}
	lsInfo, err := conn.SFTPClient.ReadDir(urlPath)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取文件信息错误：%s", err))
		return
	}
	var files = make([]response.FileInfo, 0)
	for i := range lsInfo {
		file := response.FileInfo{
			Name:  lsInfo[i].Name(),
			Path:  path.Join(urlPath, lsInfo[i].Name()),
			IsDir: lsInfo[i].IsDir(),
			Size:  utils.FormatFileSize(lsInfo[i].Size()),
			Mtime: models.LocalTime{
				Time: lsInfo[i].ModTime(),
			},
			Mode:   lsInfo[i].Mode().String(),
			IsLink: lsInfo[i].Mode()&os.ModeSymlink == os.ModeSymlink,
		}
		files = append(files, file)
	}
	response.SuccessWithData(files)
}

func UploadFileToSsh(c *gin.Context) {
	urlPath := c.Query("path")
	key := c.Query("key")
	if key == "" || urlPath == "" {
		response.FailWithCode(response.ParmError)
		return
	}
	// 读取文件
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMsg("无法读取文件")
		return
	}
	filename := file.Filename
	remoteFile := path.Join(urlPath, filename)
	conn, err := hub.get(key)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取Ssh连接失败: %s", err))
		return
	}
	dstFile, err := conn.SFTPClient.Create(remoteFile)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("sftp创建流失败：%s", err))
	}
	defer dstFile.Close()
	// 将文件流传到sftp
	src, err := file.Open()
	if err != nil {
		response.FailWithMsg("打开文件失败")
		return
	}
	buf := make([]byte, 1024)
	for {
		n, _ := src.Read(buf)
		if n == 0 {
			break
		}
		_, _ = dstFile.Write(buf)
	}
	response.Success()
}

func DownloadFileFromSsh(c *gin.Context) {
	urlPath := c.Query("path")
	key := c.Query("key")
	if key == "" || urlPath == "" {
		response.FailWithCode(response.ParmError)
		return
	}
	conn, err := hub.get(key)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取Ssh连接失败: %s", err))
		return
	}
	dstFile, err := conn.SFTPClient.Open(urlPath)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("sftp打开文件失败：%s", err))
	}
	defer dstFile.Close()
	var buff bytes.Buffer
	if _, err := dstFile.WriteTo(&buff); err != nil {
		response.FailWithMsg(fmt.Sprintf("写入文件流失败：%s", err))
	}
	c.Header("content-disposition", `attachment; filename=`+urlPath)
	c.Data(http.StatusOK, "application/octet-stream", buff.Bytes())
}


// 删除文件
func DeleteFileInSsh(c *gin.Context) {
	urlPath := c.Query("path")
	key := c.Query("key")
	if key == "" || urlPath == "" {
		response.FailWithCode(response.ParmError)
		return
	}
	conn, err := hub.get(key)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取Ssh连接失败: %s", err))
		return
	}
	if err := conn.SFTPClient.Remove(urlPath); err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
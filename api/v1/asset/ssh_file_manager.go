package asset

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/pkg/sshx"
	"anew-server/pkg/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"net/http"
	"os"
	"path"
	"strings"
)

func GetSSHByHostID(id uint) (*ssh.Client, error) {
	s := service.New()
	host, err := s.GetHostById(id)
	if err != nil {
		return nil, err
	}
	var conf *ssh.ClientConfig
	switch host.AuthType {
	case "key":
		conf, err = sshx.NewAuthConfig(host.User, "", host.PrivateKey, host.KeyPassphrase)
		if err != nil {
			return nil, err
		}
	default:
		// 默认密码模式
		conf, err = sshx.NewAuthConfig(host.User, host.Password, "", "")
		if err != nil {
			return nil, err
		}
	}
	addr := fmt.Sprintf("%s:%s", host.IpAddress, host.Port)
	client, err := ssh.Dial("tcp", addr, conf)
	return client, err
}

// 获取目录数据
func GetPathFromSSH(c *gin.Context) {
	var req request.FileReq
	err := c.Bind(&req)
	if err != nil || req.HostId == 0 {
		response.FailWithCode(response.ParmError)
		return
	}
	client, err := GetSSHByHostID(req.HostId)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取SSH客户端失败：%s", err))
	}
	if req.Path == "home" {
		session, _ := client.NewSession()
		res, _ := session.CombinedOutput("echo ${HOME}")
		req.Path = strings.Replace(utils.Bytes2Str(res), "\n", "", -1)
		defer session.Close()
	}

	defer client.Close()
	sftpx, err := sftp.NewClient(client)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取sftp客户端错误：%s", err))
		return
	}
	lsInfo, err := sftpx.ReadDir(req.Path)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取文件信息错误：%s", err))
		return
	}
	var files = make([]response.FileInfo, 0)
	for i := range lsInfo {
		file := response.FileInfo{
			Name:  lsInfo[i].Name(),
			Path:  path.Join(req.Path, lsInfo[i].Name()),
			IsDir: lsInfo[i].IsDir(),
			Size:  lsInfo[i].Size(),
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

func UploadFileToSSH(c *gin.Context) {
	hostId := utils.Str2Uint(c.Query("host_id"))
	urlPath := c.Query("path")
	if hostId == 0||urlPath == "" {
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
	client,err := GetSSHByHostID(hostId)
	if err !=nil {
		response.FailWithMsg(fmt.Sprintf("获取SSH客户端失败: %s",err))
	}
	sftpx, err := sftp.NewClient(client)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取sftp客户端错误：%s", err))
		return
	}
	dstFile, err := sftpx.Create(remoteFile)
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

func DownloadFileFromSSH(c *gin.Context) {
	hostId := utils.Str2Uint(c.Query("host_id"))
	urlPath := c.Query("path")
	if hostId == 0||urlPath == "" {
		response.FailWithCode(response.ParmError)
		return
	}
	client,err := GetSSHByHostID(hostId)
	if err !=nil {
		response.FailWithMsg(fmt.Sprintf("获取SSH客户端失败: %s",err))
	}
	sftpx, err := sftp.NewClient(client)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取sftp客户端错误：%s", err))
		return
	}
	dstFile, err := sftpx.Open(urlPath)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("sftp打开文件失败：%s", err))
	}
	defer dstFile.Close()
	var buff bytes.Buffer
	if _, err := dstFile.WriteTo(&buff); err != nil {
		response.FailWithMsg(fmt.Sprintf("写入文件流失败：%s", err))
	}
	c.Header("content-disposition", `attachment; filename=` + urlPath)
	c.Data(http.StatusOK, "application/octet-stream", buff.Bytes())
}
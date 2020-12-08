package sshx

import (
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)
// code by mylxsw  from https://github.com/mylxsw/sshx/

// 认证配置
type AuthConfig struct {
	User                 string
	Password             string
	PrivateKeyPath       string
	PrivateKeyPassphrase string
}

type SSHClient struct {
	Host   string
	Conf   *ssh.ClientConfig
	Client *ssh.Client //ssh客户端
	Logger *zap.SugaredLogger
}
type SessionOption func(session *ssh.Session) error
// ErrRemoteFileExisted 远程服务器已经存在该文件
var ErrRemoteFileExisted = errors.New("remote file already exist")
// ErrSessionCanceled 会话因为上下文对象的取消而被取消
var ErrSessionCanceled = errors.New("session canceled because context canceled")
// ErrFileFingerNotMatch 文件指纹不匹配
var ErrFileFingerNotMatch = errors.New("file finger not match")

func New(host string, conf *ssh.ClientConfig) *SSHClient {
	return &SSHClient{Host: host, Conf: conf, Logger: common.Log}
}


//func main {
//	conf, _ := NewAuthConfig("root", "123123", "", "")
//	ssha := New("192.168.1.1:22", conf)
//}

func NewAuthConfig(user string, password string, privateKeyPath string, privateKeyPassphrase string) (*ssh.ClientConfig, error) {
	conf := ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}
	if password != "" {
		conf.Auth = append(conf.Auth, ssh.Password(password))
	} else if privateKeyPath != "" {
		pk, err := getPrivateKey(privateKeyPath, privateKeyPassphrase)
		if err != nil {
			return nil, err
		}
		conf.Auth = append(conf.Auth, pk)
	} else {
		// if occur error "Failed to open SSH_AUTH_SOCK: dial unix: missing address",
		// execute command: eval `ssh-agent`,and enter passphrase
		conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			return nil, fmt.Errorf("failed to open SSH_AUTH_SOCK: %w", err)
		}
		agentClient := agent.NewClient(conn)
		// Use a callback rather than PublicKeys so we only consult the
		// agent once the remote server wants it.
		conf.Auth = append(conf.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	return &conf, nil
}

func getPrivateKey(privateKeyPath string, privateKeyPassphrase string) (ssh.AuthMethod, error) {
	if !utils.FileExist(privateKeyPath) {
		privateKeyPath = filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa")
	}
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}
	var signer ssh.Signer
	if privateKeyPassphrase != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(privateKeyPassphrase))
	} else {
		signer, err = ssh.ParsePrivateKey(key)
	}
	if err != nil {
		return nil, fmt.Errorf("parse private key failed: %v", err)
	}
	return ssh.PublicKeys(signer), nil
}

func (s *SSHClient) run(ctx context.Context, client *ssh.Client, cmd string, opts ...SessionOption) ([]byte, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("create session failed: %w", err)
	}
	defer session.Close()

	for _, opt := range opts {
		if err := opt(session); err != nil {
			return nil, err
		}
	}

	var resp []byte
	stopped := make(chan interface{}, 0)
	go func() {
		resp, err = session.CombinedOutput(cmd)
		stopped <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		_ = session.Signal(ssh.SIGKILL)
		err = ErrSessionCanceled
	case <-stopped:
	}

	return resp, err
}

func (s *SSHClient) checkFileConsistency(src string, remoteDest string) (bool, error) {
	// md5sum默认路径
	md5bin := "/usr/bin/md5sum"
	localFinger := utils.GetFileMd5(src)
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	remoteFingerBytes, err := s.Command(ctx, fmt.Sprintf("%s %s", md5bin, remoteDest))
	if err != nil {
		s.Logger.Errorf("check file md5sum for remote file failed: %w", err)
		return false, fmt.Errorf("check file md5sum for remote file failed: %w", err)
	}

	remoteFinger := strings.SplitN(strings.TrimSpace(string(remoteFingerBytes)), " ", 2)
	// debug 日志开启
	if common.Conf.Logs.Level == -1 {
		s.Logger.Debugf("Sftp文件一致性检查: local=%s, remote=%s, matched=%v", localFinger, remoteFinger[0], strings.EqualFold(localFinger, remoteFinger[0]))
	}

	return strings.EqualFold(localFinger, remoteFinger[0]), nil
}

func (s *SSHClient) transferToRemoteTmp(client *sftp.Client, destTmp string, src string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		s.Logger.Errorf("open local file %s failed: %w", src, err)
		return 0, fmt.Errorf("open local file %s failed: %w", src, err)
	}
	defer srcFile.Close()

	destFile, err := client.Create(destTmp)
	if err != nil {
		s.Logger.Errorf("create remote temp file %s failed: %w", destTmp, err)
		return 0, fmt.Errorf("create remote temp file %s failed: %w", destTmp, err)
	}
	defer destFile.Close()

	return io.Copy(destFile, srcFile)
}

func (s *SSHClient) remoteFileExist(client *sftp.Client, path string) bool {
	_, err := client.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func (s *SSHClient) transferFile(client *ssh.Client, dest string, src string, override bool, checkConsistency bool) (written int64, err error) {
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		s.Logger.Errorf("creates SFTP client failed: %w", err)
		return 0, fmt.Errorf("creates SFTP client failed: %w", err)
	}
	defer sftpClient.Close()

	if !override && s.remoteFileExist(sftpClient, dest) {
		if checkConsistency {
			matched, err := s.checkFileConsistency(src, dest)
			if err != nil {
				s.Logger.Errorf("check local & remote file (existed) consistency failed: %w", err)
				return 0, fmt.Errorf("check local & remote file (existed) consistency failed: %w", err)
			}

			if !matched {
				return 0, ErrFileFingerNotMatch
			}
		}

		return 0, ErrRemoteFileExisted
	}

	destTmp := filepath.Join(filepath.Dir(dest), fmt.Sprintf("%s.tmp_%d", filepath.Base(dest), time.Now().UnixNano()))
	written, err = s.transferToRemoteTmp(sftpClient, destTmp, src)
	if err != nil {
		s.Logger.Errorf("transfer local file to remote failed: %w", err)
		return 0, fmt.Errorf("transfer local file to remote failed: %w", err)
	}
	defer sftpClient.Remove(destTmp)

	if err := sftpClient.PosixRename(destTmp, dest); err != nil {
		return 0, err
	}

	if checkConsistency {
		matched, err := s.checkFileConsistency(src, dest)
		if err != nil {
			s.Logger.Errorf("check local & remote file consistency failed: %w", err)
			return 0, fmt.Errorf("check local & remote file consistency failed: %w", err)
		}

		if !matched {
			return 0, ErrFileFingerNotMatch
		}
	}

	return written, nil
}

// Command 在远程服务器上执行命令
func (s *SSHClient) Command(ctx context.Context, cmd string, opts ...SessionOption) ([]byte, error) {
	client, err := ssh.Dial("tcp", s.Host, s.Conf)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return s.run(ctx, client, cmd, opts...)
}

// SendFile 将本地文件传输到远程服务器
func (s *SSHClient) SendFile(dest string, src string, override bool, consistencyCheck bool) (written int64, err error) {
	conn, err := ssh.Dial("tcp", s.Host, s.Conf)
	if err != nil {
		s.Logger.Errorf("can not establish a connection to %s: %w", s.Host, err)
		return 0, fmt.Errorf("can not establish a connection to %s: %w", s.Host, err)
	}
	defer conn.Close()

	return s.transferFile(conn, dest, src, override, consistencyCheck)
}

package asciicast2

import (
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"encoding/json"
	"os"
	"time"
)

const V2OutputEvent = "o"
const V2InputEvent = "i"

type CastV2Header struct {
	Version      uint               `json:"version"`
	Width        int               `json:"width"`
	Height       int               `json:"height"`
	Timestamp    int64              `json:"timestamp,omitempty"`
	Duration     float64            `json:"duration,omitempty"`
	Title        string             `json:"title,omitempty"`
	Command      string             `json:"command,omitempty"`
	Env          *map[string]string `json:"env,omitempty"`
	outputStream *json.Encoder
}

func NewCastV2(meta CastV2Header,file string) (*CastV2Header, *os.File) {
	var c CastV2Header
	f,_ :=os.OpenFile(common.Conf.SSh.RecordDir + "/" + file,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	c.Version = 2
	c.Width = meta.Width
	c.Height = meta.Height
	c.Title = meta.Title
	c.Timestamp = meta.Timestamp
	c.Duration = c.Duration
	c.Env = meta.Env
	c.outputStream = json.NewEncoder(f)
	c.outputStream.Encode(c)
	return &c, f
}

func (c *CastV2Header) Record(t time.Time, data []byte) {
	out := make([]interface{}, 3)
	timeNow := time.Since(t).Seconds()
	out[0] = timeNow
	out[1] = V2OutputEvent
	out[2] = utils.Bytes2Str(data)
	c.Duration = timeNow // 写回结构体,暂时不知道干嘛用
	c.outputStream.Encode(out)
}

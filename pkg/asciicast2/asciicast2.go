package asciicast2

import (
	"anew-server/pkg/utils"
	"encoding/json"
	"io"
	"time"
)

const V2OutputEvent = "o"
const V2InputEvent = "i"

type CastV2Header struct {
	Version      uint               `json:"version"`
	Width        uint               `json:"width"`
	Height       uint               `json:"height"`
	Timestamp    int64              `json:"timestamp,omitempty"`
	Duration     float64            `json:"duration,omitempty"`
	Title        string             `json:"title,omitempty"`
	Command      string             `json:"command,omitempty"`
	Env          *map[string]string `json:"env,omitempty"`
	outputStream *json.Encoder
}

func NewCastV2(meta CastV2Header, fd io.Writer) *CastV2Header {
	var c CastV2Header
	c.Version = 2
	c.Width = meta.Width
	c.Height = meta.Height
	c.Title = meta.Title
	c.Timestamp = meta.Timestamp
	c.Duration = c.Duration
	c.Env = meta.Env
	c.outputStream = json.NewEncoder(fd)
	c.outputStream.Encode(c)
	return &c
}

func (c *CastV2Header) PushFrame(t time.Time, data []byte) {
	out := make([]interface{}, 3)
	timeNow := time.Since(t).Seconds()
	out[0] = timeNow
	out[1] = V2OutputEvent
	out[2] = utils.Bytes2Str(data)
	c.Duration = timeNow // 写回结构体
	c.outputStream.Encode(out)
}

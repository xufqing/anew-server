package models

import (
	"anew-server/common"
	"database/sql/driver"
	"fmt"
	"time"
)

// 由于gorm提供的base model没有json tag, 使用自定义
type Model struct {
	Id        uint       `gorm:"primary_key;comment:'自增编号'" json:"id"`
	CreatedTime LocalTime  `gorm:"comment:'创建时间'" json:"createdTime"`
	UpdatedTime LocalTime  `gorm:"comment:'更新时间'" json:"updatedTime"`
}

// 表名设置
func (Model) TableName(name string) string {
	// 添加表前缀
	return fmt.Sprintf("%s%s", common.Conf.Mysql.TablePrefix, name)
}

// 本地时间
type LocalTime struct {
	time.Time
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// ""空值不进行解析
	if len(data) == 2 {
		*t = LocalTime{Time: time.Time{}}
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+common.SecLocalTimeFormat+`"`, string(data))
	*t = LocalTime{Time: now}
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(common.SecLocalTimeFormat))
	return []byte(output), nil
}

// gorm 写入 mysql 时调用
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// gorm 检出 mysql 时调用
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to LocalTime", v)
}

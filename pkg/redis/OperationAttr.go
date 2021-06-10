package redis

import (
	"fmt"
	"time"
)

const (
	ATTR_EXPR = "expr" // 过期时间
	ATTR_NX   = "nx"   // redis锁
	ATTR_XX   = "xx"   // redis xx锁
)

type empty struct {
}
type OperationAttr struct {
	Name  string
	Value interface{}
}
type OperationAttrs []*OperationAttr

func (this OperationAttrs) Find(name string) *InterfaceResult {
	for _, attr := range this {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("OperationAttr found error:%s", name))
}

func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{ATTR_EXPR, t}
}
func WithNX() *OperationAttr {
	return &OperationAttr{ATTR_NX, empty{}}
}
func WithXX() *OperationAttr {
	return &OperationAttr{ATTR_XX, empty{}}
}

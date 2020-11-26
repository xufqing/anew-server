package redis

type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}
func (this *Iterator) HasNext() bool {
	if this.data == nil || len(this.data) == 0 {
		return false
	}
	return this.index < len(this.data)
}

func (this *Iterator) Next() (ret interface{}) {
	ret = this.data[this.index]
	this.index = this.index + 1
	return
}

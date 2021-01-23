package convert

import "strconv"

//类型转换
type StrTo string

//如果一个结构体实现了String() 方法,那么fmt.Println() 默认会调用String 方法进行输出;
func (s StrTo) String() string {
	return string(s)
}
func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}
func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}
func (s StrTo) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}

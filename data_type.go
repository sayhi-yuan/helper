package helper

// 数据类型

type UnsignedInteger interface {
	~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64
}

type Integer interface {
	UnsignedInteger | ~int8 | ~int16 | ~int | ~int32 | ~int64
}

type StringInteger interface {
	Integer | ~string
}

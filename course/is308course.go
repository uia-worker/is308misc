package is308course

import "fmt"

type C interface {
	F1(name string)
	F2(name string) (error, float32)
	F3() int64
}

type T int64

func (T) F1(name string) {
	fmt.Println(name)
}

func (T) F2(name string) (error, float32) {
	return nil, 10.2
}

func (T) F3() int64 {
	return 10
}

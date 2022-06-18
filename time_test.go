package study

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t1 := time.Second
	now := time.Now()
	fmt.Println(now.UnixNano())
	fmt.Println(time.Now().Add(t1).UnixNano())
	fmt.Println(time.Now().Add(t1 * time.Millisecond).UnixNano())
}

//type Locker interface {
//	lock()
//}

type Lock struct {
}

func (l Lock) lock() {
	fmt.Println("lock ..")
}

type Ca struct {
	name string
	l    Lock
}

func TestCa(t *testing.T) {
	nc := new(*Ca)

	fmt.Println(nc)
}

type myInt int64

func (i myInt) add(a, b int64) int64 {
	return a + b
}

func TestMyInt(t *testing.T) {
	var mi myInt
	fmt.Println(&mi)
	fmt.Println(mi.add(1, 2))
}

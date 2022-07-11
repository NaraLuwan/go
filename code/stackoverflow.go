package main

import (
	"fmt"
)

type people struct {
	Name string
}

func (p *people) std() {
	fmt.Printf("print: %v", p)
}

func (p *people) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &people{}
	p.std()
}

/* 运行报错：
/private/var/folders/mt/_gyb9cjn2mn389mz8rrq1f5r0000gn/T/GoLand/___go_build_stackoverflow_go
runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0xc020160378 stack=[0xc020160000, 0xc040160000]
fatal error: stack overflow

runtime stack:
runtime.throw(0x10ca5f0, 0xe)
/usr/local/Cellar/go@1.16/1.16.7/libexec/src/runtime/panic.go:1117 +0x72

原因是自定义的类有string方法，fmt的stringer接口也有个string方法，自定义类间接的实现了stringer接口，产生了递归。

解决方法：
方法1. 换个string这个函数名，这个跟fmt的stringer interface冲突了
方法2. 不要直接把类传进去，而是把属性传递进去。
*/

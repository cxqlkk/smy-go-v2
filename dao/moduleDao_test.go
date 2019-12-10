package dao

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestString(t *testing.T){
	bts:=[]byte{'a','b','c'}
	str:=*(*string)(unsafe.Pointer(&bts))
	fmt.Println(str)
}

func TestModuleDao_List(t *testing.T) {

	d:=NewModuleDao(MasterEngin())
	r:=d.ListByName("")
	fmt.Println(r)
}


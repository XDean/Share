package main

import (
	"fmt"
	"github.com/xdean/goex/xconfig"
	"testing"
)

func TestGen(t *testing.T) {
	fmt.Println(xconfig.EncryptString("my-password", "123456"))
}

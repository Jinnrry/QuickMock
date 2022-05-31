package services

import (
	"github.com/Jinnrry/gop"
	"github.com/robertkrimen/otto"
	"testing"
)

func TestGetApiInfo(t *testing.T) {
	ret, _ := GetApiInfo("/user/info")
	gop.Print(ret)
}

func TestJS(t *testing.T) {
	vm := otto.New()
	vm.Run(`
    a=1;
	b=2;
	c=3;
`)

	v, e := vm.Get("b")
	gop.Print(v)
	gop.Print(e)

}

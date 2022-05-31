package services

import (
	"fmt"
	"github.com/Jinnrry/gop"
	"testing"
)

func TestUpsertApi(t *testing.T) {
	err := UpsertApi("user/get", &ApiInfo{
		HttpHeader: map[string]any{
			"Content-Type": "json; charset=UTF-8",
		},
		HttpCode:       200,
		Response:       []byte("{\"success\":1}"),
		Method:         "GET|POST",
		CustomFunction: "",
		Wait:           100,
	})

	fmt.Println(err)
}

func TestGetApiList(t *testing.T) {
	res, err := GetApiList()

	gop.Print(err)
	gop.Print(res)
}

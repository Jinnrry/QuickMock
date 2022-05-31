package controllers

import (
	"QuickMock/services"
	"QuickMock/tools"
	"QuickMock/views"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"net/http"
)

type response struct {
	Success  int    `json:"success"`
	Message  string `json:"msg"`
	ServerIp string `json:"server_ip"`
	Data     any    `json:"data"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, views.IndexHtml)
}

func Jquery(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, views.JqueryJS)
}

func Vue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, views.VueJS)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, views.EditHtml)
	}
}

func DelApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json; charset=utf-8")
	r.PostFormValue("")

	path := r.PostForm.Get("path")

	err := services.DeleteApi(path)
	ret := response{}
	if err != nil {
		ret.Success = 500
		ret.Message = err.Error()
	}

	retByte, _ := json.Marshal(ret)
	fmt.Fprint(w, string(retByte))

}

func UpsertApi(w http.ResponseWriter, r *http.Request) {
	r.PostFormValue("")
	apiInfo := &services.ApiInfo{
		Response:       []byte(r.PostFormValue("response")),
		Method:         r.PostFormValue("methods"),
		CustomFunction: r.PostFormValue("script"),
		HttpCode:       cast.ToInt(r.PostFormValue("httpCode")),
		Wait:           cast.ToInt(r.PostFormValue("wait")),
	}

	responseHeaderStr := r.PostFormValue("responseHead")
	_ = json.Unmarshal([]byte(responseHeaderStr), &apiInfo.HttpHeader)

	err := services.UpsertApi(r.PostFormValue("apiPath"), apiInfo)
	res := response{}
	if err != nil {
		res.Success = 500
		res.Message = err.Error()
	}

	ret, _ := json.Marshal(res)

	fmt.Fprint(w, string(ret))
}

func GetApiList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json; charset=utf-8")

	list, err := services.GetApiList()
	ret := response{}
	if err != nil {
		ret.Success = 500
	} else {
		ret.ServerIp = tools.GetLocalIp()
		ret.Success = 0
		ret.Data = list
	}
	retByte, _ := json.Marshal(ret)
	fmt.Fprint(w, string(retByte))
}

type apiInfoResponse struct {
	HttpHeader     map[string]any `json:"header"`
	Response       string         `json:"response"`
	Method         string         `json:"method"`
	CustomFunction string         `json:"custom_function"`
	HttpCode       int            `json:"http_code"`
	Wait           int            `json:"wait"`
}

func GetApiDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json; charset=utf-8")
	UrlData := r.URL.Query()

	apiPath := cast.ToString(UrlData.Get("api"))

	apiInfo, err := services.GetApiInfo(apiPath)
	apiInfoRes := apiInfoResponse{
		HttpHeader:     apiInfo.HttpHeader,
		Response:       string(apiInfo.Response),
		Method:         apiInfo.Method,
		CustomFunction: apiInfo.CustomFunction,
		HttpCode:       apiInfo.HttpCode,
		Wait:           apiInfo.Wait,
	}

	ret := response{}
	if err != nil {
		ret.Success = 500
	} else {
		ret.Success = 0
		ret.Data = apiInfoRes
	}
	retByte, _ := json.Marshal(ret)
	fmt.Fprint(w, string(retByte))

}

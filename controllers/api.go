package controllers

import (
	"QuickMock/services"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AllPathHandler(w http.ResponseWriter, r *http.Request) {
	// 避免postForm为空
	r.PostFormValue("")

	requestPath := r.URL.Path
	UrlData := r.URL.Query()
	BodyData, _ := ioutil.ReadAll(r.Body)
	PostForm := r.PostForm
	reqHeader := r.Header

	apiInfo, err := services.GetApiInfo(requestPath)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	services.BuildApiResponse(w, r.Method, apiInfo, UrlData, PostForm, reqHeader, BodyData)

}

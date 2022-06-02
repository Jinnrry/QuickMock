package services

import (
	"QuickMock/config"
	"QuickMock/tools"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type ApiInfo struct {
	HttpHeader     map[string]any `json:"header"`
	Response       []byte         `json:"-"`
	Method         string         `json:"method"`
	CustomFunction string         `json:"custom_function"`
	HttpCode       int            `json:"http_code"`
	Wait           int            `json:"wait"`
}

func GetApiInfo(apiPath string) (*ApiInfo, error) {
	ret := &ApiInfo{}
	// 加载设置文件
	dataPath := strings.Replace(apiPath, "/", config.Separator, -1)
	filePtr, err := os.Open(config.RunConfig.DataPath + dataPath + "/settings.json")
	if err != nil {
		return nil, errors.New("接口不存在")
	}
	defer filePtr.Close()

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&ret)
	if err != nil {
		return nil, errors.New("接口配置文件损坏")
	}
	// 加载response文件
	filePtr, err = os.Open(config.RunConfig.DataPath + dataPath + "/response")
	if err != nil {
		return nil, errors.New("接口配置文件损坏")
	}
	ret.Response, _ = ioutil.ReadAll(filePtr)

	return ret, nil
}

func BuildApiResponse(w http.ResponseWriter, reqMethod string, apiInfo *ApiInfo, urlData url.Values, postForm url.Values, reqHeader http.Header, reqBody []byte) {

	if apiInfo.CustomFunction != "" {
		vm := otto.New()
		jsObjUrlData, _ := json.Marshal(urlData)
		jsObjPostForm, _ := json.Marshal(postForm)
		jsObjReqHeader, _ := json.Marshal(reqHeader)
		jsObjReqBody, _ := json.Marshal(reqBody)

		jsObjResponseHeader, _ := json.Marshal(apiInfo.HttpHeader)

		initScript := fmt.Sprintf("var UrlData = %s;"+
			"var PostForm = %s;"+
			"var RequestHeader = %s;"+
			"var RequestBody = %s;"+
			"var Wait = %d;"+
			"var ResponseHeader = %s;"+
			"var HttpCode = %d;"+
			"var Response = %s;", jsObjUrlData, jsObjPostForm, jsObjReqHeader, jsObjReqBody, apiInfo.Wait, jsObjResponseHeader, apiInfo.HttpCode, apiInfo.Response)

		vm.Run(initScript)
		vm.Run(apiInfo.CustomFunction)
		vm.Run(`
			var ResponseHeader = JSON.stringify(ResponseHeader);
			var Response = JSON.stringify(Response);
		`)

		// 重写等待时间
		v, err := vm.Get("Wait")
		if err == nil {
			tmp, _ := v.ToInteger()
			apiInfo.Wait = cast.ToInt(tmp)
		} else {
			fmt.Println(err.Error())
			return
		}

		// 重写响应头
		v, err = vm.Get("ResponseHeader")
		if err == nil {
			tmp, _ := v.ToString()
			err = json.Unmarshal([]byte(tmp), &apiInfo.HttpHeader)
			if err != nil {
				fmt.Println("自定义脚本ResponseHeader值错误")
			}
		} else {
			fmt.Println(err.Error())
			return
		}

		// 重写状态码
		v, err = vm.Get("HttpCode")
		if err == nil {
			tmp, _ := v.ToInteger()
			apiInfo.HttpCode = cast.ToInt(tmp)
		} else {
			fmt.Println(err.Error())
			return
		}

		// 重写响应数据
		v, err = vm.Get("Response")
		if err == nil {
			tmp, _ := v.ToString()
			apiInfo.Response = []byte(tmp)
		} else {
			fmt.Println(err.Error())
			return
		}

	}

	if !tools.InArray(reqMethod, strings.Split(apiInfo.Method, "|")) {
		w.WriteHeader(403)
		fmt.Fprint(w, "接口不支持该方式请求")
		return
	}

	if apiInfo.Wait != 0 {
		time.Sleep(time.Duration(apiInfo.Wait) * time.Millisecond)
	}

	for name, value := range apiInfo.HttpHeader {
		w.Header().Set(name, cast.ToString(value))
	}
	w.WriteHeader(apiInfo.HttpCode)
	w.Write(apiInfo.Response)
}

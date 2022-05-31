package services

import (
	"QuickMock/config"
	"QuickMock/tools"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// UpsertApi 更新或者新增api
func UpsertApi(apiPath string, info *ApiInfo) error {

	if apiPath[0:1] != "/" {
		apiPath = "/" + apiPath
	}

	dataPath := strings.Replace(apiPath, "/", "|||", -1)

	_ = os.MkdirAll(config.RunConfig.DataPath+dataPath, 0777)

	apiSettings, _ := json.Marshal(info)
	err := ioutil.WriteFile(config.RunConfig.DataPath+dataPath+"/settings.json", apiSettings, 0777)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(config.RunConfig.DataPath+dataPath+"/response", info.Response, 0777)

	return err

}

type apiItem struct {
	ApiPath    string `json:"api_path"`
	RequestUrl string `json:"request_url"`
}

// GetApiList 获取api列表
func GetApiList() ([]apiItem, error) {
	var api []apiItem

	dirList, err := ioutil.ReadDir(config.RunConfig.DataPath)
	if err != nil {
		return nil, err
	}
	for _, file := range dirList {
		apiPath := strings.Replace(file.Name(), "|||", "/", -1)

		if config.RunConfig.ApiServerPort != 80 {
			api = append(api, apiItem{
				ApiPath:    apiPath,
				RequestUrl: fmt.Sprintf("http://%s:%d%s", tools.GetLocalIp(), config.RunConfig.ApiServerPort, apiPath),
			})
		} else {
			api = append(api, apiItem{
				ApiPath:    apiPath,
				RequestUrl: fmt.Sprintf("http://%s%s", tools.GetLocalIp(), apiPath),
			})
		}

	}
	return api, nil
}

func DeleteApi(path string) error {
	dataPath := strings.Replace(path, "/", "|||", -1)
	dirPath := config.RunConfig.DataPath + dataPath
	return os.RemoveAll(dirPath)
}

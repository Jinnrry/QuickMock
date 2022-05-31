package main

import (
	"QuickMock/config"
	"QuickMock/controllers"
	"QuickMock/tools"
	"fmt"
	"net/http"
	"os"
)

func startApiServer() {

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/", controllers.AllPathHandler)
	apiServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.RunConfig.ApiServerPort),
		Handler: apiMux,
	}

	go func() {
		err := apiServer.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}()
}

func startAdminServer() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/", controllers.Home)
	adminMux.HandleFunc("/apiList", controllers.GetApiList)
	adminMux.HandleFunc("/apiDetail", controllers.GetApiDetail)
	adminMux.HandleFunc("/upsert", controllers.UpsertApi)
	adminMux.HandleFunc("/delete", controllers.DelApi)
	adminMux.HandleFunc("/edit.html", controllers.Edit)
	adminMux.HandleFunc("/jquery3.js", controllers.Jquery)
	adminMux.HandleFunc("/vue2.js", controllers.Vue)

	adminServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.RunConfig.AdminServerPort),
		Handler: adminMux,
	}

	go func() {
		err := adminServer.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}()

}

func main() {
	startApiServer()
	startAdminServer()

	fmt.Printf("启动成功！api接口运行端口%d\r\n"+
		"admin后台运行端口%d\r\n"+
		"你可以通过http://127.0.0.1:%d 或者 http://%s:%d 进行接口管理",
		config.RunConfig.ApiServerPort, config.RunConfig.AdminServerPort, config.RunConfig.AdminServerPort, tools.GetLocalIp(), config.RunConfig.AdminServerPort)
	select {}
}

package config

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"strings"
)

const Separator = "```"

type Config struct {
	ApiServerPort   int
	AdminServerPort int
	DataPath        string
}

var RunConfig = Config{
	AdminServerPort: 8080,
	ApiServerPort:   80,
	DataPath:        "./run_data/",
}

func init() {
	if len(os.Args) > 1 {
		if os.Args[1] == "help" ||
			os.Args[1] == "-h" ||
			os.Args[1] == "--h" ||
			os.Args[1] == "--help" ||
			os.Args[1] == "-help" {
			fmt.Printf("支持命令如下：\r\n " +
				"1、help 显示帮助信息\r\n " +
				"2、api_port 指定api端口，示例：api_port=80 \r\n " +
				"3、admin_port 指定管理后台端口，示例：admin_port=8080\r\n" +
				"4、data_path 指定数据存储路径，示例:data_path=./run_data/\r\n")

			os.Exit(0)
		}

		for i := 1; i < len(os.Args); i++ {
			datas := strings.Split(os.Args[i], "=")
			if len(datas) == 2 {

				switch datas[0] {
				case "api_port":
					RunConfig.ApiServerPort = cast.ToInt(datas[1])
				case "admin_port":
					RunConfig.AdminServerPort = cast.ToInt(datas[1])
				case "data_path":
					RunConfig.DataPath = datas[1]
					if RunConfig.DataPath[len(RunConfig.DataPath)-1:] != "/" {
						RunConfig.DataPath += "/"
					}

				}

			}

		}

	}
}

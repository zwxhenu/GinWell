package core

import (
	"GinWell-Server/config"
	"GinWell-Server/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func Viper(path ...string) *viper.Viper {
	var ConfigInfo string
	if len(path) == 0 {
		flag.StringVar(&ConfigInfo, "c", "", "choose config file.")
		flag.Parse()
		if ConfigInfo == "" {
			if configEnv := os.Getenv(config.ConfigEnv); configEnv == "" {
				ConfigInfo = config.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config.ConfigFile)
			} else {
				ConfigInfo = configEnv
				fmt.Printf("您正在使用GW_CONFIG环境变量,config的路径为%v\n", ConfigInfo)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", ConfigInfo)
		}
	} else {
		ConfigInfo = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", ConfigInfo)
	}

	v := viper.New()
	v.SetConfigFile(ConfigInfo)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GW_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GW_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}

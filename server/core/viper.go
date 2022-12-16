package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/weilaim/wm-admin/server/core/internal"
	"github.com/weilaim/wm-admin/server/global"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			// 判断internal.ConfigEnv 常量存储的环境变量是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称，config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称，config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称，config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)

				}
			} else {
				// internal.ConfigEnv 常量存储的环境变量不为空 将赋值于config
				config = configEnv
			}
		} else {
			// 命令行参数不为空将赋值于config
			fmt.Printf("您正在使用命令行的-c参数进行传递的值，config的路径为%s\n", internal.ConfigEnv, config)
		}
	} else {
		// 函数传递的可变参数的第一个赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值，config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error ocnfig file:%s\n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.WM_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.WM_CONFIG); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root 位置去找到对应迁移位置，保证root路径有效
	global.WM_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

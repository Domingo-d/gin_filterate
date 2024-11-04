package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"server/global"
)

func Viper(confName string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(confName)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s \n", err.Error()))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}

	return v
}

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//初始化
var AppConfig Config

func init()  {
	var ViperConfig = viper.New()
	InitConfig(ViperConfig)
	if err := ViperConfig.Unmarshal(&AppConfig); err != nil {
		log.Printf("初始话配置文件失败！")
	}
	fmt.Println(AppConfig)
}

func InitConfig(viperConfig *viper.Viper)  {
	viperConfig.AddConfigPath("./")          //设置读取的文件路径
	viperConfig.SetConfigName("application") //设置读取的文件名
	viperConfig.SetConfigType("yaml")        //设置文件的类型
	//读取配置
	if err := viperConfig.ReadInConfig(); err != nil {
		panic(err)
	}
}

type Config struct {
	Redis Redis
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

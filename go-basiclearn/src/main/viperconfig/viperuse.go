package viperconfig

import (
	"fmt"
	"github.com/spf13/viper"
	_ "gopkg.in/yaml.v2"
	"log"
)

//初始化
var AppConfig Config

func Main()  {
	var ViperConfig = viper.New()
	InitConfig(ViperConfig)
	//1、读取yaml中配置项转为map
	c := ViperConfig.AllSettings()
	//marshal, err := yaml.Unmarshal(c)
	//AppConfig = marshal
	fmt.Println(c) //map结构
	//2、序列化yaml转struct
	if err := ViperConfig.Unmarshal(&AppConfig); err != nil {
		log.Printf("初始话配置文件失败！")
	}
	fmt.Println(AppConfig)
	//拼接
	//1、数据库连接url
	dsn := AppConfig.DataSource.UserName + ":" + AppConfig.DataSource.Password + "@" + AppConfig.DataSource.Url
	fmt.Println(dsn)
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
	Server Server
	DataSource DataSource
	Redis Redis
}

type Server struct {
	Port string `mapstructure:"port"`
}

type DataSource struct {
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url string `mapstructure:"url"`
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

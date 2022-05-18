package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	*ServerConfig `mapstructure:"app"`
	*LogConfig    `mapstructure:"log"`
	*RedisConfig  `mapstructure:"redis"`
	*MysqlConfig  `mapstructure:"mysql"`
}

type ServerConfig struct {
	Name        string `mapstructure:"name"`
	Mode        string `mapstructure:"mode"`
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	Description string `mapstructure:"description"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Database           string `mapstructure:"database"`
	MaxConnections     int    `mapstructure:"max_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
}

type RedisConfig struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	Password           string `mapstructure:"password"`
	Database           int    `mapstructure:"database"`
	MaxConnections     int    `mapstructure:"max_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
	PoolSize           int    `mapstructure:"pool_size"`
}

func Init(filename string) (err error) {
	// 方式1 通过用户指定配置文件
	viper.SetConfigName(filename)
	// 方式2 指定配置文件名和配置文件的位置,viper 自行查找可用配置文件
	//viper.SetConfigFile("config.yaml")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		// Handle errors reading the config file
		fmt.Printf("Fatal error config file: %s \n\n", err)
		return err
	}
	// 配置信息反序列化到Conf变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err: %s \n\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err: %s \n\n", err)
			return
		}
	})

	return nil

}

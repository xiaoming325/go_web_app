package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// 初始化viper，读取配置文件并监视更改
func Init() error {
	viper.SetConfigName("config")  // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")    // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./conf/") // 查找配置文件所在的路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		//读取配置信息失败
		panic(fmt.Errorf("ReadInConfig failed, err: %v", err))
	}

	err = viper.Unmarshal(&Conf) // 将配置文件信息反序列化到Conf结构体中
	if err != nil {
		panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
	}

	//监视配置文件的更改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了...")
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(fmt.Errorf("unmarshal to Conf failed, err:%v", err))
		}
	})
	return err
}

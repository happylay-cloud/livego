package configure

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/kr/pretty"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
{
  "server": [
    {
      "appname": "live",
      "live": true,
	  "hls": true,
	  "static_push": []
    }
  ]
}
*/

type Application struct {
	Appname    string   `mapstructure:"appname"`
	Live       bool     `mapstructure:"live"`
	Hls        bool     `mapstructure:"hls"`
	Flv        bool     `mapstructure:"flv"`
	Api        bool     `mapstructure:"api"`
	StaticPush []string `mapstructure:"static_push"`
}

type Applications []Application

type JWT struct {
	Secret    string `mapstructure:"secret"`
	Algorithm string `mapstructure:"algorithm"`
}
type ServerCfg struct {
	Level           string       `mapstructure:"level"`
	ConfigFile      string       `mapstructure:"config_file"`
	FLVArchive      bool         `mapstructure:"flv_archive"`
	FLVDir          string       `mapstructure:"flv_dir"`
	RTMPNoAuth      bool         `mapstructure:"rtmp_noauth"`
	RTMPAddr        string       `mapstructure:"rtmp_addr"`
	HTTPFLVAddr     string       `mapstructure:"httpflv_addr"`
	HLSAddr         string       `mapstructure:"hls_addr"`
	HLSKeepAfterEnd bool         `mapstructure:"hls_keep_after_end"`
	APIAddr         string       `mapstructure:"api_addr"`
	RedisAddr       string       `mapstructure:"redis_addr"`
	RedisPwd        string       `mapstructure:"redis_pwd"`
	ReadTimeout     int          `mapstructure:"read_timeout"`
	WriteTimeout    int          `mapstructure:"write_timeout"`
	GopNum          int          `mapstructure:"gop_num"`
	JWT             JWT          `mapstructure:"jwt"`
	Server          Applications `mapstructure:"server"`
}

// default config
var defaultConf = ServerCfg{
	ConfigFile:      "config/config.yaml",
	FLVArchive:      false,
	RTMPNoAuth:      false,
	RTMPAddr:        ":1935",
	HTTPFLVAddr:     ":7001",
	HLSAddr:         ":7002",
	HLSKeepAfterEnd: false,
	APIAddr:         ":8090",
	WriteTimeout:    10,
	ReadTimeout:     10,
	GopNum:          1,
	Server: Applications{{
		Appname:    "live",
		Live:       true,
		Hls:        true,
		Flv:        true,
		Api:        true,
		StaticPush: nil,
	}},
}

// 初始化配置文件对象
var Config = viper.New()

// 初始化日志
func initLog() {
	if l, err := log.ParseLevel(Config.GetString("level")); err == nil {
		log.SetLevel(l)
		log.SetReportCaller(l == log.DebugLevel)
	}
}

func init() {
	defer Init()

	// 默认配置
	b, _ := json.Marshal(defaultConf)
	defaultConfig := bytes.NewReader(b)
	viper.SetConfigType("json")
	_ = viper.ReadConfig(defaultConfig)
	_ = Config.MergeConfigMap(viper.AllSettings())

	// 命令行标志
	pflag.String("rtmp_addr", ":1935", "RTMP服务监听地址")
	pflag.String("httpflv_addr", ":7001", "HTTP-FLV服务监听地址")
	pflag.String("hls_addr", ":7002", "HLS服务监听地址")
	pflag.String("api_addr", ":8090", "HTTP接口管理服务监听地址")
	pflag.String("config_file", "config/config.yaml", "配置文件名称")
	pflag.String("level", "info", "日志等级")
	pflag.Bool("hls_keep_after_end", false, "流结束后保持HLS")
	pflag.String("flv_dir", "tmp", "输出flv文件路径：flvDir/APP/KEY_TIME.flv")
	pflag.Int("read_timeout", 10, "读超时时间")
	pflag.Int("write_timeout", 10, "写超时时间")
	pflag.Int("gop_num", 1, "gop数量")
	pflag.Parse()
	_ = Config.BindPFlags(pflag.CommandLine)

	// 设置配置文件
	Config.SetConfigFile(Config.GetString("config_file"))
	// 配置文件路径
	Config.AddConfigPath(".")
	err := Config.ReadInConfig()
	if err != nil {
		log.Warning(err)
		log.Info("使用默认配置")
	} else {
		_ = Config.MergeInConfig()
	}

	// 环境
	replacer := strings.NewReplacer(".", "_")
	Config.SetEnvKeyReplacer(replacer)
	Config.AllowEmptyEnv(true)
	Config.AutomaticEnv()

	// 日志
	initLog()

	// 打印最终配置
	c := ServerCfg{}
	_ = Config.Unmarshal(&c)
	log.Debugf("当前配置：\n%# v", pretty.Formatter(c))
}

func CheckAppName(appName string) bool {
	apps := Applications{}
	_ = Config.UnmarshalKey("server", &apps)
	for _, app := range apps {
		if app.Appname == appName {
			return app.Live
		}
	}
	return false
}

func GetStaticPushUrlList(appName string) ([]string, bool) {
	apps := Applications{}
	_ = Config.UnmarshalKey("server", &apps)
	for _, app := range apps {
		if (app.Appname == appName) && app.Live {
			if len(app.StaticPush) > 0 {
				return app.StaticPush, true
			} else {
				return nil, false
			}
		}
	}
	return nil, false
}

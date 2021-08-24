package setting

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type App struct {
	JwtSecret string `mapstructure:"jwt_secret"`
	PageSize  int    `mapstructure:"page_size"`
	PrefixUrl string `mapstructure:"prefix_url"`

	RuntimeRootPath string `mapstructure:"runtime_root_path"`

	ImageSavePath  string   `mapstructure:"image_save_path"`
	ImageMaxSize   int      `mapstructure:"image_max_size"`
	ImageAllowExts []string `mapstructure:"image_allow_exts"`

	ExportSavePath string `mapstructure:"export_save_path"`
	QrCodeSavePath string `mapstructure:"qr_code_save_path"`
	FontSavePath   string `mapstructure:"font_save_path"`

	LogSavePath string `mapstructure:"log_save_path"`
	LogSaveName string `mapstructure:"log_save_name"`
	LogFileExt  string `mapstructure:"log_file_ext"`
	TimeFormat  string `mapstructure:"time_format"`
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func Setup() {
	viper.SetConfigName("app")
	// location related to str.go
	viper.AddConfigPath("conf/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse `conf/app.yml`: %v", err)
	}
	parse(AppSetting)
	parse(ServerSetting)
	parse(DatabaseSetting)
	parse(RedisSetting)
}

func parse(setting interface{}) {
	err := viper.Unmarshal(setting)

	if err != nil {
		log.Fatalf("parse error when parsing setting %v, %v", setting, err)
	}
}

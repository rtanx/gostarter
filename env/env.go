package env

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/rtanx/gostarter/internal/infrastructure/logger"
	"github.com/spf13/viper"
)

type Env string

const (
	PRODUCTION  Env = "PRODUCTION"
	PROD        Env = "PROD"
	STAGING     Env = "STAGING"
	STG         Env = "STG"
	DEVELOPMENT Env = "DEVELOPMENT"
	DEV         Env = "DEV"
)

var runningMode Env
var configBasePath string

func init() {
	var ok bool
	configBasePath, ok = confBasePath()
	if !ok {
		panic("cannot retrive base path of configuration definition")
	}
	initConfig()
}

func confBasePath() (string, bool) {
	rootpath, ok := AppRootPath()
	if !ok {
		return "", ok
	}
	return fmt.Sprintf("%s/.conf", rootpath), ok
}

func ConfigBasePath() string {
	return configBasePath
}

func AppRootPath() (string, bool) {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		return "", ok
	}
	rootpath := filepath.Join(filepath.Dir(b), "..")
	return rootpath, ok
}

func initConfig() {
	viper.AddConfigPath(configBasePath)

	ev := Env(strings.ToUpper(os.Getenv("ENV")))
	switch ev {
	case PRODUCTION, PROD:
		viper.SetConfigName(".prod")
		runningMode = PROD
	case STAGING, STG:
		viper.SetConfigName(".stg")
		runningMode = STG
	case DEVELOPMENT, DEV:
		viper.SetConfigName(".dev")
		runningMode = DEV
	default:
		logger.Warn("environment mode is not specified or there maybe a typo while specifying the environment mode")
		logger.Info("using default environment mode: DEV")
		viper.SetConfigName(".dev")
		runningMode = DEV
	}
	fmt.Printf("looking environment configuration in %s\n", configBasePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error while reading environment configuration file: %w", err))
	}
	logger.Info("using environment configuration file:", logger.String("file", viper.ConfigFileUsed()))

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		logger.Info("environment configuration file changed ", logger.String("file", in.Name))
	})

}

func RunningMode() Env {
	return runningMode
}

package config

import (
	"log"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var conf *Config

type Config struct {
	envPrefix string
	configDir string
	fileType  string
	val       map[string]*viper.Viper
	mu        sync.Mutex
}

func New(opts ...Option) *Config {
	c := Config{
		envPrefix: "app",
		fileType:  fileTypeYaml,
		val:       make(map[string]*viper.Viper),
	}
	for _, opt := range opts {
		opt(&c)
	}

	conf = &c
	return &c
}

func Load(filename string, val any, hook func(v *viper.Viper)) error {
	return conf.Load(filename, val, hook)
}

// Load scan data to struct

func (c *Config) Load(filename string, val any, hook func(v *viper.Viper)) error {
	v, err := c.LoadWithType(filename, hook)
	if err != nil {
		return err
	}

	if err = v.Unmarshal((val)); err != nil {
		return err
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		if err := v.Unmarshal(val); err != nil {
			log.Printf("unmarshal config error: %v, 保留旧配置", err)
			return
		}
	})

	v.WatchConfig()
	return nil
}

// LoadWithType load conf by file type
func LoadWithType(filename string, hook func(v *viper.Viper)) (*viper.Viper, error) {
	return conf.LoadWithType(filename, hook)
}

func (c *Config) LoadWithType(filename string, hook func(v *viper.Viper)) (v *viper.Viper, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.val[filename]
	if ok {
		return v, nil
	}

	v, err = c.load(filename, hook)
	if err != nil {
		return nil, err
	}

	c.val[filename] = v
	return v, nil
}

func (c *Config) load(filename string, hook func(v *viper.Viper)) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(c.configDir)
	v.SetConfigName(filename)
	v.SetConfigType((c.fileType))
	v.AutomaticEnv()
	v.SetEnvPrefix(c.envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if hook != nil {
		hook(v)
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	log.Println("Using config file:", v.ConfigFileUsed(), " settings: ", v.AllSettings())

	return v, nil

}

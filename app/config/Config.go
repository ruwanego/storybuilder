package config

import "time"

// Config is the master config struct that holds all other config structs.
type Config struct {
	AppConfig      AppConfig
	DBConfig       DBConfig
	LogConfig      LogConfig
	ServicesConfig []ServiceConfig
}

// AppConfig holds application configurations.
type AppConfig struct {
	Name     string        `yaml:"name"`
	Mode     string        `yaml:"mode"`
	Host     string        `yaml:"host"`
	Port     int           `yaml:"port"`
	Timeout  TimeoutConfig `yaml:"timeout"`
	Timezone string        `yaml:"timezone"`
	Metrics  MetricConfig  `yaml:"metrics"`
	Cache    CacheConfig   `yaml:"cache"`
}

type (
	DurString     string
	TimeoutConfig struct {
		Read  DurString `yaml:"read"`
		Write DurString `yaml:"write"`
		Idle  DurString `yaml:"idle"`
	}
)

func (d DurString) Dur() time.Duration {
	tDur, _ := time.ParseDuration(string(d))
	return tDur
}

// DBConfig holds database configurations.
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"pool_size"`
	Check    bool   `yaml:"check"`
}

// LogConfig holds application log configurations.
type LogConfig struct {
	Level     string `yaml:"level"`
	Colors    bool   `yaml:"colors"`
	Console   bool   `yaml:"console"`
	File      bool   `yaml:"file"`
	Directory string `yaml:"directory"`
}

// MetricConfig holds application metric configurations.
type MetricConfig struct {
	Enabled bool   `yaml:"enabled"`
	Port    int    `yaml:"port"`
	Route   string `yaml:"route"`
}

// ServiceConfig holds service configurations.
type ServiceConfig struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Timeout int    `yaml:"timeout"`
}

// CacheConfig holds cache configurations.
type CacheConfig struct {
	LifeWindow  string `yaml:"life-window"`
	HardMaxSize int    `yaml:"hard-max-size"`
}

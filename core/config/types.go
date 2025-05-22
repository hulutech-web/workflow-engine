package config

type Config struct {
	Server   Server   `yaml:"server" json:"server"`
	Database Database `yaml:"database" json:"database"`
	Redis    Redis    `yaml:"redis" json:"redis"`
	Jwt      Jwt      `yaml:"jwt" json:"jwt"`
	Logging  Logging  `yaml:"logging" json:"logging"`
	Storage  Storage  `yaml:"storage" json:"storage"`
}

type Server struct {
	Port      int    `yaml:"port" json:"port"`
	Host      string `yaml:"host" json:"host"`
	Mode      string `yaml:"module" json:"module"`
	PublicUrl string `yaml:"publicUrl" json:"publicUrl"`
}

// Logging 配置日志文件
type Logging struct {
	Level      string `json:"level" yaml:"level"`
	FilePath   string `json:"file_path" yaml:"file_path"`
	MaxSize    int    `json:"maxsize" yaml:"maxsize"`
	MaxBackups int    `json:"maxbackups" yaml:"maxbackups"`
	MaxAge     int    `json:"maxage" yaml:"maxage"`
}

type Database struct {
	Driver       string `yaml:"driver" json:"driver"`
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	Database     string `yaml:"database" json:"database"`
	TablePrefix  string `yaml:"table_prefix" json:"table_prefix"`
	Params       string `yaml:"params" json:"params"`
	AutoMigrate  bool   `mapstructure:"auto_migrate"` // 显式映射字段名
	MaxOpenConns int    `yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxLifeTime  int    `yaml:"max_life_time" json:"max_life_time"`
}

type Redis struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
	Prefix   string `yaml:"prefix" json:"prefix"`
	PoolSize int    `yaml:"pool_size" json:"pool_size"`
}

type Jwt struct {
	Secret        string `yaml:"secret" json:"secret"`
	AccessExpiry  string `yaml:"access_expiry" json:"access_expiry"`
	RefreshExpiry string `yaml:"refresh_expiry" json:"refresh_expiry"`
}

type Storage struct {
	PublicPrefix string `yaml:"publicPrefix" json:"publicPrefix"`
	LocalPath    string `yaml:"localPath" json:"localPath"`
}

package config

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Mode         string `yaml:"mode"`
		Dsn          string `yaml:"dsn"`
		MaxIdleConns int    `yaml:"max_idle_conns"`
		MaxOpenCons  int    `yaml:"max_open_cons"`
	} `yaml:"database"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
	Jwt struct {
		Expires int    `yaml:"expires"`
		Issuer  string `yaml:"issuer"`
		Key     string `yaml:"key"`
	} `yaml:"jwt"`
	Upload struct {
		Size int64  `yaml:"size"`
		Dir  string `yaml:"dir"`
	} `yaml:"upload"`
}

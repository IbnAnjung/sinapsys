package config

type Config struct {
	Http  HttpConfig
	Mysql MysqlConfig
	Jwt   JwtConfig
	Redis RedisConfig
}

type HttpConfig struct {
	Port string
}

type MysqlConfig struct {
	User               string
	Password           string
	Host               string
	Port               string
	Schema             string
	Timeout            int
	MaxIddleConnection int
	MaxOpenConnection  int
	MaxLifeTime        int
}

type JwtConfig struct {
	SecretKey            string
	AccessTokenLifeTime  int
	RefreshTokenLifeTime int
}

type RedisConfig struct {
	Addr           string
	Username       string
	Password       string
	DB             int
	MinIdleConns   int
	MaxIdleConns   int
	MaxActiveConns int
}

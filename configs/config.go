package configs

type Db struct {
	Name     string
	User     string
	Host     string
	Password string
	Port     int
}

type Server struct {
	Host string
	Port int
}

type Config struct {
	Db     Db
	Server Server
}

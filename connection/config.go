package connection

// Config represents  the basic information to connects to database
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Driver   string
	Schema   string
	Orcl     string
}

package config

// Configuration is a struct represeting the database config
type Configuration struct {
	DatabaseType      string
	Database          string
	User              string
	Password          string
	MaxConnectionTime int32
	MaxIdleConns      int
	MaxOpenConns      int
}

// ConnectionString returns the MySQL connection string for database
func (c Configuration) ConnectionString() string {
	return c.User + ":" + c.Password + "@/" + c.Database + "?charset=utf8&parseTime=True&loc=Local"
}

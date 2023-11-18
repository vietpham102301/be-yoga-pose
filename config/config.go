package config

const (
	DBUsername = "root"
	DBPassword = "viet1234"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "yoga_support"
)

func GetMySQLURL() string {
	return DBUsername + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?parseTime=true"
}

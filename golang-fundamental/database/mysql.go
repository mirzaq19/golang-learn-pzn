package database

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabaseConnection() string {
	return connection
}

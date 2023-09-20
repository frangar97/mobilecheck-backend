package config

import (
	"os"
)

type config struct {
	Port                 string
	DatabaseUrl          string
	SqlServerDatabaseUrl string
}

func InitConfig() (*config, error) {
	// err := godotenv.Load()

	// if err != nil {
	// 	return nil, err
	// }

	//sqlServerDSN := BuildSqlServerConnection()

	cfg := config{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		//SqlServerDatabaseUrl: sqlServerDSN,
	}

	return &cfg, nil
}

// func BuildSqlServerConnection() string {
// 	server := os.Getenv("DB_SERVER")
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	database := os.Getenv("DB_NAME")
// 	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
// 		server, user, password, port, database)
// 	return connectionString
// }

package config

import "os"

var Global AppConfig

// Application configuration options
type AppConfig struct {
	Scope string

	// Database
	DBDriverName string
	DBUsername   string
	DBPassword   string
	DBHost       string
	DBName       string
}

func init() {
	Global = AppConfig{
		Scope: os.Getenv("SCOPE"),

		// Database
		DBDriverName: "mysql",
		DBUsername:   "root",
		DBPassword:   "adrianj16",
		DBHost:       "localhost",
		DBName:       "biller",
	}
}

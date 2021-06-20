package config

import (
	"flag"
	"fmt"
	"os"
)

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
	SSL  string
}

type PusherConfig struct {
	Host   string
	Port   string
	App    string
	Key    string
	Secret string
	Secure bool
}

type Config struct {
	Port         string
	Identifier   string
	Domain       string
	InProduction bool
	DBConfig     DBConfig
	PusherConfig PusherConfig
	Config       map[string]string
}

func NewConfig() *Config {
	insecurePort := flag.String("port", ":4000", "port to listen on")
	identifier := flag.String("identifier", "vigilate", "unique identifier")
	domain := flag.String("domain", "localhost", "domain name (e.g. example.com)")
	inProduction := flag.Bool("production", false, "application is in production")
	dbHost := flag.String("dbhost", "localhost", "database host")
	dbPort := flag.String("dbport", "5432", "database port")
	dbUser := flag.String("dbuser", "", "database user")
	dbPass := flag.String("dbpass", "", "database password")
	dbName := flag.String("dbname", "vigilate", "database name")
	dbSsl := flag.String("dbssl", "disable", "database ssl setting")
	pusherHost := flag.String("pusherHost", "", "pusher host")
	pusherPort := flag.String("pusherPort", "443", "pusher port")
	pusherApp := flag.String("pusherApp", "9", "pusher app id")
	pusherKey := flag.String("pusherKey", "", "pusher key")
	pusherSecret := flag.String("pusherSecret", "", "pusher secret")
	pusherSecure := flag.Bool("pusherSecure", false, "pusher server uses SSL (true or false)")

	flag.Parse()

	if *dbUser == "" || *dbHost == "" || *dbPort == "" || *dbName == "" {
		fmt.Println("Missing required flags.")
		os.Exit(1)
	}

	return &Config{
		Port:         *insecurePort,
		Identifier:   *identifier,
		Domain:       *domain,
		InProduction: *inProduction,
		DBConfig: DBConfig{
			Host: *dbHost,
			Port: *dbPort,
			User: *dbUser,
			Pass: *dbPass,
			Name: *dbName,
			SSL:  *dbSsl,
		},
		PusherConfig: PusherConfig{
			Host:   *pusherHost,
			Port:   *pusherPort,
			App:    *pusherApp,
			Key:    *pusherKey,
			Secret: *pusherSecret,
			Secure: *pusherSecure,
		},
	}
}

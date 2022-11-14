package config

import (
	"hippo/logging"
	"io/ioutil"
	"os"

	firebase "firebase.google.com/go"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   server        `yaml:"server"`
	Database database      `yaml:"database"`
	Firebase firebaseAdmin `yaml:"firebase"`
}

type server struct {
	Port     int    `yaml:"port"`
	BasePath string `yaml:"basePath"`
}

type database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type firebaseAdmin struct {
	Path string `yaml:"path"`
	App  *firebase.App
}

func GetConfig() *Config {
	logging.Info.Print("Reading in server configurations")
	configFile, err := ioutil.ReadFile("config/config.yaml")

	if err != nil {
		logging.Fatal.Fatalf("Fatal error reading in config.yaml: %v", err)
	}

	conf := &Config{}

	err = yaml.Unmarshal(configFile, conf)

	if err != nil {
		logging.Fatal.Fatalf("Fatal error decoding config.yaml: %v", err)
	}

	app, err := InitFirebase(conf)

	if err != nil {
		logging.Fatal.Fatalf("Fatal error initializing Firebase admin: %v", err)
	}

	dbPassword := os.Getenv(conf.Database.Password)
	if len(dbPassword) == 0 {
		logging.Fatal.Fatal("Fatal error retrieving database password")
	}

	return &Config{
		Server: conf.Server,
		Database: database{
			Host:     conf.Database.Host,
			Port:     conf.Database.Port,
			Name:     conf.Database.Name,
			Username: conf.Database.Username,
			Password: dbPassword,
		},
		Firebase: firebaseAdmin{
			Path: conf.Firebase.Path,
			App:  app,
		},
	}
}

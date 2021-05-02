package config

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	Port            string
	DatabaseDialect string
	DatabaseUrl     string
	GrpcPort        string
}

var (
	databaseDialect = *flag.String("dialect", os.Getenv("DATABASE_DIALECT"), "Database connection string")
	databaseUrl     = *flag.String("db", os.Getenv("DATABASE_URL"), "Database connection string")
	port            = *flag.String("HTTP port", os.Getenv("PORT"), "port for HTTP")
	grpcport        = *flag.String("GRPC port", os.Getenv("GRPCPORT"), "port for HTTP")
)

func Load() (config Config) {
	config.parse()
	config.validate()
	flag.Parse()

	return config
}

// validate configuration.
func (c *Config) validate() {
	if len(c.DatabaseDialect) == 0 || len(c.DatabaseUrl) == 0 {
		log.Panic("DATABASE_DIALECT and DATABASE_URL required.")
	}
}

func (c *Config) parse() {
	flag.Parse()
	// Switches
	c.DatabaseDialect = databaseDialect
	c.DatabaseUrl = databaseUrl
	c.Port = port
	c.GrpcPort = grpcport
}

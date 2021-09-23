package support

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const prefix = "XAGENDA_%s"

// Config is the configuration for the server. It is populated from flags and environment variables.
// Please see the Load method for more information.
type Config struct {
	Addr     string
	Timeout  time.Duration
	Dsn      string
	LogLevel uint
}

// Load loads the configuration from flags and environment variables, in that order.
// If a value is not set in either place, the default value is used, if defined, otherwise returns a MissingConfigError.
// Environment variables are expected to be in the form of XAGENDA_<field name in uppercase>.
func (cfg *Config) Load() error {
	flag.StringVar(&cfg.Addr, "host", lookupEnvOrString("ADDR", "localhost:8080"), "TCP address for the server to listen on")
	flag.DurationVar(&cfg.Timeout, "timeout", lookupEnvOrDuration("TIMEOUT", time.Second*10), "Timeout for the server to wait for a response")
	flag.StringVar(&cfg.Dsn, "dsn", lookupEnvOrString("DSN", ""), "DSN for the database")
	flag.UintVar(&cfg.LogLevel, "log-level", lookupEnvOrUint("LOG_LEVEL", uint(LevelError)), "Log level for the server")
	flag.Parse()
	return cfg.validate()
}

func (cfg *Config) validate() error {
	if cfg.Dsn == "" {
		return &MissingConfigError{Field: "DSN"}
	}
	return nil
}

func lookupEnvOrString(env string, def string) string {
	if val, ok := os.LookupEnv(fmt.Sprintf(prefix, env)); ok {
		return val
	}
	return def
}

func lookupEnvOrDuration(env string, def time.Duration) time.Duration {
	if val, ok := os.LookupEnv(fmt.Sprintf(prefix, env)); ok {
		d, err := time.ParseDuration(val)
		if err != nil {
			return def
		}
		return d
	}
	return def
}

func lookupEnvOrUint(env string, def uint) uint {
	if val, ok := os.LookupEnv(fmt.Sprintf(prefix, env)); ok {
		i, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return def
		}
		return uint(i)
	}
	return def
}

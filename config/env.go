package config

import "os"

//configuration template

const (
	envKey   = "APP_ENV" //environment key
	EnvProd  = "prod"    // environment production
	EnvLocal = "local"   // environment local, default environment
)

var env = GetEnv(envKey, EnvLocal) //env variable

func GetEnv(key, def string) string {
	envr, ok := os.LookupEnv(key) //returns the environment name and a boolean
	if ok {
		return envr //if ok is true, return env name by the key
	} else {
		return def //if ok is false, return default env name
	}
}

func Environment() string { //return env variable
	return env
}

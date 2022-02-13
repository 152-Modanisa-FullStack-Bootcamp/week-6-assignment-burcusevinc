package config

import (
	"encoding/json"
	"io"
	"os"
)

//configuration read

type Config struct {
	//json field match
	InitialBalance int `json:"initialBalanceAmount"`
	MinumumBalance int `json:"minumumBalanceAmount"`
}

//linking json fields with application

//struct
var C = &Config{}

func init() {
	//open file
	file, err := os.Open(".config/" + env + ".json")
	if err != nil {
		panic(err)
	}

	//close file, when the process finish.
	defer file.Close()

	//read file
	read, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// json fields are kept in our struct
	err = json.Unmarshal(read, C)
	if err != nil {
		panic(err)
	}
}

func NewConfig() *Config {
	return C
}

package config

import (
	"os"
)

func ListenPort() string {
	listenport := ":" + os.Getenv("SERVER_PORT")
	return listenport
}
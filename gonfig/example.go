package main

import (
	"fmt"

	flags "github.com/vrianta/golang/gonfig/v1"
)

var Flags = flags.New[struct {
	Host   string `env:"APP_HOST" arg:"host" default:"localhost"`
	Port   int    `env:"APP_PORT" default:"8080"`
	Debug  bool   `arg:"debug" default:"false"`
	ApiKey string `env:"APP_API_KEY" arg:"apikey" required:""`
}](true)

func main() {

	fmt.Println(Flags.Host)

}

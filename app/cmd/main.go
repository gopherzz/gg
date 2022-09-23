package main

import (
	"fmt"
	"log"

	"github.com/gopherzz/gg/app/internal/ui"
	"github.com/gopherzz/gg/app/pkg/config"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg config.Config

func main() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("read config error: %s", err.Error())
	}
	ui := ui.NewUi(&cfg)
	log.Fatal(ui.Render())
	fmt.Println(cfg)
}

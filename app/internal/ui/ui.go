package ui

import (
	"fmt"

	"github.com/gopherzz/gg/app/pkg/config"
)

type Ui struct {
	Config *config.Config
}

func (ui Ui) Render() {
	fmt.Println("rendered!")
}

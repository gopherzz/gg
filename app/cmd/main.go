package main

import (
	"fmt"

	"github.com/gopherzz/gg/app/internal/ui"
)

func main() {
	fmt.Println("Hello!")
	ui := ui.Ui{Config: nil}
	ui.Render()
}

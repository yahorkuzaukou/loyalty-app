package main

import (
	"github.com/yahorkuzaukou/loyalty-app/cmd"
)

func main() {
	app := cmd.NewApp()
	app.Listen("0.0.0.0:3000")
}

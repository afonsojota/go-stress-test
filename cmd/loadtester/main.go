package main

import (
	"github.com/afonsojota/go-stress-test/internal/infrastructure"
)

func main() {
	container := infrastructure.NewContainer()

	cli := container.GetCLI()

	cli.Execute()
}

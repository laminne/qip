package main

import (
	"fmt"
	"os"

	"github.com/approvers/qip/pkg/router"
	"github.com/approvers/qip/pkg/utils/config"
)

func main() {
	f, _ := os.Open("config.yml")
	fmt.Println(config.LoadConfig(f))
	router.StartServer(7000)
}

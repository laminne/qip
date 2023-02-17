package main

import (
	"fmt"
	"github.com/approvers/qip/pkg/router"
	"github.com/approvers/qip/pkg/utils/config"
	"os"
)

func main() {
	f, _ := os.Open("config.yml")
	fmt.Println(config.LoadConfig(f))
	router.StartServer(7000)
}

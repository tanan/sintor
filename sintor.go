package main

import (
	"flag"
	"github.com/tanan/sintor/infrastructure/web"
	"github.com/tanan/sintor/config"
	"strconv"
	"log"
)

var configPath = flag.String("c", "config.toml", "config file path")

func main() {
	flag.Parse()

	con, err := config.LoadFile(*configPath)
	if err != nil {
		log.Fatalf("cannot read a config file. error: %v\n", err)
	}
	router := web.LoadRouter(con)
	router.Logger.Fatal(router.Start(":" + strconv.Itoa(con.Application.Port)))
}

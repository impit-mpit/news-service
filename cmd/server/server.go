package main

import (
	"neuro-most/news-service/config"
	"neuro-most/news-service/internal/infra"
)

func main() {
	cfg, err := config.NewLoadConfig()
	if err != nil {
		panic(err)
	}
	infra.Config(cfg).Database().Serve().Start()
}

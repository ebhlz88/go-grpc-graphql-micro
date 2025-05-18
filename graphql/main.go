package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	accountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	orderURL   string `envconfig:"ORDER_SERVICE_URL"`
	catalogURL string `envconfig:"CATALOG_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	s, err := NewGraphQLServer(cfg.accountURL, cfg.catalogURL, cfg.orderURL)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("emad", "/graphql"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

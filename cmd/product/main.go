package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"

	"github.com/aditiapratama1231/graphql-example/graph"
	"github.com/aditiapratama1231/graphql-example/graph/generated"
	"github.com/aditiapratama1231/graphql-example/internal/product"
	"github.com/aditiapratama1231/graphql-example/internal/repository"
	"github.com/aditiapratama1231/graphql-example/pkg/database"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.DBInit()

	var resolver = &graph.Resolver{
		ProductResolver: productResolverInit(db),
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func productResolverInit(db *gorm.DB) product.Product {
	productRepo := repository.NewProductRepository(db)
	return product.Product{
		ProductRepository: productRepo,
	}
}

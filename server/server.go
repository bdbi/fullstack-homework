package main

import (
	"context"
	"homework-backend/ent"
	"homework-backend/graph"
	"homework-backend/graph/generated"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"

	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"
const defaultSQLite = "file:data/data.db?cache=shared&_fk=1"

func init() {
	// create a database an populate it
	if _, ok := os.LookupEnv("SQLITE_CONN"); !ok {
		os.Setenv("SQLITE_CONN", defaultSQLite)
	}
	client, err := ent.Open("sqlite3", os.Getenv("SQLITE_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	if _, err := os.Stat(".nop"); err == nil {
		log.Println("database population skipped. delete .nop file to populate database")
		return
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	tx, err := client.Tx(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = tx.Question.Create().SetBody("What is your favourite food?").SetWeight(1).SetType("TextQuestion").Save(context.Background())
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}
	opts, err := tx.Opt.CreateBulk(client.Opt.Create().SetBody("East").SetWeight(1), client.Opt.Create().SetBody("West").SetWeight(0)).Save(context.Background())
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}
	_, err = tx.Question.Create().SetBody("Where does the sun set?").SetWeight(0.5).SetType("ChoiceQuestion").AddOptions().AddOptions(opts...).Save(context.Background())
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}
	ioutil.WriteFile(".nop", []byte{'a'}, fs.ModeAppend)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	h := cors.Default().Handler(mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, h))
}

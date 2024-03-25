//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./gql_generated/ent.graphql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	epex, err := entproto.NewExtension(
		entproto.WithProtoDir("./ent/proto"),
	)

	// Create a list of entc options to pass to entc.Generate.
	opts := []entc.Option{
		entc.Extensions(ex),
		entc.Extensions(epex),
	}

	// Run the ent codegen, which will generate the Ent schema and the GraphQL schema.
	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

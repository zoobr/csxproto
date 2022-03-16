# csxproto


**Example using Hasura client with hasura http client**
```
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zoobr/csxproto/hasura"
	"github.com/hasura/go-graphql-client"
)

func main(){

  url := "http://localhost/v1/graphql"

	authOption := map[string]string{
    "X-Hasura-User-Id": "1",
    "X-Hasura-Role": "user_role",
  }
	client, err := hasura.NewHttpClient(url, authOption)
	if err != nil {
		log.Fatal(err)
	}

	type Query struct {
		Institution []struct {
			ID              graphql.Int    `graphql:"id"`
			InstitutionType graphql.String `graphql:"institution_type"`
			MaxCapacity     graphql.Int    `graphql:"max_capacity"`
			Name            graphql.String `graphql:"name"`
			Number          graphql.String `graphql:"number"`
			Path            graphql.String `graphql:"path"`
		}
	}

	q := &Query{}
	err = client.Query(context.Background(), q, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)
}
```
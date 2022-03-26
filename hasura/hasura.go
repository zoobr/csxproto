package hasura

import (
	"net/http"

	"github.com/hasura/go-graphql-client"
)

// NewHttpClient returns hasura client.
// url - hasura url,
// authOption - authentication option
func NewHttpClient(url string, authOption map[string]string) (*graphql.Client, error) {

	f := func(request *http.Request) {
		request.Header.Add("Content-type", "application/json")
		for key, val := range authOption {
			request.Header.Add(key, val)
		}
	}

	hasuraClient := graphql.NewClient(url, nil)
	client := hasuraClient.WithRequestModifier(f)

	return client, nil
}

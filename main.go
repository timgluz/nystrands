package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
	"github.com/mailru/easyjson"

	openapi "github.com/timgluz/nystrands/pkg/openapi"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		doc, err := renderDocumentation()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error rendering documentation: %v", err), http.StatusInternalServerError)
			return
		}

		w.Write(doc)
	})
}

func main() {}

func renderDocumentation() ([]byte, error) {
	doc := openapi.OpenAPI{
		OpenAPIVersion: "3.0.0",
		Info: openapi.OpenAPIInfo{
			Title:       "API Documentation",
			Description: "API to help with NYStrands",
		},
		Servers: []openapi.OpenAPIServer{
			{URL: "localhost:8080", Description: "Test instance"},
		},
	}

	return easyjson.Marshal(doc)
}

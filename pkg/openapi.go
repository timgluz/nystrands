package openapi

// OpenAPI is the root object of the OpenAPI specification.
// if you changes any of the fields, you need to update the easyjson

//easyjson:json
type OpenAPI struct {
	OpenAPIVersion string                 `json:"openapi"`
	Info           OpenAPIInfo            `json:"info"`
	Servers        []OpenAPIServer        `json:"servers"`
	Paths          map[string]OpenAPIPath `json:"paths"`
	Components     []OpenAPIComponent     `json:"components"`
}

//easyjson:json
type OpenAPIInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

//easyjson:json
type OpenAPIServer struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

//easyjson:json
type OpenAPIPath struct {
	Get OpenAPIOperation `json:"get"`
}

//easyjson:json
type OpenAPIOperation struct {
	Tags        []string                   `json:"tags"`
	Summary     string                     `json:"summary"`
	Description string                     `json:"description"`
	Parameters  []OpenAPIParameter         `json:"parameters"`
	Responses   map[string]OpenAPIResponse `json:"responses"`
}

//easyjson:json
type OpenAPIParameter struct {
}

//easyjson:json
type OpenAPIResponse struct {
	Description string `json:"description"`
}

//easyjson:json
type OpenAPIComponent struct {
	Schemas map[string]OpenAPISchema `json:"schemas"`
}

//easyjson:json
type OpenAPISchema struct {
	Type     string   `json:"type"`
	Required []string `json:"required"`
}

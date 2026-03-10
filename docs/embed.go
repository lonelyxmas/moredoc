package docs

import _ "embed"

var (
	//go:embed openapi.swagger.yaml
	OpenAPIYAML []byte

	//go:embed swagger.json
	SwaggerJSON []byte

	//go:embed swagger.yaml
	SwaggerYAML []byte
)

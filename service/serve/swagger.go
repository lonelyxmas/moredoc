package serve

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/yaml.v3"

	moredocdocs "moredoc/docs"
)

const swaggerIndexHTML = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>MoreDoc API Docs</title>
  <link rel="stylesheet" type="text/css" href="./static/swagger-ui.css">
  <link rel="icon" type="image/png" href="./static/favicon-32x32.png" sizes="32x32">
  <link rel="icon" type="image/png" href="./static/favicon-16x16.png" sizes="16x16">
  <style>
    html { box-sizing: border-box; overflow-y: scroll; }
    *, *:before, *:after { box-sizing: inherit; }
    body { margin: 0; background: #fafafa; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="./static/swagger-ui-bundle.js" charset="UTF-8"></script>
  <script src="./static/swagger-ui-standalone-preset.js" charset="UTF-8"></script>
  <script src="./swagger-initializer.js" charset="UTF-8"></script>
</body>
</html>
`

const swaggerInitializerJS = `window.onload = function() {
  window.ui = SwaggerUIBundle({
    urls: [
      { url: "./openapi.swagger.yaml", name: "MoreDoc gRPC Gateway API" },
      { url: "./doc.json", name: "Gin Upload API" }
    ],
    "urls.primaryName": "MoreDoc gRPC Gateway API",
    dom_id: '#swagger-ui',
    deepLinking: true,
    docExpansion: 'list',
    persistAuthorization: true,
    defaultModelsExpandDepth: 1,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    layout: "StandaloneLayout"
  });
};
`

const gatewayOpenAPITitle = "MoreDoc API"
const gatewayOpenAPIVersion = "1.0"
const gatewayOpenAPIDescription = "MoreDoc 后端 API 文档，包含 gRPC Gateway 暴露的接口。支持匿名访问与携带 Bearer Token 的差异化返回。"

// RegisterSwaggerRoutes 挂载统一 Swagger 文档入口。
func RegisterSwaggerRoutes(app *gin.Engine) {
	assetHandler := ginSwagger.CustomWrapHandler(&ginSwagger.Config{}, swaggerFiles.Handler)

	app.GET("/swagger", redirectSwaggerIndex)
	app.GET("/swagger/", redirectSwaggerIndex)
	app.GET("/swagger/index.html", serveSwaggerIndex)
	app.GET("/swagger/swagger-initializer.js", serveSwaggerInitializer)
	app.GET("/swagger/doc.json", serveGinSwaggerDoc)
	app.GET("/swagger/openapi.swagger.yaml", serveGatewayOpenAPI)
	app.GET("/swagger/static/*any", assetHandler)
}

func redirectSwaggerIndex(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/swagger/index.html")
}

func serveSwaggerIndex(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(swaggerIndexHTML))
}

func serveSwaggerInitializer(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/javascript", []byte(swaggerInitializerJS))
}

func serveGinSwaggerDoc(ctx *gin.Context) {
	if len(moredocdocs.SwaggerJSON) == 0 {
		ctx.String(http.StatusInternalServerError, "embedded swagger.json is empty")
		return
	}
	ctx.Data(http.StatusOK, "application/json; charset=utf-8", moredocdocs.SwaggerJSON)
}

func serveGatewayOpenAPI(ctx *gin.Context) {
	if len(moredocdocs.OpenAPIYAML) == 0 {
		ctx.String(http.StatusInternalServerError, "embedded openapi.swagger.yaml is empty")
		return
	}
	content, err := normalizeGatewayOpenAPI(moredocdocs.OpenAPIYAML)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "normalize openapi.swagger.yaml failed: %v", err)
		return
	}
	ctx.Data(http.StatusOK, "application/yaml; charset=utf-8", content)
}

func normalizeGatewayOpenAPI(content []byte) ([]byte, error) {
	var doc yaml.Node
	if err := yaml.Unmarshal(content, &doc); err != nil {
		return nil, err
	}

	root := documentRoot(&doc)
	if root == nil || root.Kind != yaml.MappingNode {
		return nil, http.ErrNotSupported
	}

	info := ensureMapValue(root, "info")
	setScalarMapValue(info, "title", gatewayOpenAPITitle)
	setScalarMapValue(info, "version", gatewayOpenAPIVersion)
	setScalarMapValue(info, "description", gatewayOpenAPIDescription)

	securityDefinitions := ensureMapValue(root, "securityDefinitions")
	bearerAuth := ensureMapValue(securityDefinitions, "BearerAuth")
	setScalarMapValue(bearerAuth, "type", "apiKey")
	setScalarMapValue(bearerAuth, "name", "Authorization")
	setScalarMapValue(bearerAuth, "in", "header")
	setScalarMapValue(bearerAuth, "description", "Optional JWT bearer token. Example: Bearer <token>")

	security := ensureSequenceValue(root, "security")
	ensureOptionalAnonymousSecurity(security)
	ensureBearerSecurityRequirement(security)

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(&doc); err != nil {
		return nil, err
	}
	if err := encoder.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func documentRoot(doc *yaml.Node) *yaml.Node {
	if doc.Kind == yaml.DocumentNode && len(doc.Content) > 0 {
		return doc.Content[0]
	}
	return doc
}

func ensureMapValue(parent *yaml.Node, key string) *yaml.Node {
	if value := mapValue(parent, key); value != nil && value.Kind == yaml.MappingNode {
		return value
	}

	value := &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"}
	setMapValue(parent, key, value)
	return value
}

func ensureSequenceValue(parent *yaml.Node, key string) *yaml.Node {
	if value := mapValue(parent, key); value != nil && value.Kind == yaml.SequenceNode {
		return value
	}

	value := &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
	setMapValue(parent, key, value)
	return value
}

func ensureOptionalAnonymousSecurity(security *yaml.Node) {
	for _, item := range security.Content {
		if item.Kind == yaml.MappingNode && len(item.Content) == 0 {
			return
		}
	}

	security.Content = append(security.Content, &yaml.Node{Kind: yaml.MappingNode, Tag: "!!map"})
}

func ensureBearerSecurityRequirement(security *yaml.Node) {
	for _, item := range security.Content {
		if item.Kind != yaml.MappingNode {
			continue
		}
		for i := 0; i+1 < len(item.Content); i += 2 {
			if item.Content[i].Value == "BearerAuth" {
				if item.Content[i+1].Kind != yaml.SequenceNode {
					item.Content[i+1] = &yaml.Node{Kind: yaml.SequenceNode, Tag: "!!seq"}
				}
				return
			}
		}
	}

	security.Content = append(security.Content, &yaml.Node{
		Kind: yaml.MappingNode,
		Tag:  "!!map",
		Content: []*yaml.Node{
			scalarNode("BearerAuth"),
			{Kind: yaml.SequenceNode, Tag: "!!seq"},
		},
	})
}

func setScalarMapValue(parent *yaml.Node, key, value string) {
	setMapValue(parent, key, scalarNode(value))
}

func setMapValue(parent *yaml.Node, key string, value *yaml.Node) {
	if parent.Kind != yaml.MappingNode {
		parent.Kind = yaml.MappingNode
		parent.Tag = "!!map"
		parent.Content = nil
	}

	for i := 0; i+1 < len(parent.Content); i += 2 {
		if parent.Content[i].Value == key {
			parent.Content[i+1] = value
			return
		}
	}

	parent.Content = append(parent.Content, scalarNode(key), value)
}

func mapValue(parent *yaml.Node, key string) *yaml.Node {
	if parent == nil || parent.Kind != yaml.MappingNode {
		return nil
	}

	for i := 0; i+1 < len(parent.Content); i += 2 {
		if parent.Content[i].Value == key {
			return parent.Content[i+1]
		}
	}

	return nil
}

func scalarNode(value string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: value}
}

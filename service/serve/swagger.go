package serve

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"

	_ "moredoc/docs"
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
      { url: "./openapi.yaml", name: "gRPC Gateway API" },
      { url: "./doc.json", name: "Gin Upload API" }
    ],
    "urls.primaryName": "gRPC Gateway API",
    dom_id: '#swagger-ui',
    deepLinking: true,
    docExpansion: 'list',
    persistAuthorization: true,
    defaultModelsExpandDepth: 1,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });
};
`

// RegisterSwaggerRoutes 挂载统一 Swagger 文档入口。
func RegisterSwaggerRoutes(app *gin.Engine) {
	assetHandler := ginSwagger.CustomWrapHandler(&ginSwagger.Config{}, swaggerFiles.Handler)

	app.GET("/swagger", redirectSwaggerIndex)
	app.GET("/swagger/", redirectSwaggerIndex)
	app.GET("/swagger/index.html", serveSwaggerIndex)
	app.GET("/swagger/swagger-initializer.js", serveSwaggerInitializer)
	app.GET("/swagger/doc.json", serveGinSwaggerDoc)
	app.GET("/swagger/openapi.yaml", serveGatewayOpenAPI)
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
	doc, err := swag.ReadDoc(swag.Name)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Data(http.StatusOK, "application/json; charset=utf-8", []byte(doc))
}

func serveGatewayOpenAPI(ctx *gin.Context) {
	ctx.File("./docs/openapi.yaml")
}

/*
Copyright © 2022 MNT.Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"moredoc/cmd"
)

// @title MoreDoc API
// @version 1.0
// @description MoreDoc 后端 API 文档，包含 Gin 原生接口的 Swagger 描述。
// @description 通过 gRPC Gateway 暴露的接口仍按 proto/openapi 流程维护，Swagger 页面主要覆盖 Gin 侧文件上传与静态资源相关接口。
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}

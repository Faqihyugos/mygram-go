package main

import (
	"github.com/Faqihyugos/mygram-go/handler"
)

// @title          Mygram Golang
// @version        1.0
// @description    Final Project FGA Golang MyGram, Go Programming Language
// @termsOfService http://swagger.io/terms/

// @contact.name  Faqih Yugos
// @contact.url   http://www.swagger.io/support
// @contact.email faqihyugos@gmail.com
// @BasePath       /

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description				   Description for what is this security definition being used

func main() {
	handler.StartApp()
}

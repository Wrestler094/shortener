// @title           Сервис сокращения ссылок
// @version         1.0
// @description     API для создания, получения и удаления сокращённых URL
// @termsOfService  https://example.com/terms

// @contact.name   Михаил Ковалев
// @contact.email  kovalev094@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"fmt"

	_ "github.com/Wrestler094/shortener/docs" // путь к сгенерированной документации

	"github.com/Wrestler094/shortener/internal/app"
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func main() {
	printBuildInfo()

	app.Run()
}

func printBuildInfo() {
	fmt.Println("Build version:", getOrNA(buildVersion))
	fmt.Println("Build date:   ", getOrNA(buildDate))
	fmt.Println("Build commit: ", getOrNA(buildCommit))
}

func getOrNA(s string) string {
	if s == "" {
		return "N/A"
	}
	return s
}

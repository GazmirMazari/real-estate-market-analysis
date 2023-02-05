package main //ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/general_api_info.html
import (
	"fmt" // @title Swagger Example API
	// @version 1.0
	// @description This is a sample server celler server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8080
	// @BasePath /api/v1

	// @securityDefinitions.basic BasicAuth

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	// @securitydefinitions.oauth2.application OAuth2Application
	// @tokenUrl https://example.com/oauth/token
	// @scope.write Grants write access
	// @scope.admin Grants read and write access to administrative information

	redfin "real_estate_market/cmd/api"
)

func main() {
	// // Start the proxy server
	// go func() {
	// 	err := http.ListenAndServe("localhost:8080", colly.NewCollector())
	// 	if err != nil {
	// 		log.Fatalf("ListenAndServe: %v", err)
	// 	}
	// }()

	properties, err := redfin.Scrape("https://www.redfin.com/city/9170/IN/Indianapolis")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, property := range properties {
		fmt.Printf("%s, %s, %s %s, %s\n", property.Address, property.City, property.State, property.Zip, property.Price)
	}
}

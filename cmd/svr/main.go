package main //ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/general_api_info.html
// @title Swagger Example API
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

import (
	"log"
	"net/http"

	httpRoutes "github.com/GazmirMazari/real-estate-market-analysis/routes"
)

func main() {
	//TODO: add swagger
	//TODO: add swagger ui
	//TODO: Create a system design document which will layout how will we be implementing the system
	//TODO:	Data collection: Use the Open API provided by Redfin and Zillow to collect real estate market data. This can be done using the Go standard library's net/http package to makeHTTP requests to the APIs and the encoding/json package to parse the responses.
	//TODO:	Data validation: Implement data validation to ensure that the data received by the API is in the correct format and contains all the required fields. This can be done using apackage such as go-validator or govalidator.
	//TODO:	Data storage: Store the collected and validated data in a GraphDB. This will allow for efficient querying and data relationships.
	//TODO:	Redis for caching: Use Redis as a cache for frequently accessed data. This can be done using the redigo package or other Go Redis clients.
	//TODO:	API endpoint: Create API endpoints to handle requests from clients. These endpoints should handle authentication and authorization using package such as jwt-go or go-oauth2,and //use the cache system to return data quickly.
	//TODO:	Concurrency: Use Go's concurrency features, such as goroutines and channels, to handle multiple requests simultaneously. This will improve the performance of the API.
	//TODO:	Middleware: Implement middleware to handle common tasks such as logging, error handling, and monitoring.
	//TODO:	Gomock for testing: Use Gomock to mock dependencies and test your code in isolation.
	//TODO:	Logging: Implement logging to track events and errors that occur during the execution of the API. This can be done using a Go package such as logrus or zap.
	//TODO:	Monitoring: Implement monitoring to track the performance and health of the API. This can be done using a Go package such as prometheus or go-metrics, which can providemetrics //such as request count, error rate, and response time.
	//TODO:	Go Fuzz for integration testing: Use Go Fuzz to test the API by providing it with randomized input. This can help to detect and prevent potential bugs and crashes.
	//TODO:	Caching data for integration testing: Use the data stored in Redis cache to test the API's integration with the GraphDB, this will make testing more efficient and provide more /realistic testing scenarios.
	//TODO:	Migration: Implement data migration to handle changes in the data model. This can be done using a migration tool such as goose or gormigrate.
	//TODO:	Session management: Implement session management to track user sessions and handle authentication and authorization. This can be done using a session management library suchas //gorilla/sessions or context-session.
	//TODO:	Deployment: Consider a containerization solution such as Docker to deploy your service, this will make it easy to manage and scaling your service.
	//TODO:	Scaling: Consider scaling the service horizontally by adding more instances of the service behind a load balancer. This can be done using a Kubernetes or other container orchestration platform.

	mux := httpRoutes.

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

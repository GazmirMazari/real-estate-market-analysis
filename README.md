# real-estate-market-analysis
Backend built in go to analyze real estate market

# This repo it is built to have a better analysis of how go routines and channel works. There is a massive a TODO list which I would need to go through


- [ ]    Add swagger for API documentation
- [ ]    Add swagger UI for easy visualization and testing of the API
- [ ]    Create a system design document outlining the implementation of the system
- [ ]    Data collection: Use Go's colly library and go routines to scrape real estate market data from websites such as Redfin and Zillow
- [ ]    Data validation: Implement data validation to ensure that the collected data is in the correct format and contains all necessary fields, using a package such as go-validator or govalidator
- [ ]    Data storage: Store the collected and validated data in a GraphDB for efficient querying and data relationships
- [ ]    Implement Redis for caching frequently accessed data, using a package such as redigo or other Go Redis clients
- [ ]    Create API endpoints to handle requests from clients, including authentication and authorization using packages such as jwt-go or go-oauth2, and utilizing the cache system for quick data retrieval
- [ ]    Utilize Go's concurrency features such as goroutines and channels to handle multiple requests simultaneously, improving API performance
- [ ]    Implement middleware for common tasks such as logging, error handling, and monitoring
- [ ]    Use Gomock for testing by mocking dependencies and testing code in isolation
- [ ]    Implement logging to track events and errors using a package such as logrus or zap
- [ ]    Implement monitoring to track performance and health of the API using packages such as prometheus or go-metrics
- [ ]    Use Go Fuzz for integration testing by providing randomized input to detect and prevent potential bugs and crashes
- [ ]    Use Redis cache data for integration testing to efficiently test the API's integration with the GraphDB and provide more realistic testing scenarios
- [ ]    Implement data migration to handle changes in the data model using a migration tool such as goose or gormigrate
- [ ]    Implement session management to track user sessions and handle authentication and authorization using a library such as gorilla/sessions or context-session
- [ ]    Consider containerization using Docker for easy management and scaling of the service
- [ ]    Consider horizontal scaling by adding more instances of the service behind a load balancer, using a container orchestration platform such as Kubernetes.

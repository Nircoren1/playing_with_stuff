module myapi

go 1.18

require (
	github.com/gin-gonic/gin v1.7.7
	// Add any other dependencies your API requires here
	// For example, if you're using PostgreSQL:
	github.com/lib/pq v1.10.2
)

require github.com/gorilla/mux v1.8.1

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
)

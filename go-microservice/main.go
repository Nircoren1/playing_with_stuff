package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
    connectionString := 
        "user=" + user + " " +
        "password=" + password + " " +
        "dbname=" + dbname + " sslmode=disable"

    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
    a.initializeRoutes()

    // cors
    a.Router.Use(handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    ))
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/test", a.test).Methods("GET")
    a.Router.HandleFunc("/test1", a.test).Methods("POST")

}

func (a *App) test(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"message": "Test successful"}
    respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
    a := App{}
    a.Initialize("user", "password", "dbname")
    a.Run(":8080")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	// Router library
	"github.com/gorilla/mux"
)

type server struct{}

/*
// Remember the Handler interface!
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "POST called"}`))
}

func patch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "PATCH called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "DELETE called"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	userId := -1
	var err error
	if val, ok := pathParams["userId"]; ok {
		userId, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Need a number"}`))
			return
		}
	}
	commentId := -1
	if val, ok := pathParams["commentId"]; ok {
		commentId, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Need a number"}`))
			return
		}
	}
	query := r.URL.Query()
	location := query.Get("location")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"userId":%d,  "commentId":%d, "location": "%s" }`, userId, commentId, location)))
}

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", post).Methods(http.MethodPost)
	api.HandleFunc("/", patch).Methods(http.MethodPatch)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)
	api.HandleFunc("/user/{userId}/comment/{commentId}", params).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}

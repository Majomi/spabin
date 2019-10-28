//go:generate go run spa/generate.go

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/majomi/spabin/spa"
)

func main() {
	r := chi.NewRouter()

	// Return index.html if path does not begin with /api
	r.Handle("/*", spa.Handler{})

	// Api Routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/todos", getTodos)
	})

	// Server listening on port 3000 with 15s timeout
	srv := http.Server{
		Handler:      r,
		Addr:         ":3000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func getTodos(w http.ResponseWriter, r *http.Request) {

	type todo struct {
		Title   string `json:"title"`
		Message string `json:"msg"`
	}

	todos := []*todo{
		&todo{
			Title:   "Todo 1",
			Message: "Message 1",
		},
		&todo{
			Title:   "Todo 2",
			Message: "Message 2",
		},
		&todo{
			Title:   "Todo 3",
			Message: "Message 3",
		},
		&todo{
			Title:   "Todo 4",
			Message: "Message 4",
		},
		&todo{
			Title:   "Todo 5",
			Message: "Message 5",
		},
	}

	render.JSON(w, r, todos)
}

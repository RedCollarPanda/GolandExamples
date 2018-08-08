package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type limitHandler struct {
	connc   chan struct{}
	handler http.Handler
}

func NewLimitHandler(maxConns int, handler http.Handler) http.Handler {
	h := &limitHandler{
		connc:   make(chan struct{}, maxConns),
		handler: handler,
	}
	for i := 0; i < maxConns; i++ {
		h.connc <- struct{}{}
	}
	return h
}

func newHandlerImpl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("incomming connection")

		panic("this must me recovered")
		w.Write([]byte("stub"))
	})
}

func (h *limitHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	select {
	case <-h.connc:

		h.handler.ServeHTTP(w, req)
		h.connc <- struct{}{}

	default:
		http.Error(w, "503 too busy", http.StatusServiceUnavailable)
	}
}

func panicMiddlewareHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("panicMiddleware", r.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	handlerImpl := NewLimitHandler(10, newHandlerImpl())

	panicWrapper := panicMiddlewareHttp(handlerImpl)

	mux.Handle("/panic", panicWrapper)
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println(server.ListenAndServe())

}

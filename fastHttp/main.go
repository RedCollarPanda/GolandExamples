package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

type limitHandler struct {
	connc chan struct{}
}

func NewFastHandler(maxConns int) limitHandler {
	h := limitHandler{
		connc: make(chan struct{}, maxConns),
	}
	for i := 0; i < maxConns; i++ {
		h.connc <- struct{}{}
	}
	return h
}

func panicMiddleware(next limitHandler) func(ctx *fasthttp.RequestCtx) {

	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println("panicMiddleware")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)

				ctx.Response.SetStatusCode(http.StatusServiceUnavailable)
				ctx.Response.BodyWriter().Write([]byte("500 panic"))
			}
		}()

		next.HandleFastHTTP(ctx)
	}
}

func panicPage(ctx *fasthttp.RequestCtx) {
	panic("this must me recovered")
}

func simplePage(ctx *fasthttp.RequestCtx) {
	ctx.Response.BodyWriter().Write([]byte("stub"))
}

func (h *limitHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {

	select {
	case <-h.connc:
		fmt.Println("ServeHTTP")
		switch string(ctx.Path()) {

		case "/messages":
			simplePage(ctx)
		case "/panic":
			panicPage(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)

		}

		h.connc <- struct{}{}
	default:
		ctx.Response.SetStatusCode(http.StatusNotFound)
		ctx.Response.BodyWriter().Write([]byte("404 not found"))
	}
}

func main() {
	handlr := NewFastHandler(100)

	panicWrapper := panicMiddleware(handlr)
	fasthttp.ListenAndServe(":8080", panicWrapper)
}

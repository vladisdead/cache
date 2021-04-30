package server

import (
	"github.com/fasthttp/router"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"

	"time"
)

func (s *Server) newRouter() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Accept,Origin,Accept-Encoding,DNT,User-Agent,Content-Type")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH")

		var (
			path = string(ctx.Path())
			method = string(ctx.Method())
			start = time.Now()
		)

		handlerLogger := s.log.With().Time("received_time", start).
			Str("method", method).
			Str("url", path).
			Str("agent", string(ctx.Request.Header.UserAgent())).
			Str("server_ip", ctx.LocalAddr().String()).
			Logger()


		r := router.New()
		s.newRouterAPI(r, &handlerLogger)

		r.Handler(ctx)
	}
}

func (s *Server) newRouterAPI(r *router.Router, log *zerolog.Logger) {
	r.GET("/user", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetUsersWithMinAgeHandler(ctx, log)
	}))

	r.GET("/users", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetAllUsersHandler(ctx, log)
	}))

	r.GET("/users/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetUserByIdHandler(ctx, log)
	}))
	r.POST("/users", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.CreateUserHandler(ctx, log)
	}))
	r.PUT("/users/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.UpdateUserHandler(ctx, log)
	}))
	r.DELETE("/users/{id}", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.DeleteUserHandler(ctx, log)
	}))

	r.GET("/users/cache", fasthttp.CompressHandler(func(ctx *fasthttp.RequestCtx) {
		s.GetAllCacheHandler(ctx, log)
	}))

	r.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.Response.SetBodyString(`"` + string(ctx.Path()) + `"`)
		ctx.Response.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	}
}


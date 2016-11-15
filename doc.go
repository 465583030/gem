// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

/*
Package gem is a simple and fast web framework,
for building web or restful application.

Install

Run the following command to install Gem:

	go get github.com/go-gem/gem


Example

Run the following code, and then navigate to http://127.0.0.1:8080.

	package main
	import (
		"log"

		"github.com/go-gem/gem"
		"github.com/valyala/fasthttp"
	)

	func main() {
		server := gem.New()

		router := gem.NewRouter()

		router.GET("/", func(c *gem.Context) {
			c.HTML(fasthttp.StatusOK, "Hello world.")
		})

		log.Fatal(server.ListenAndServe(":8080", router.Handler))
	}


Server

The Server is extend edition of fasthttp.Server.

1. Create Server:

	server := gem.New()

2. Create router:

	router := gem.NewRouter()

3. Launch the server:

	// HTTP server.
	server.ListenAndServe(":8080", router.Handler)

Router

1. Create router's instance:

	router := gem.NewRouter()

2. Register Middlewares:

	router.Use(middleware)
	...

3. Register handler:

	// Handle registers a new request handle with the given path and method.
	router.Handle("GET", "/", func(c *gem.Context) {
		c.Write([]byte("Hello world."))
	})

	// Router also provides some shortcuts: GET, POST, DELETE, PUT...
	router.GET("/get", func(c *gem.Context) {
		c.Write([]byte("Hello world."))
	})

4. Register specific middlewares:

All of router's APIs can registers specific middlewares through the last parameters.

	router.GET("/specific-middleware", func(c *gem.Context) {
		c.Write([]byte("Hello world."))
	}, middlewareOne, middlewareTwo ...)

5. Static resource files:

	router.ServeFiles("/js/*filepath", "/path/to/js")

6. Router's parameters.

	router.GET("/user/:name", func(c *gem.Context) {
		c.HTML(200, fmt.Sprintf("Hello %s", c.UserValue("name")))
	})

If you need to handle CORS, pass your CORS middleware as the third parameter like the following:

	router.ServeFiles("/js/*filepath", "/path/to/js", corsMiddleware)

Context

Context is extend edition of fasthttp.RequestCtx.

Context provides some convenient methods:

	HTML(code int, body string)

	JSON(code int, v interface{})

	JSONP(code int, v interface{}, callback []byte)

	XML(code int, v interface{}, headers ...string)

	IsAjax() bool

Middleware

Middleware is an useful feature, you can use it to implement some useful functions,
such as BasicAuth, Gzip compress, request body limit, IP white list or blacklist etc.

It is easy to write a middleware, you just need to implement the Handle method:

	Handle(next Handler) Handler

Built-in middlewares:

1. Compress(Gzip) Middleware.

2. ...

Logger

Gem defines Logger's interface, it is easy to choose your favorite logging package.

Gem provide a leveled logging package, see https://github.com/go-gem/log for more details.

How to set logger:

	// Create logger instance.
	logger := log.New(os.Stderr, log.LstdFlags, log.LevelAll)

	// Set server's logger.
	server.SetLogger(logger)

	// Using logger.
	router.Get("/log", func(c *gem.Context) {
		c.Logger().Info("log info")
	})

Logger interface following:

	type Logger interface {
		Debug(v ...interface{})
		Debugf(format string, v ...interface{})
		Debugln(v ...interface{})

		Info(v ...interface{})
		Infof(format string, v ...interface{})
		Infoln(v ...interface{})

		Warning(v ...interface{})
		Warningf(format string, v ...interface{})
		Warningln(v ...interface{})

		Error(v ...interface{})
		Errorf(format string, v ...interface{})
		Errorln(v ...interface{})

		Fatal(v ...interface{})
		Fatalf(format string, v ...interface{})
		Fatalln(v ...interface{})
	}


Sessions Store

See https://github.com/go-gem/sessions for more details.

How to set sessions store:

	// Create sessions store instance.
	store := sessions.NewCookieStore([]byte("something-very-secret"))

	// Set server's sessions store.
	server.SetSessionsStore(store)

	// Using sessions store.
	router.Get("/log", func(c *gem.Context) {
		session,err := c.SessionsStore().Get(c.RequestCtx, "GOSESSION")
		if err != nil {
			...
		}
		...
	})
*/
package gem

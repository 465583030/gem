// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package middleware

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/go-gem/gem"
)

// Default configuration.
var (
	AllowOrigins = []string{"*"}
	AllowMethods = []string{gem.StrGet, gem.StrHead, gem.StrPut, gem.StrPatch, gem.StrPost, gem.StrDelete}
)

// Cross-Origin Resource Sharing middleware.
type CORS struct {
	//  Skipper defines a function to skip middleware.
	Skipper Skipper

	// Access-Control-Allow-Origin
	AllowOrigins []string

	// Access-Control-Allow-Methods
	AllowMethods []string

	// Access-Control-Allow-Headers
	AllowHeaders []string

	// Access-Control-Expose-Headers
	ExposeHeaders []string

	// Access-Control-Max-Age
	MaxAge int

	// Access-Control-Allow-Credentials
	AllowCredentials bool
}

// NewCORS returns CORS instance with the default configuration.
func NewCORS() *CORS {
	return &CORS{
		Skipper:       defaultSkipper,
		AllowOrigins:  AllowOrigins,
		AllowMethods:  AllowMethods,
		AllowHeaders:  []string{},
		ExposeHeaders: []string{},
	}
}

// Handle implements Middleware's Handle function.
func (cors *CORS) Handle(next gem.Handler) gem.Handler {
	if cors.Skipper == nil {
		cors.Skipper = defaultSkipper
	}
	if len(cors.AllowOrigins) == 0 {
		cors.AllowOrigins = AllowOrigins
	}
	if len(cors.AllowMethods) == 0 {
		cors.AllowMethods = AllowMethods
	}
	allowMethods := strings.Join(cors.AllowMethods, ",")
	allowHeaders := strings.Join(cors.AllowHeaders, ",")
	exposeHeaders := strings.Join(cors.ExposeHeaders, ",")
	maxAge := strconv.Itoa(cors.MaxAge)

	return gem.HandlerFunc(func(ctx *gem.Context) {
		if cors.Skipper(ctx) {
			next.Handle(ctx)
			return
		}

		next.Handle(ctx)

		origin := string(ctx.RequestCtx.Request.Header.Peek(gem.HeaderOrigin))

		allowedOrigin := ""
		for _, o := range cors.AllowOrigins {
			if o == "*" || o == origin {
				allowedOrigin = o
				break
			}
		}

		// Simple request
		if bytes.Equal(ctx.RequestCtx.Request.Header.Method(), gem.MethodOptions) {
			ctx.RequestCtx.Response.Header.Add(gem.HeaderVary, gem.HeaderOrigin)
			if origin == "" || allowedOrigin == "" {
				next.Handle(ctx)
				return
			}
			ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlAllowOrigin, allowedOrigin)
			if cors.AllowCredentials {
				ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlAllowCredentials, "true")
			}
			if exposeHeaders != "" {
				ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlExposeHeaders, exposeHeaders)
			}
			next.Handle(ctx)
			return
		}

		// Preflight request
		ctx.RequestCtx.Response.Header.Add(gem.HeaderVary, gem.HeaderOrigin)
		ctx.RequestCtx.Response.Header.Add(gem.HeaderVary, gem.HeaderAccessControlRequestMethod)
		ctx.RequestCtx.Response.Header.Add(gem.HeaderVary, gem.HeaderAccessControlRequestHeaders)
		if origin == "" || allowedOrigin == "" {
			next.Handle(ctx)
			return
		}
		ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlAllowOrigin, allowedOrigin)
		ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlAllowMethods, allowMethods)
		if cors.AllowCredentials {
			ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlAllowCredentials, "true")
		}
		if allowHeaders != "" {
			ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlAllowHeaders, allowHeaders)
		} else {
			h := ctx.RequestCtx.Response.Header.Peek(gem.HeaderAccessControlRequestHeaders)
			if len(h) > 0 {
				ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlAllowHeaders, string(h))
			}
		}
		if cors.MaxAge > 0 {
			ctx.RequestCtx.Response.Header.Set(gem.HeaderAccessControlMaxAge, maxAge)
		}

		ctx.Response.ResetBody()
	})
}

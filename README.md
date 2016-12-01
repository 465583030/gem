# Gem [![GoDoc](https://godoc.org/github.com/go-gem/gem?status.svg)](https://godoc.org/github.com/go-gem/gem) [![Build Status](https://travis-ci.org/go-gem/gem.svg?branch=master)](https://travis-ci.org/go-gem/gem) [![Go Report Card](https://goreportcard.com/badge/github.com/go-gem/gem)](https://goreportcard.com/report/github.com/go-gem/gem) [![Coverage Status](https://coveralls.io/repos/github/go-gem/gem/badge.svg?branch=master)](https://coveralls.io/github/go-gem/gem?branch=master) [![Join the chat at https://gitter.im/go-gem/gem](https://badges.gitter.im/go-gem/gem.svg)](https://gitter.im/go-gem/gem?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Gem, a simple and fast web framework, it built top of [fasthttp](https://github.com/valyala/fasthttp).

#### The API is currently unstable until the version `v1.0.0` being released, see [milestones](https://github.com/go-gem/gem/milestones) for more details.

## Install

```
go get github.com/go-gem/gem
```


## Features

- [Graceful shutdown and restart](https://github.com/go-gem/gem/wiki/Graceful-shutdown-and-restart)
- [Leveled logging](https://github.com/go-gem/gem/wiki/Logger)
- High-performance and pretty router, very friendly to RESTful APIs
- Sessions support
- Various Middlewares:
    - JSON WEB TOKEN middleware
    - Compress middleware
    - Basic Auth middleware
    - Request Body Limit middleware
    - CSRF middleware
    - CORS middleware

## Performance

![Gem benchmark](benchmark.png)

| Parameter |                       Info                       |
|:----------|:-------------------------------------------------|
| Date      | 2016-11-01                                       |
| Go Version| 1.7.1                                            |
| OS        | Ubuntu 16.04.1 LTS                               |
| CPU       | Intel(R) Core(TM) i7-4720HQ CPU @ 2.60GHz 8 Cores|
| Memory    | 8G                                               |

The benchmark was tested at my laptop, I will send a pull request to 
[go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) when the first stable version was released.

## Quick Start

```
package main

import (
    "log"
    
    "github.com/go-gem/gem"
)

func main() {
	log.Fatal(gem.ListenAndServe(":8080", func(ctx *gem.Context) {
        ctx.HTML(200, "Hello world.")
    }))
}
```

Run the code above, and then navigate to [127.0.0.1:8080](http://127.0.0.1:8080).

## [Examples](https://github.com/go-gem/examples)

- [**Graceful shutdown and restart**](https://github.com/go-gem/examples/tree/master/graceful)
- [**Leveled logging**](https://github.com/go-gem/examples/tree/master/logger)


## [Wiki](https://github.com/go-gem/gem/wiki)

- [Router](https://github.com/go-gem/gem/wiki/Router)

- [Logger](https://github.com/go-gem/gem/wiki/Logger)

- [Graceful shutdown and restart](https://github.com/go-gem/gem/wiki/Graceful-shutdown-and-restart)

- [Middleware](https://github.com/go-gem/gem/wiki/Middleware)


## Semantic Versioning

Gem follows [semantic versioning 2.0.0](http://semver.org/) managed through GitHub releases.


## Support Us

- :star: the project.
- Spread the word.
- [Contribute](#contribute) to the project.


## Contribute

- [Report issues](https://github.com/go-gem/gem/issues/new)
- Send PRs.
- Improve/fix documentation.

**We’re always looking for help, so if you would like to contribute, we’d love to have you!**


## Motivation

Just for figuring out the web framework's workflow, and try to design a simple and strong web framework.


## Related Packages

- [**tests**](https://github.com/go-gem/tests) - a tests package for fasthttp and Gem.
- [**sessions**](https://github.com/go-gem/sessions) - a sessions manager package for fasthttp.
- [**log**](https://github.com/go-gem/log) - a simple, leveled logging package.


## FAQ

1. What should I pay attention to?

    **At present**, Gem is incompatible with some APIs of fasthttp, the incompatible APIs following:
    
    - **fasthttp.TimeoutHandler**

2. What is the difference between of fasthttp and gem?

    **Gem built on top of fasthttp, and use `Server.ServeConn` to serve connections, it can not works without fasthttp**.
    
    **Advantages**: 
    
    Gem provides some useful built-in features, such as:
    
    - High-performance router
    - Leveled logger
    - Various middlewares
    - Sessions support
    - Graceful shutdown and restart
    
    **Disadvantages**: 
    
    at present, Gem dose not provide APIs to Serve the custom `Listener` like `fasthttp.Serve`,
    but that will not be a problem, We will add support for this feature in the future.


## LICENSE

MIT licensed. See [LICENSE](LICENSE) file for more information.

**Inspiration & Credits**

I have read the code of the following open source projects, and integrate their designs into this project.

I respect these projects and it's authors, and follow their LICENSE.

If your LICENSE is missing, please contact me, I will add it ASAP.

- [**fasthttp**](https://github.com/valyala/fasthttp) - [LICENSE](https://github.com/valyala/fasthttp/blob/master/LICENSE)
- [**httprouter**](https://github.com/julienschmidt/httprouter) - [LICENSE](https://github.com/julienschmidt/httprouter/blob/master/LICENSE)
- [**fasthttprouter**](https://github.com/buaazp/fasthttprouter) - [LICENSE](https://github.com/buaazp/fasthttprouter/blob/master/LICENSE)
- [**echo**](https://github.com/labstack/echo) - [LICENSE](https://github.com/labstack/echo/blob/master/LICENSE)
- [**endless**](https://github.com/fvbock/endless) - [LICENSE](https://github.com/fvbock/endless/blob/master/LICENSE)
- [**go-graceful-restart-example**](https://github.com/Scalingo/go-graceful-restart-example) - [LICENSE](https://github.com/Scalingo/go-graceful-restart-example/blob/master/LICENSE)
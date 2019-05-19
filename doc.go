// Copyright 2015 The Tango Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tango is a micro & pluggable web framework for Go language.

// 	package main

// 	import "gitea.com/lunny/tango"

// 	type Action struct {
// 	}

// 	func (Action) Get() string {
// 	    return "Hello tango!"
// 	}

// 	func main() {
// 	    t := tango.Classic()
// 	    t.Get("/", new(Action))
// 	    t.Run()
// 	}

// Middlewares allow you easily plugin/unplugin features for your Tango applications.

// There are already many [middlewares](https://gitea.com/tango) to simplify your work:

// - recovery - recover after panic
// - compress - Gzip & Deflate compression
// - static - Serves static files
// - logger - Log the request & inject Logger to action struct
// - param - get the router parameters
// - return - Handle the returned value smartlly
// - ctx - Inject context to action struct

// - [session](https://gitea.com/tango/session) - Session manager, with stores support:
//   * Memory - memory as a session store
//   * [Redis](https://gitea.com/tango/session-redis) - redis server as a session store
//   * [nodb](https://gitea.com/tango/session-nodb) - nodb as a session store
//   * [ledis](https://gitea.com/tango/session-ledis) - ledis server as a session store)
// - [xsrf](https://gitea.com/tango/xsrf) - Generates and validates csrf tokens
// - [binding](https://gitea.com/tango/binding) - Bind and validates forms
// - [renders](https://gitea.com/tango/renders) - Go template engine
// - [dispatch](https://gitea.com/tango/dispatch) - Multiple Application support on one server
// - [tpongo2](https://gitea.com/tango/tpongo2) - Pongo2 teamplte engine support
// - [captcha](https://gitea.com/tango/captcha) - Captcha
// - [events](https://gitea.com/tango/events) - Before and After
// - [flash](https://gitea.com/tango/flash) - Share data between requests
// - [debug](https://gitea.com/tango/debug) - Show detail debug infomaton on log

package tango

Tango [![Build Status](https://drone.gitea.com/api/badges/lunny/tango/status.svg)](https://drone.gitea.com/lunny/tango)  [![](http://gocover.io/_badge/gitea.com/lunny/tango)](http://gocover.io/gitea.com/lunny/tango)
[![](https://goreportcard.com/badge/gitea.com/lunny/tango)](https://goreportcard.com/report/gitea.com/lunny/tango)
[![Join the chat at https://img.shields.io/discord/323705316027924491.svg](https://img.shields.io/discord/323705316027924491.svg)](https://discord.gg/7Ckxjwu)
[English](README.md)
=======================

![Tango Logo](logo.png)

Tango 是一个微内核的Go语言Web框架，采用模块化和注入式的设计理念。开发者可根据自身业务逻辑来选择性的装卸框架的功能，甚至利用丰富的中间件来搭建一个全栈式Web开发框架。

## 特性
- 强大而灵活的路由设计
- 兼容已有的 `http.Handler`
- 基于中间件的模块化设计，灵活定制框架功能
- 高性能的依赖注入方式

## 安装Tango：

    go get gitea.com/lunny/tango

## 快速入门

一个经典的Tango例子如下：

```go
package main

import (
    "errors"

    "gitea.com/lunny/tango"
)

type Action struct {
    tango.JSON
}

func (Action) Get() interface{} {
    if true {
        return map[string]string{
            "say": "Hello tango!",
        }
    }
    return errors.New("something error")
}

func main() {
    t := tango.Classic()
    t.Get("/", new(Action))
    t.Run()
}
```

然后在浏览器访问`http://localhost:8000`, 将会得到一个json返回

```
{"say":"Hello tango!"}
```

如果将上述例子中的 `true` 改为 `false`, 将会得到一个json返回

```
{"err":"something error"}
```

这段代码因为拥有一个内嵌的`tango.JSON`，所以返回值会被自动的转成Json

## 文档

- [Manual](http://gobook.io/read/gitea.com/tango/manual-en-US/), And you are welcome to contribue for the book by git PR to [gitea.com/tango/manual-en-US](https://gitea.com/tango/manual-en-US)
- [操作手册](http://gobook.io/read/gitea.com/tango/manual-zh-CN/)，您也可以访问 [gitea.com/tango/manual-zh-CN](https://gitea.com/tango/manual-zh-CN)为本手册进行贡献
- [API Reference](https://godoc.org/gitea.com/lunny/tango)

## 交流讨论

- QQ群：369240307
- [论坛](https://groups.google.com/forum/#!forum/go-tango)

## 使用案例

- [会计人论坛](https://www.kuaijiren.com) - 会计人论坛
- [GopherTC](https://github.com/jimmykuu/gopher/tree/2.0) - Golang China
- [Wego](https://gitea.com/tango/wego)  tango结合[xorm](http://www.xorm.io/)开发的论坛
- [Pugo](https://github.com/go-xiaohei/pugo) 博客
- [DBWeb](https://github.com/go-xorm/dbweb) 基于Web的数据库管理工具
- [Godaily](http://godaily.org) - [github](https://github.com/godaily/news) RSS聚合工具
- [Gos](https://gitea.com/tango/gos)  简易的Web静态文件服务端
- [GoFtpd](https://github.com/goftp/ftpd) - 纯Go的跨平台FTP服务器

## 中间件列表

[中间件](https://gitea.com/tango)可以重用代码并且简化工作：

- [recovery](https://gitea.com/lunny/tango/wiki/Recovery) - recover after panic
- [compress](https://gitea.com/lunny/tango/wiki/Compress) - Gzip & Deflate compression
- [static](https://gitea.com/lunny/tango/wiki/Static) - Serves static files
- [logger](https://gitea.com/lunny/tango/wiki/Logger) - Log the request & inject Logger to action struct
- [param](https://gitea.com/lunny/tango/wiki/Params) - get the router parameters
- [return](https://gitea.com/lunny/tango/wiki/Return) - Handle the returned value smartlly
- [context](https://gitea.com/lunny/tango/wiki/Context) - Inject context to action struct
- [session](https://gitea.com/tango/session) - [![CircleCI](https://circleci.com/gh/tango-contrib/session/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/session/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/session/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/session) Session manager, [session-redis](http://gitea.com/tango/session-redis), [session-nodb](http://gitea.com/tango/session-nodb), [session-ledis](http://gitea.com/tango/session-ledis), [session-ssdb](http://gitea.com/tango/session-ssdb)
- [xsrf](https://gitea.com/tango/xsrf) - [![CircleCI](https://circleci.com/gh/tango-contrib/xsrf/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/xsrf/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/xsrf/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/xsrf) Generates and validates csrf tokens
- [binding](https://gitea.com/tango/binding) - [![CircleCI](https://circleci.com/gh/tango-contrib/binding/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/binding/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/binding/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/binding) Bind and validates forms
- [renders](https://gitea.com/tango/renders) - [![CircleCI](https://circleci.com/gh/tango-contrib/renders/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/renders/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/renders/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/renders) Go template engine
- [dispatch](https://gitea.com/tango/dispatch) - [![CircleCI](https://circleci.com/gh/tango-contrib/dispatch/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/dispatch/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/dispatch/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/dispatch) Multiple Application support on one server
- [tpongo2](https://gitea.com/tango/tpongo2) - [![CircleCI](https://circleci.com/gh/tango-contrib/tpongo2/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/tpongo2/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/tpongo2/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/tpongo2) [Pongo2](https://github.com/flosch/pongo2) teamplte engine support
- [captcha](https://gitea.com/tango/captcha) - [![CircleCI](https://circleci.com/gh/tango-contrib/captcha/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/captcha/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/captcha/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/captcha) Captcha
- [events](https://gitea.com/tango/events) - [![CircleCI](https://circleci.com/gh/tango-contrib/events/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/events/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/events/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/events) Before and After
- [flash](https://gitea.com/tango/flash) - [![CircleCI](https://circleci.com/gh/tango-contrib/flash/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/flash/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/flash/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/flash) Share data between requests
- [debug](https://gitea.com/tango/debug) - [![CircleCI](https://circleci.com/gh/tango-contrib/debug/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/debug/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/debug/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/debug) show detail debug infomaton on log
- [basicauth](https://gitea.com/tango/basicauth) - [![CircleCI](https://circleci.com/gh/tango-contrib/basicauth/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/basicauth/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/basicauth/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/basicauth) basicauth middleware
- [authz](https://gitea.com/tango/authz) - [![Build Status](https://travis-ci.org/tango-contrib/authz.svg?branch=master)](https://travis-ci.org/tango-contrib/authz) [![Coverage Status](https://coveralls.io/repos/github/tango-contrib/authz/badge.svg?branch=master)](https://coveralls.io/github/tango-contrib/authz?branch=master) manage permissions via ACL, RBAC, ABAC
- [cache](https://gitea.com/tango/cache) - [![CircleCI](https://circleci.com/gh/tango-contrib/cache/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/cache/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/cache/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/cache) cache middleware - cache-memory, cache-file, [cache-ledis](https://gitea.com/tango/cache-ledis), [cache-nodb](https://gitea.com/tango/cache-nodb), [cache-mysql](https://gitea.com/tango/cache-mysql), [cache-postgres](https://gitea.com/tango/cache-postgres), [cache-memcache](https://gitea.com/tango/cache-memcache), [cache-redis](https://gitea.com/tango/cache-redis)
- [rbac](https://gitea.com/tango/rbac) - [![CircleCI](https://circleci.com/gh/tango-contrib/rbac/tree/master.svg?style=svg)](https://circleci.com/gh/tango-contrib/rbac/tree/master) [![codecov](https://codecov.io/gh/tango-contrib/rbac/branch/master/graph/badge.svg)](https://codecov.io/gh/tango-contrib/rbac) rbac control

## License
This project is under BSD License. See the [LICENSE](LICENSE) file for the full license text.

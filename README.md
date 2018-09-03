<p align="center"><img src="http://hprose.com/banner.@2x.png" alt="Hprose" title="Hprose" width="650" height="200" /></p>

# [Hprose-gateway-etcd](https://github.com/vlorc/hprose-gateway-etcd)
[简体中文](https://github.com/vlorc/hprose-gateway-etcd/blob/master/README_CN.md)

[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![codebeat badge](https://codebeat.co/badges/c41b426c-4121-4dc8-99c2-f1b60574be64)](https://codebeat.co/projects/github-com-vlorc-hprose-gateway-etcd-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vlorc/hprose-gateway-etcd)](https://goreportcard.com/report/github.com/vlorc/hprose-gateway-etcd)
[![GoDoc](https://godoc.org/github.com/vlorc/hprose-gateway-etcd?status.svg)](https://godoc.org/github.com/vlorc/hprose-gateway-etcd)
[![Build Status](https://travis-ci.org/vlorc/hprose-gateway-etcd.svg?branch=master)](https://travis-ci.org/vlorc/hprose-gateway-etcd?branch=master)
[![Coverage Status](https://coveralls.io/repos/github/vlorc/hprose-gateway-etcd/badge.svg?branch=master)](https://coveralls.io/github/vlorc/hprose-gateway-etcd?branch=master)

Hprose gateway etcd service discovery based on golang

## Features
+ lazy client
+ discovery
+ register
+ water

## Installing
	go get github.com/vlorc/hprose-gateway-etcd

## Quick Start

* Service discovery
```golang
r := resolver.NewResolver(cli, ctx, "rpc" /*prefix*/)
// print event 
go r.Watch("*", watcher.NewPrintWatcher(fmt.Printf))
```

* Service register
```golang
m := manager.NewManager(cli, context.Background(), "rpc" /*prefix*/, 5 /*ttl*/)
s := manage.Register("user" /*name*/, "1" /*uuid*/)
s.Update(&types.Service{
	Id:       "1",
	Name:     "user",
	Url:      "http://localhost:8080",
})
```
## License
This project is under the apache License. See the LICENSE file for the full license text.


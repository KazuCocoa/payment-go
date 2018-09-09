learn by https://github.com/po3rin/vue-golang-payment-app


```
$ go get -u [some packages]
```

You need set `$GOPATH`. If you use asdf, you probably should specify the path like below in `bash_profile`, for example.

- `/Users/kazu/.asdf/installs/golang/1.9/packages`

---

Run mysql as `brew`

## Dep for version management

https://golang.github.io/dep/docs/installation.html

## MySQL

```
$ brew services start mysql
$ brew services list
```

change password

```
$ mysqladmin -u root password 'yourpassword'
$ mysql -s -u root -p
$ mysql> create database itemsdb;
```

## Run
### Server

```
$ MYSQL_USER=root MYSQL_PASSWORD=mysql MYSQL_DATABASE=itemsDB API_SERVER_PORT=:8888 CLIENT_CORS_ADDR=http://localhost:8888 GIN_MODE=debug go run backend-api/main.go
[GIN-debug] [WARNING] Now Gin requires Go 1.6 or later and Go 1.7 will be required soon.

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/items             --> _/Users/kazuaki/tmp/payment/backend-api/infrastructure.init.0.func1 (4 handlers)
[GIN-debug] GET    /api/v1/items/:id         --> _/Users/kazuaki/tmp/payment/backend-api/infrastructure.init.0.func2 (4 handlers)
[GIN-debug] POST   /api/v1/charge/items/:id  --> _/Users/kazuaki/tmp/payment/backend-api/infrastructure.init.0.func3 (4 handlers)
[GIN-debug] Listening and serving HTTP on :8888
[GIN] 2018/09/09 - 18:17:06 | 200 |    9.002378ms |             ::1 | GET      /api/v1/items/1
[GIN] 2018/09/09 - 18:17:08 | 200 |    1.485424ms |             ::1 | GET      /api/v1/items/1
[GIN] 2018/09/09 - 18:17:24 | 200 |      931.81µs |             ::1 | GET      /api/v1/items
[GIN] 2018/09/09 - 18:17:29 | 403 |   14.419868ms |             ::1 | POST     /api/v1/charge/items/1
```

### Request by curl
```
$ curl -X GET localhost:8888/api/v1/items/1
{"ID":1,"Name":"toy","Description":"test-toy","Amount":2000}

$ curl -X GET localhost:8888/api/v1/items/1
{"ID":1,"Name":"toy","Description":"test-toy","Amount":2000}

$ curl -X GET localhost:8888/api/v1/items
[{"ID":1,"Name":"toy","Description":"test-toy","Amount":2000},{"ID":2,"Name":"game","Description":"test-game","Amount":6000}]

$ curl -X POST localhost:8888/api/v1/charge/items/1
{"code":14,"message":"all SubConns are in TransientFailure, latest connection error: connection error: desc = \"transport: Error while dialing dial tcp [::1]:50051: getsockopt: connection refused\""}
```


# iotserver-apibench
Simple Golang concurrency example showing how to use goroutine, channel and WaitGroup
to benchmark with server API.

# Signup API

```
POST /user/signup HTTP/1.1
Host: localhost:3000
Content-Type: application/x-www-form-urlencoded
Cache-Control: no-cache
Postman-Token: 19437a39-cb19-e846-331c-5894019d3db9

username=demo&email=demo%40email.com&password=1234qwer
```


# Prerequesites

* Golang (1.8+ recommended)

# Build

```sh
go build -o iotserver-apibench
```

# Run

```sh
➜  iotserver-apibench ./iotserver-apibench
2017/03/06 16:01:44 user868 OK
2017/03/06 16:01:44 user869 OK
2017/03/06 16:01:44 user860 OK
2017/03/06 16:01:44 user865 OK
2017/03/06 16:01:44 user861 OK
2017/03/06 16:01:44 user867 OK
2017/03/06 16:01:44 user863 OK
2017/03/06 16:01:44 user864 OK
2017/03/06 16:01:44 user866 OK
Time: 0.620937s#
```

# page-views-counter
Golang page views counter

Usage:

1. Download and Install ```go get github.com/gksbrandon/page-view-counter```
2. Change directory to the root path ```cd $GOPATH/src/github.com/gksbrandon/page-views-counter```
3. Run the API ```go run main.go```

You should see a message, stating server is running on localhost:8888
The server connects to a postgresql database on elephansql

Now you can test a payload with Postman against the two endpoints:
``` POST http://localhost:8888/counter/v1/statistics```
```➜  ~ curl -H "Content-Type: application/json" -X POST -d '{"data": {"article_id": "abc"}}' http://localhost:8888/counter/v1/statistics -v
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8888 (#0)
> POST /counter/v1/statistics HTTP/1.1
> Host: localhost:8888
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 31
>
* upload completely sent off: 31 out of 31 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 09 Mar 2020 02:01:01 GMT
< Content-Length: 30
<
{"data":{"article_id":"abc"}}
* Connection #0 to host localhost left intact```

``` GET http://localhost:8888/counter/v1/statistics/article_id/{{article_id}}```
```➜  ~ curl http://localhost:8888/counter/v1/statistics/article_id/abc -v
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8888 (#0)
> GET /counter/v1/statistics/article_id/abc HTTP/1.1
> Host: localhost:8888
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 09 Mar 2020 02:02:10 GMT
< Content-Length: 284
<
{"data":{"article_id":"abc","type":"statistics_article_view_count","attributes":{"count":[{"reference":"5 minutes ago","Count":2},{"reference":"1 hour ago","Count":7},{"reference":"1 day ago","Count":12},{"reference":"2 days ago","Count":12},{"reference":"3 days ago","Count":12}]}}}
* Connection #0 to host localhost left intact
➜  ~ curl http://localhost:8888/counter/v1/statistics/article_id/def -v
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8888 (#0)
> GET /counter/v1/statistics/article_id/def HTTP/1.1
> Host: localhost:8888
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 09 Mar 2020 02:02:17 GMT
< Content-Length: 285
<
{"data":{"article_id":"def","type":"statistics_article_view_count","attributes":{"count":[{"reference":"5 minutes ago","Count":0},{"reference":"1 hour ago","Count":13},{"reference":"1 day ago","Count":18},{"reference":"2 days ago","Count":18},{"reference":"3 days ago","Count":18}]}}}
* Connection #0 to host localhost left intact
➜  ~```
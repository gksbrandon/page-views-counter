# page-views-counter
Golang page views counter

Usage:

1. Download and Install ```go get github.com/gksbrandon/page-view-counter```
2. Change directory to the root path ```cd $GOPATH/src/github.com/gksbrandon/page-views-counter```
3. Run the API ```go run main.go```

You should see a message, stating server is running on localhost:8888
The server connects to a postgresql database on elephansql

Now you can test a payload with Postman against the two endpoints:
```POST http://localhost:8888/counter/v1/statistics```

```curl -H "Content-Type: application/json" -X POST -d '{"data": {"article_id": "abc"}}' http://localhost:8888/counter/v1/statistics -v```
{"data":{"article_id":"abc"}}

```GET http://localhost:8888/counter/v1/statistics/article_id/{{article_id}}```
```➜  ~ curl http://localhost:8888/counter/v1/statistics/article_id/abc -v```
{"data":{"article_id":"abc","type":"statistics_article_view_count","attributes":{"count":[{"reference":"5 minutes ago","Count":2},{"reference":"1 hour ago","Count":7},{"reference":"1 day ago","Count":12},{"reference":"2 days ago","Count":12},{"reference":"3 days ago","Count":12}]}}}

```curl http://localhost:8888/counter/v1/statistics/article_id/def -v```
{"data":{"article_id":"def","type":"statistics_article_view_count","attributes":{"count":[{"reference":"5 minutes ago","Count":0},{"reference":"1 hour ago","Count":13},{"reference":"1 day ago","Count":18},{"reference":"2 days ago","Count":18},{"reference":"3 days ago","Count":18}]}}}
* Connection #0 to host localhost left intact

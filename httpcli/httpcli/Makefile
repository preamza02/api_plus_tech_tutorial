build:
	go build -o bin/httpcli main.go

URL=https://httpbin.org
URL_GET=${URL}/get
URL_POST=${URL}/post
URL_DELETE=${URL}/delete
URL_PUT=${URL}/put
test-get-dev:
	go run main.go get ${URL_GET}
test-get-query-dev:
	go run main.go get ${URL_GET} -query q=golang -query lang=en
test-get-header-dev:
	go run main.go get ${URL_GET} -header Accept-Language=en-US -header Accept-Encoding=gzip
test-get-both-dev:
	go run main.go get ${URL_GET} -query q=golang -header Accept-Language=en-US  

test-post-dev:
	go run main.go post ${URL_POST} -json '{"name":"test","lang":"golang"}'

test-put-dev:
	go run main.go put ${URL_PUT} -json '{"update":"true"}' 

test-delete-dev:
	go run main.go delete ${URL_DELETE} 

test-get:
	httpcli https://www.google.com

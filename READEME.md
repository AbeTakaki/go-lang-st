github.com/AbeTakaki/go-lang-st

1. HTTP サーバー
2. JSON
```
curl.exe  -X GET -w "%{http_code}\n" http://localhost:8080/hello
curl.exe  -X POST -w "%{http_code}\n" http://localhost:8080/article
curl.exe  -X GET -w "%{http_code}\n" http://localhost:8080/article/list
curl.exe  -X GET -w "%{http_code}\n" http://localhost:8080/article/{id}
curl.exe  -X POST -w "%{http_code}\n" http://localhost:8080/article/nice
curl.exe  -X POST -w "%{http_code}\n" http://localhost:8080/comment
```
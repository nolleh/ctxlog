## ctxlog
echo middleware that print log for a api's request / response

### Example 

#### Output

colored by log level

```go
2020-10-08T16:40:46.4543276+09:00 [INFO] {
 "name": "nolleh",
 "say": "hello happy new year!"
}
2020-10-08T16:40:46.5171606+09:00 [TRACE] {
 "request": {
  "body": {
   "name": "nolleh",
   "say": "hello happy new year!"
  },
  "header": {
   "Accept": [
    "*/*"
   ],
   "Accept-Encoding": [
    "gzip, deflate, br"
   ],
   "Cache-Control": [
    "no-cache"
   ],
   "Connection": [
    "keep-alive"
   ],
   "Content-Length": [
    "54"
   ],
   "Content-Type": [
    "application/json"
   ],
   "Postman-Token": [
    "b1556c81-bf31-4e6a-8c2a-e1bd0787ebdd"
   ],
   "User-Agent": [
    "PostmanRuntime/7.22.0"
   ]
  },
  "method": "POST",
  "uri": "/"
 },
 "requestId": "",
 "response": {
  "body": {
   "message": "Hello, World!"
  },
  "status": 200
 }
}

```

#### Source

```go
    // Echo instance
    // e := echo.New()
    
    /// this one line!
    e.Use(middleware.CtxLogger())

    // e.POST("/", hello)
    // e.Logger.Fatal(e.Start(":1323"))
```


### Dependency 
* echo framework (for middlware implementation) [](http://github.com/labstack/echo)
* caption_json_formatter [](http://github.com/nolleh/caption_json_formatter)
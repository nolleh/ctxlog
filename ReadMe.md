## ctxlog
echo middleware that print log for a api's request / response

### Example 

#### Output

```go
2020-01-23T13:33:42.84136+09:00 [TRACE] [null] {
 "request": {
  "body": {
   "name": "nolleh",
   "say": "hello happy new year!"
  },
  "header": {
   "Content-Length": [
    "60"
   ],
   "Content-Type": [
    "application/json"
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
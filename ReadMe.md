## ctxlog
echo middleware that print log for a api's request / response

### Example 

#### Output

```go
2020-01-21T23:48:53.6996825+09:00 [TRACE] [null] {
 "RequestId": "", // todo
 "request": {
  "body": {
   "name": "nolleh",
   "say": "hello, Happy new year!"
  },
  "header": {
   "Content-Length": [
    "41"
   ],
   "Content-Type": [
    "application/json"
   ]
  },
  "method": "POST",
  "uri": "/"
 },
 "response": {
  "body": null, // todo
  "status": 0
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
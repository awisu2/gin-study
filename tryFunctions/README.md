# try gin functions

```bash
go mod tidy
go run . run

# http://localhost:8080/user/123
```

## NOTES

- create ginEngine: `gin.New()`
  - `gin.Default()` with the Logger and Recovery middleware
- middleware: ex: `engine.Use(gin.Logger())`
- router
  - parameter: `g.GET("/show", func(c *gin.Context){c.Param("name")})`
- logger: write later
- validator: [go\-playground/validator](https://github.com/go-playground/validator)
- binding
  - support type JSON, XML, YAML, standard form value(foo=bar&buz=fiz)
- about this codes architecture. write later.

## logger

```go
# set target
f, _ := os.Create(opt.LogFile)
gin.DefaultWriter = io.MultiWriter(f)

# use logger
engine.Use(gin.LoggerWithFormatter(opt.Formatter))

# if not want color
gin.DisableConsoleColor()
```


## about architecture

- [bxcodec/go\-clean\-arch](https://github.com/bxcodec/go-clean-arch) を参考に、usecase毎にpackageを分割しています
  - そのため gin の実装も分割されているのでこれでいいのか迷っています
  - notesなどはその実装にあたり、例えばMailなどのusecaseが増えればそれに合わせて増減する予定です
- wireを利用しているので、Usecaseの追加更新時は、`go generate` または `wire ./app/wire` を実行のこと

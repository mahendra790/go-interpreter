package server

import (
	"bytes"
	"monkey/src/evaluator"
	"monkey/src/lexer"
	"monkey/src/object"
	"monkey/src/parser"
	"net/http"

	"github.com/gin-gonic/gin"

	"flag"
	"log"
)

type ExecuteBody struct {
	Code string `json:"code" binding:"required"`
}

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/execute", func(ctx *gin.Context) {
		var payload = ExecuteBody{}
		ctx.ShouldBindJSON(&payload)

		l := lexer.New(payload.Code)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) > 0 {
			ctx.JSON(200, gin.H{
				"errors": p.Errors(),
			})
			return
		}

		evaluated := evaluator.Eval(program, object.NewEnvironment(), &bytes.Buffer{})

		ctx.JSON(200, gin.H{
			"message": "success",
			"output":  evaluated.Inspect(),
		})
	})
	http.ListenAndServe(":8080", r)
}

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func RunServer() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}

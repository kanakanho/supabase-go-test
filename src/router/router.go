package router

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/supabase-community/postgrest-go"
)

func Init() {
	f, _ := os.Create("../log/server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	client := postgrest.NewClient("http://localhost:5432/rest/v1", "", nil)
	if client.ClientError != nil {
		panic(client.ClientError)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!!")
	})
	r.GET("/add_them", func(c *gin.Context) {
		result := client.Rpc("add_them", "", map[string]int{"a": 12, "b": 3})
		c.JSON(http.StatusOK, result)
	})

	// サーバーの起動状態を表示しながら、ポート8084でサーバーを起動する
	if err := r.Run("0.0.0.0:8000"); err != nil {
		fmt.Println("サーバーの起動に失敗しました:", err)
	} else {
		fmt.Println("サーバーが正常に起動しました")
	}
}

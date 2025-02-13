package main

import (
	"fmt"
	"net/http"
	"time"

	components "github.com/aadielpr/gtl/templates/components"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
		components.Hello("My Website", "World").Render(c, c.Writer)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("Error when starting server:", err)
		return
	}

	fmt.Println("Server running on port: 8080")
}

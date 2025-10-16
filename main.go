package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Stack string `json:"stack"`
}

type result struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	TimeStamp string `json:"timestamp"`
	Fact      string `json:"fact"`
}

func rateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(2, 4)
	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			message := gin.H{
				"status":    "unsuccessful",
				"error":     "The API is at capacity, try again later.",
				"timeStamp": time.Now().UTC().Format(time.RFC3339),
			}
			ctx.JSON(http.StatusTooManyRequests, message)
			ctx.Abort()
			return
		} else {
			ctx.Next()
		}
	}
}

func factHandler(ctx *gin.Context) {
	tojumi := User{
		Email: "odelolatojumi@gmail.com",
		Name:  "Odelola Oluwantojumi",
		Stack: "Go/Gin",
	}

	status := "success"
	catfact, statusCode, err := makeRequest()
	if err != nil || statusCode != http.StatusOK {
		ctx.JSON(statusCode, err)
		ctx.Abort()
	}

	result := result{
		Status:    status,
		User:      tojumi,
		TimeStamp: time.Now().UTC().Format(time.RFC3339),
		Fact:      catfact.CatFact,
	}

	ctx.JSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	router.GET("/me", rateLimiter(), factHandler)
	router.Run(":7070")
}

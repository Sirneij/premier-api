package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type teamSchema struct {
	teamName    []string  `json:"teamName"`
	teamMembers []string  `json:"teamMembers"`
	description string    `json:"description"`
	createdAt   time.Time `json:"createdAt"`
	updatedAt   time.Time `json:"updatedAt"`
}

func main() {
	router := gin.Default()

	router.Run()
}

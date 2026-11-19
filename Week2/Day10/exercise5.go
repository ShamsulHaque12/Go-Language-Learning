package main

/*
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sample User data
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Sujon"},
	{ID: 2, Name: "Rahim"},
}

// APIKeyMiddleware checks for valid X-API-Key header
func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Header থেকে X-API-Key এর মান নেওয়া
		apiKey := c.GetHeader("X-API-Key")

		// API Key অনুপস্থিত (missing) বা ভুল (invalid) হলে 401 Unauthorized দেওয়া হবে
		if apiKey == "" || apiKey != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort() // Request যাতে পরবর্তী handler-এ না যায়, তাই abort করা হলো
			return
		}

		// API Key ঠিক থাকলে পরবর্তী handler-এ যাবে
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// GET /users route-এ APIKeyMiddleware apply করা হলো
	r.GET("/users", APIKeyMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Users fetched successfully",
			"users":   users,
		})
	})

	// Server Run
	r.Run(":8080")
}

*/
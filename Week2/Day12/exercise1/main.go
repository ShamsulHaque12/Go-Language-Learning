package main

/*
import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"go-language-learning/Week2/Day12/exercise1/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := database.CreateUsersTable(db); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		users, err := database.GetUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   users,
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid user ID",
			})
			return
		}

		if id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "User ID must be a positive integer",
			})
			return
		}

		user, err := database.GetUserByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"status": "error",
					"error":  "User not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   user,
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var user database.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request payload: " + err.Error(),
			})
			return
		}

		createdUser, err := database.CreateUser(db, &user)
		if err != nil {
			var pqErr *pq.Error

			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				c.JSON(http.StatusConflict, gin.H{
					"status": "error",
					"error":  "Email already exists",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Failed to create user",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "success",
			"data":   createdUser,
		})
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid user ID",
			})
			return
		}

		var user database.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid request payload: " + err.Error(),
			})
			return
		}

		user.ID = id

		if err := database.UpdateUser(db, &user); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"status": "error",
					"error":  "User not found",
				})
				return
			}

			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				c.JSON(http.StatusConflict, gin.H{
					"status": "error",
					"error":  "Email already exists",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Failed to update user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   user,
		})
	})

	r.PATCH("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid user ID",
			})
			return
		}

		existingUser, err := database.GetUserByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"status": "error",
					"error":  "User not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Failed to fetch user",
			})
			return
		}

		var input struct {
			Name     *string `json:"name"`
			Email    *string `json:"email"`
			Age      *int    `json:"age"`
			IsActive *bool   `json:"is_active"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid request payload: " + err.Error(),
			})
			return
		}

		if input.Name != nil {
			existingUser.Name = *input.Name
		}
		if input.Email != nil {
			existingUser.Email = *input.Email
		}
		if input.Age != nil {
			existingUser.Age = *input.Age
		}
		if input.IsActive != nil {
			existingUser.IsActive = *input.IsActive
		}

		if err := database.UpdateUser(db, existingUser); err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				c.JSON(http.StatusConflict, gin.H{
					"status": "error",
					"error":  "Email already exists",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Failed to update user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   existingUser,
		})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid user ID",
			})
			return
		}

		if err := database.DeleteUser(db, id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{
					"status": "error",
					"error":  "User not found",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Failed to delete user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "User deleted successfully",
		})
	})

	log.Println("Starting server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

*/
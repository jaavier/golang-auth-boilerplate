package handlers

import (
	"app/db"
	"app/middlewares"
	"app/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Private(c echo.Context) error {
	userID := c.Get("user_id").(string)
	return c.JSON(http.StatusOK, echo.Map{"message": "Welcome " + userID})
}

func Signup(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser models.User
	err := db.Users.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return c.JSON(http.StatusConflict, echo.Map{"message": "Username already exists"})
	}
	if err != mongo.ErrNoDocuments {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error checking user existence"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err = db.Users.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error creating user"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "User created successfully"})
}

func Login(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var dbUser models.User
	err := db.Users.FindOne(ctx, bson.M{"username": user.Username}).Decode(&dbUser)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid credentials"})
	}

	token, err := middlewares.CreateJWT(dbUser.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error generating token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func Logout(c echo.Context) error {
	// Handle logout logic if needed (e.g., blacklisting tokens)
	return c.JSON(http.StatusOK, echo.Map{"message": "Logged out successfully"})
}

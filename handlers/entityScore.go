package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/auth/auth"
	"github.com/PoteeDev/scores/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func EntityScore(c *gin.Context) {
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"detail": err.Error()})
		return
	}
	var score models.Score
	col := database.GetCollection(database.DB, "scoreboard")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var entityName string
	if queryName := c.Query("name"); queryName != "" {
		entityName = queryName
	} else {
		entityName = metadata.UserId
	}
	err = col.FindOne(ctx, bson.M{"id": entityName}).Decode(&score)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, score)
}

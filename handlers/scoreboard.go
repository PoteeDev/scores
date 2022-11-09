package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/scores/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func ShowScoreboard(c *gin.Context) {
	var scoreboard []models.Score
	col := database.GetCollection(database.DB, "scoreboard")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find()
	opts.SetSort(bson.M{"place": 1})
	results, err := col.Find(ctx, bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	defer results.Close(ctx)

	if err = results.All(context.TODO(), &scoreboard); err != nil {
		panic(err)
	}

	// for results.Next(ctx) {
	// 	var entityScore models.Score
	// 	if err = results.Decode(&entityScore); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
	// 		return
	// 	}

	// 	scoreboard = append(scoreboard, entityScore)
	// }

	c.JSON(http.StatusOK, gin.H{"scoreboard": scoreboard})
}

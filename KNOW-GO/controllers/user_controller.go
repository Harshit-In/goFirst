package controllers

import (
	"context"
	"fastearn/configs"
	"fastearn/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var user *mongo.Collection = configs.GetCollection(configs.DB, "user")

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user []models.USER
		defer cancel()

		results, _ := user.Find(ctx, bson.M{})

	

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleOrder models.USER
			if err = results.Decode(&singleOrder); err != nil {
				c.JSON( Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()})
			}

			buy_orders = append(user, singleOrder)
		}
		c.JSON( Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data":user})
	}
}

package routes

import (
	"events-api/models"
	"events-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateEvent(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	userId, err := utils.ValidateToken(token)
	fmt.Printf("uid id %v", userId)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized - invalid token"})
		return
	}
	event := models.Event{}
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	event.UserId = userId
	errSave := event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "event created",
			"error":   errSave,
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":   event,
	})
}

func GetEvent(ctx *gin.Context) {
	event_id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cant get data", "message": err})
	}
	event, err := getEvent(event_id)
	fmt.Println("EVENT ", event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "event not exists", "message": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": event})
}

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	fmt.Println("EVENTS: ", events)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant get data", "message": err})
	}
	ctx.JSON(http.StatusOK, events)
}

func UpdateEvent(ctx *gin.Context) {
	event_id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cant get data", "message": err})
	}
	_, err = getEvent(event_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "event not exists", "message": err})
	}
	var event models.Event
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cant parse request", "error": err})
	}
	event.Id = event_id
	err = event.Update()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "cant update event", "error": err})
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "event updates"})
}

func RemoveEvent(ctx *gin.Context) {
	event_id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "event not exist", "message": err})
	}
	event, err := getEvent(event_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "event not exists", "message": err})
	}
	err = event.DeleteEvent()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete event", "message": err})
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "event deleted"})

}

func getEvent(id int64) (*models.Event, error) {
	event, err := models.GetEventById(id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

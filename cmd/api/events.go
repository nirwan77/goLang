package main

import (
	"net/http"
	"res-api-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var events database.Event

	if err := c.ShouldBindJSON(&events); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.Insert(&events)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to the server"})
		return
	}

	c.JSON(http.StatusCreated, events)
}

func (app *application) getAllEvents(c *gin.Context) {
	event, err := app.models.Events.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide an id"})
		return
	}

	event, err := app.models.Events.GetEvent(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event with particular id not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to the server"})
		return
	}

	c.JSON(http.StatusCreated, event)
}

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide an id"})
	}

	event, err := app.models.Events.GetEvent(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event with particular id was not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to the server"})
		return
	}

	updatedEvent := &database.Event{}

	if err := c.BindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedEvent.Id = id

	if err := app.models.Events.UpdateEvent(updatedEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to the server"})
		return
	}

	c.JSON(http.StatusCreated, updatedEvent)
}

func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": "provide the id"})
	}

	if err := app.models.Events.DeleteEvent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
	}

	c.JSON(http.StatusNoContent, nil)
}

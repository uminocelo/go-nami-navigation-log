package controllers

import (
	"nami_navigation_log/helpers"
	"nami_navigation_log/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEntry(cxt *gin.Context) {
	var input models.Entry

	if err := cxt.ShouldBindJSON(&input); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(cxt)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID
	savedEntry, err := input.Save()

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(cxt *gin.Context) {
	user, err := helpers.CurrentUser(cxt)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"data": user.Entries})
}

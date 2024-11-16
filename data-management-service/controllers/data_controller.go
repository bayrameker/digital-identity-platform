package controllers

import (
    "com.digitalidentity.datamanagementservice/models"
    "com.digitalidentity.datamanagementservice/services"
    "com.digitalidentity.datamanagementservice/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CreateData(c *gin.Context) {
    userID := c.GetString("user_id")

    var input models.UserDataInput
    if err := c.ShouldBindJSON(&input); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, err.Error())
        return
    }

    data, err := services.CreateData(userID, input)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create data")
        return
    }

    utils.RespondWithJSON(c, http.StatusCreated, data)
}

func GetData(c *gin.Context) {
    userID := c.GetString("user_id")
    dataID := c.Param("data_id")

    data, err := services.GetData(userID, dataID)
    if err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Data not found")
        return
    }

    utils.RespondWithJSON(c, http.StatusOK, data)
}

// Diğer fonksiyonlar (UpdateData, DeleteData, GrantPermission, GetDataLogs)

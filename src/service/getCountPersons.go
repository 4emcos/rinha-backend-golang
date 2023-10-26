package service

import (
	"rinha-backend/src/infra"

	"github.com/gin-gonic/gin"
)

func GetCountPersons(c *gin.Context) {
	db := infra.ConnectDB()

	count, err := infra.CountPresons(db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, count)

}

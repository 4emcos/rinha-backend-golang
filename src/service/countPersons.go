package service

import (
	"rinha-backend/src/infra"
	"rinha-backend/src/types"

	"github.com/gin-gonic/gin"
)

func CountPersons(c *gin.Context, db types.IPgx) {

	count, err := infra.CountPresons(db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, count)

}

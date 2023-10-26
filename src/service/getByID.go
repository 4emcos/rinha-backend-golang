package service

import (
	"rinha-backend/src/infra"

	"github.com/gin-gonic/gin"
)

func GetByID(c *gin.Context) {
	db := infra.ConnectDB()

	person, err := infra.GetPersonByID(db, c.Param("id"))

	if err != nil {
		c.Status(404)
		return
	}

	c.IndentedJSON(200, person)

}

package service

import (
	"rinha-backend/src/infra"
	"rinha-backend/src/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetByTerm(c *gin.Context) {
	var term string

	queryPart := strings.Split(c.Request.URL.RawQuery, "t=")

	if len(queryPart) > 1 {
		term = queryPart[1]
	} else {
		c.JSON(400, gin.H{"error": "termo não informado"})
		return
	}

	db := infra.ConnectDB()
	person, err := infra.GetByTerm(db, term)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(person) == 0 || person == nil {
		c.IndentedJSON(200, []types.Person{})
		return
	}

	c.IndentedJSON(200, person)
}

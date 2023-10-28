package service

import (
	"rinha-backend/src/infra"
	"rinha-backend/src/types"

	"github.com/gin-gonic/gin"
)

func GetByTerm(c *gin.Context, db types.IPgx) {
	term := c.Query("t")
	if term == "" {
		c.Status(400)
		return
	}

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

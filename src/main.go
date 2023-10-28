package main

import (
	"os"
	"rinha-backend/src/infra"
	"rinha-backend/src/service"
	"rinha-backend/src/types"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := os.Getenv("HTTP_PORT")

	if port == "" {
		port = "666"
	}
	db := infra.ConnectDB()
	batchChannelCreatePerson := make(chan types.Person, 5000)
	resultChannel := make(chan infra.CreateResult)

	go processBatch(batchChannelCreatePerson, resultChannel, db)

	router.POST("/pessoas", func(c *gin.Context) {
		service.SavePerson(c, batchChannelCreatePerson, resultChannel)
	})
	router.GET("/pessoas/:id",
		func(c *gin.Context) {
			service.GetByID(c, db)
		})
	router.GET("/pessoas", func(c *gin.Context) {
		service.GetByTerm(c, db)
	})

	router.GET("/contagem-pessoas", func(c *gin.Context) {
		service.CountPersons(c, db)
	})
	router.Run(":" + port)
}

func processBatch(batchChannel chan types.Person, resultChannel chan infra.CreateResult, db types.IPgx) {
	for person := range batchChannel {

		createdPerson, err, code := infra.InsertPerson(db, person)

		if err != nil {
			resultChannel <- infra.CreateResult{Person: person, Error: err, StatusCode: code}
		} else {
			resultChannel <- infra.CreateResult{Person: createdPerson, Error: nil, StatusCode: code}
		}
	}
}

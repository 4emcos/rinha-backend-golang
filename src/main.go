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

	batchChannelCreatePerson := make(chan types.Person, 5000)
	resultChannel := make(chan infra.CreateResult)

	go processBatch(batchChannelCreatePerson, resultChannel)

	router.POST("/pessoas", func(c *gin.Context) {
		service.SavePerson(c, batchChannelCreatePerson, resultChannel)
	})
	router.GET("/pessoas/:id", service.GetByID)
	router.GET("/pessoas", service.GetByTerm)
	router.GET("/contagem-pessoas", service.GetCountPersons)
	router.Run(":" + port)
}

func processBatch(batchChannel chan types.Person, resultChannel chan infra.CreateResult) {
	for person := range batchChannel {
		db := infra.ConnectDB()
		createdPerson, err, code := infra.InsertPerson(db, person)

		if err != nil {
			resultChannel <- infra.CreateResult{Person: person, Error: err, StatusCode: code}
		} else {
			resultChannel <- infra.CreateResult{Person: createdPerson, Error: nil, StatusCode: code}
		}
	}
}

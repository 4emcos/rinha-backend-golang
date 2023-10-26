package service

import (
	"bytes"
	"encoding/json"
	"reflect"
	"rinha-backend/src/infra"
	"rinha-backend/src/types"

	"github.com/gin-gonic/gin"
)

func SavePerson(c *gin.Context, batchChannel chan types.Person, resultChannel chan infra.CreateResult) {
	var person types.Person

	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var buf bytes.Buffer
	if err := json.Compact(&buf, raw); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	data := buf.Bytes()

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	stack, ok := result["stack"].([]interface{})
	if !ok {
		c.JSON(422, gin.H{"error": "Erro ao obter a lista 'stack'"})
		return
	}

	for _, item := range stack {
		if _, isString := item.(string); !isString {
			c.JSON(400, gin.H{"error": "array não contém apenas strings"})
			return
		}
	}

	if err := json.Unmarshal(data, &person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if person.Nome == "" || person.Apelido == "" || reflect.TypeOf(person.Nome).Kind() != reflect.String {
		c.JSON(422, gin.H{"error": "tipo inválido ou nulo"})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	batchChannel <- types.Person{
		Nome:       person.Nome,
		Apelido:    person.Apelido,
		Nascimento: person.Nascimento,
		Stack:      person.Stack,
	}

	res := <-resultChannel

	if res.Error != nil {
		c.JSON(res.StatusCode, gin.H{"error": res.Error.Error()})
		return
	}

	c.Header("Location", "/pessoas/"+*res.Person.ID)
	c.JSON(res.StatusCode, res.Person)

}

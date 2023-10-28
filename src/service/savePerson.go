package service

import (
	"errors"
	"rinha-backend/src/infra"
	"rinha-backend/src/types"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	InvalidDtoErr = errors.New("invalid request")
)

func Validate(c *types.Person) error {
	if len(c.Apelido) > 32 {
		return InvalidDtoErr
	}

	if len(c.Nome) > 100 {
		return InvalidDtoErr
	}

	dateLayout := "2006-01-02"
	if _, err := time.Parse(dateLayout, c.Nascimento); err != nil {
		return InvalidDtoErr
	}

	for _, tech := range c.Stack {
		if len(tech) > 32 {
			return InvalidDtoErr
		}
	}

	return nil
}

func SavePerson(c *gin.Context, batchChannel chan types.Person, resultChannel chan infra.CreateResult) {
	input := &types.Person{}

	err := c.ShouldBindJSON(input)

	if err != nil {
		c.Status(422)
		return
	}

	if err := Validate(input); err != nil {
		c.Status(422)
		return
	}

	// raw, err := c.GetRawData()
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// var buf bytes.Buffer
	// if err := json.Compact(&buf, raw); err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// data := buf.Bytes()

	// var result map[string]interface{}
	// if err := json.Unmarshal(data, &result); err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	// stack, ok := result["stack"].([]interface{})
	// if !ok {
	// 	c.JSON(422, gin.H{"error": "Erro ao obter a lista 'stack'"})
	// 	return
	// }

	// for _, item := range stack {
	// 	if _, isString := item.(string); !isString {
	// 		c.JSON(400, gin.H{"error": "array não contém apenas strings"})
	// 		return
	// 	}
	// }

	// if err := json.Unmarshal(data, &person); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// if person.Nome == "" || person.Apelido == "" || reflect.TypeOf(person.Nome).Kind() != reflect.String {
	// 	c.JSON(422, gin.H{"error": "tipo inválido ou nulo"})
	// 	return
	// }

	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	batchChannel <- types.Person{
		Nome:       input.Nome,
		Apelido:    input.Apelido,
		Nascimento: input.Nascimento,
		Stack:      input.Stack,
	}

	res := <-resultChannel

	if res.Error != nil {
		c.JSON(res.StatusCode, gin.H{"error": res.Error.Error()})
		return
	}

	c.Header("Location", "/pessoas/"+*res.Person.ID)
	c.JSON(res.StatusCode, res.Person)

}

package infra

import (
	"context"
	"database/sql"
	"fmt"
	"rinha-backend/src/types"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func InsertPerson(db types.IPgx, person types.Person) (types.Person, error, int) {
	var id = uuid.New().String()

	create, err := db.Exec(context.Background(), `
        INSERT INTO person (id, nome, apelido, nascimento, stack)
        SELECT $1, $2, $3, $4, $5
        WHERE NOT EXISTS (
            SELECT 1 FROM person WHERE apelido = $3
        )
    `, &id, person.Nome, person.Apelido, person.Nascimento, pq.Array(person.Stack))

	if err != nil {
		return types.Person{}, err, 500
	}

	if create.RowsAffected() == 0 {
		return types.Person{}, fmt.Errorf("apelido já cadastrado: %s", person.Apelido), 422
	}

	person.ID = &id

	fmt.Printf("Inserted person with ID %s: %+v\n", id, person)
	fmt.Println("Data inserted successfully.")

	return person, nil, 201
}

func GetByTerm(db types.IPgx, term string) ([]types.Person, error) {

	fmt.Println("Searching for: " + term)

	rows, err := db.Query(context.Background(), "SELECT id, nome, apelido, nascimento, stack FROM person WHERE stack @> $1 OR nome ILIKE $2 OR apelido ILIKE $2", pq.Array([]string{term}), "%"+term+"%")

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	var people []types.Person

	for rows.Next() {
		var person types.Person
		err := rows.Scan(&person.ID, &person.Nome, &person.Apelido, &person.Nascimento, &person.Stack)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return people, nil
}

func GetPersonByID(db types.IPgx, id string) (types.Person, error) {
	var person types.Person
	err := db.QueryRow(context.Background(), "SELECT id,nome, apelido, nascimento, stack FROM person WHERE id = $1", id).
		Scan(&person.ID, &person.Nome, &person.Apelido, &person.Nascimento, &person.Stack)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Person{}, fmt.Errorf("ID não encontrado: %s", id)
		}
		return types.Person{}, err
	}

	return person, nil
}

func CountPresons(db types.IPgx) (int, error) {
	var count int
	err := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM person").Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

type CreateResult struct {
	Person     types.Person
	Error      error
	StatusCode int
}

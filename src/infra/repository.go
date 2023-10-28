package infra

import (
	"context"
	"database/sql"
	"fmt"
	"rinha-backend/src/types"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lib/pq"
)

func InsertPerson(db types.IPgx, person types.Person) (types.Person, error, int) {
	var id = uuid.New().String()

	fmt.Println("Inserting person: " + person.Apelido)
	fmt.Println("ID: " + id)
	_, err := db.Exec(context.Background(), `
		INSERT INTO pessoas(id, apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4, $5)`, &id, person.Apelido, person.Nome, person.Nascimento, pq.Array(person.Stack))

	if err != nil {
		fmt.Println(err)
		pqErr := err.(*pgconn.PgError)
		if pqErr.Code == "23505" {
			return types.Person{}, fmt.Errorf("apelido já cadastrado: %s", person.Apelido), 422
		}
		return types.Person{}, err, 500
	}

	// if create.RowsAffected() == 0 {
	// 	return types.Person{}, fmt.Errorf("apelido já cadastrado: %s", person.Apelido), 422
	// }

	person.ID = &id

	fmt.Printf("Inserted person with ID %s: %+v\n", id, person)
	fmt.Println("Data inserted successfully.")

	return person, nil, 201
}

func GetByTerm(db types.IPgx, term string) ([]types.Person, error) {

	fmt.Println("Searching for: " + term)

	rows, err := db.Query(context.Background(), `SELECT id, apelido, nome, nascimento, stack FROM pessoas p WHERE p.BUSCA_TRGM ILIKE '%' || $1 || '%' LIMIT 50;`, term)

	if err != nil {

		return nil, err
	}

	fmt.Println("Data retrieved successfully.")
	var people []types.Person

	for rows.Next() {
		person := types.Person{}
		err := rows.Scan(&person.ID, &person.Nome, &person.Apelido, &person.Nascimento, pq.Array(&person.Stack))
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}

	return people, nil
}

func GetPersonByID(db types.IPgx, id string) (types.Person, error) {
	var person types.Person
	fmt.Println("Searching for: " + id)
	err := db.QueryRow(context.Background(), "SELECT id,nome, apelido, nascimento, stack FROM pessoas WHERE id = $1", id).
		Scan(&person.ID, &person.Nome, &person.Apelido, &person.Nascimento, pq.Array(&person.Stack))

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
	err := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM pessoas").Scan(&count)

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

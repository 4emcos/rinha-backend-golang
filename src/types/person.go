package types

type Person struct {
	ID         *string  `json:"id"`
	Apelido    string   `json:"apelido" validate:"required,max=32"`
	Nome       string   `json:"nome" validate:"required,max=100"`
	Nascimento string   `json:"nascimento" validate:"required,datetime=2006-01-02"`
	Stack      []string `json:"stack" validate:"dive,max=32"`
}

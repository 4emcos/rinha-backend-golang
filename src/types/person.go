package types

type Person struct {
	ID         *string  `json:"id"`
	Nome       string   `json:"nome"`
	Apelido    string   `json:"apelido"`
	Nascimento string   `json:"nascimento"`
	Stack      []string `json:"stack"`
}

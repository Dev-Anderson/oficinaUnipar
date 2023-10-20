package models

type User struct {
	Email          string `json:"email"`
	Senha          string `json:"senha"`
	Nome           string `json:"nome"`
	DataNascimento string `json:"datanascimento"`
	Documento      string `json:"documento"`
	DataCadastro   string `json:"datacadastro"`
	DataAlteracao  string `json:"dataalteracao"`
}

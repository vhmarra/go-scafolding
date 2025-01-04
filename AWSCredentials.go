package main

type AWSCredentials struct {
	Id     string
	Secret string
	Token  string //TODO acho que nao precisa disso
	Region string
	Url    string
	Json   string
}

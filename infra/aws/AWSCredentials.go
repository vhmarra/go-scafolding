package aws

type Credentials struct {
	Id     string
	Secret string
	Token  string //TODO acho que nao precisa disso
	Region string
	Url    string
	Json   string
}

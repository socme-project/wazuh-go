package wazuhapi

type WazuhAPI struct {
	Host     string
	Port     string
	Username string
	Password string
	Token    string
	Indexer  Indexer
	Insecure bool
}

type Indexer struct {
	Username string
	Password string
	Host     string
	Port     string
}

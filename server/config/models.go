package config

type Configuration struct {
	Environment Environment
}

type Environment struct {
	Port    string
	Address string
}

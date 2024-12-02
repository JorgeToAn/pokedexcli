package api

type Config struct {
	Next     *string
	Previous *string
}

var DefaultConfig = Config{
	Next:     nil,
	Previous: nil,
}

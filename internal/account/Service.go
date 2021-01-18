package account

import "fmt"

// Service struct to hold repository
type Service struct {
	config Config
}

// NewService create service struct
func NewService(config Config) *Service {
	return &Service{
		config: config,
	}
}

func (s Service) insertTransaction() {
	fmt.Println("cheguei aqui")
}

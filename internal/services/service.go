package services

import (
	"fmt"
)

type Service struct {
}

func (s Service) Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

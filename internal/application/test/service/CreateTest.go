package service

import "fmt"

type CreateTestService struct {
}

func (c *CreateTestService) CreateTest() {
	fmt.Println("Creating tests")
}

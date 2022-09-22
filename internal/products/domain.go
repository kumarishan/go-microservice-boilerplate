package products

import "github.com/kumarishan/go-microservice-boilerplate/pkg/repo"

type ProductId string

func (p ProductId) String() string {
	return string(p)
}

type Product struct {
	repo.Model[ProductId]
	Name   string `gorm:"index"`
	Status Status `gorm:"index"`
}

type Status uint

const (
	StatusActive Status = iota
	StatusInActive
)

func (s Status) String() (ret string) {
	switch s {
	case StatusActive:
		ret = "Active"
	case StatusInActive:
		ret = "InActive"
	}
	return
}

func StatusFor(status string) (s Status) {
	switch status {
	case "Active":
		s = StatusActive
	case "InActive":
		s = StatusInActive
	}
	return
}

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
	StatusUnknown Status = iota
	StatusActive
	StatusInActive
)

func (p *Product) IsActive() bool {
	return p.Status == StatusActive
}

func (p *Product) SetStatus(status Status) bool {
	// sample logic to showacase pointer methods, will be building a complete state machine
	if status == StatusActive && p.Status == StatusInActive {
		p.Status = StatusActive
		return true
	}
	return false
}

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "Active"
	case StatusInActive:
		return "InActive"
	}
	return "Unknown"
}

func StatusFor(status string) Status {
	switch status {
	case "Active":
		return StatusActive
	case "InActive":
		return StatusInActive
	}
	return StatusUnknown
}

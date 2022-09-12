package products

type ProductId string

func (p ProductId) String() string {
	return string(p)
}

type Product struct {
	ID     ProductId
	Name   string
	Status string
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

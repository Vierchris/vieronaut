package vieronaut

import "fmt"

// go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
type Vieronaut struct {
	Name   string
	Health int // 0 dead, 100 full health
}

func Quest(v Vieronaut) (success bool, err error) {
	// do something
	// do something
	// do something

	v, err = fight(v)
	if err != nil {
		return false, err
	}

	switch v.Health {
	case 0:
		return false, nil
	case 50:
		// sth magically happens when his health is at 50 here
		fmt.Println("Vieronaut is at 50 health")
	case 1:
		// sth magically happens when his health is at 1 here
		fmt.Println("Vieronaut is at 50 health")
	default:
		// do sth else
	}
	return true, nil
}

func fight(Vieronaut) (Vieronaut, error) {
	// do something
	// do something with vieronaut
	// whatever happens to him, we might return error or /and his health is 0
	return Vieronaut{}, nil
}

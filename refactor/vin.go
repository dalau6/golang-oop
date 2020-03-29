package vin

import "fmt"

type VIN interface {
	Manufacturer() string
}

type vin string

func NewVIN(code string) (vin, error) {

	if len(code) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
	}

	// ... check for disallowed characters ...

	return vin(code), nil
}

func (v vin) Manufacturer() string {

	return string(v[:3])
}

type vinEU vin

func NewEUVIN(code string) (vinEU, error) {

	// call super constructor
	v, err := NewVIN(code)

	// and cast to own type
	return vinEU(v), err
}

func (v vinEU) Manufacturer() string {

	// call manufacturer on supertype
	manufacturer := vin(v).Manufacturer()

	// add EU specific postfix if appropriate
	if manufacturer[2] == '9' {
		manufacturer += string(v[11:14])
	}

	return manufacturer
}

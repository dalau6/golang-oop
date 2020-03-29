package vin

type VIN string

func (v VIN) Manufacturer() string {

	manufacturer := v[:3]
	if manufacturer[2] == '9' {
		manufacturer += v[11:14]
	}

	return string(manufacturer)
}

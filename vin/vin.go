package vin

func Manufacturer(vin string) string {

  manufacturer := vin[: 3]
  // if the last digit of the manufacturer ID is a 9
  // the digits 12 to 14 are the second part of the ID
  if manufacturer[2] == '9' {
    manufacturer += vin[11: 14]
  }

  return manufacturer
}
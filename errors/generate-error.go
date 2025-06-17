package main

import (
	"strconv"
  "errors"
)

func ParseLatitude(str string) (float64, error) {
  latitude, err := strconv.ParseFloat(str, 64)
  if err != nil {
    return 0, err
  } else if latitude < -90 || latitude > 90 {
    return 0, errors.New("the value is out of range")
  }
  return latitude, nil
}

func main() {
  ParseLatitude("50")
  ParseLatitude("100")
}


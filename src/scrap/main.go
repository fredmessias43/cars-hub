package main

import (
	"fmt"
	"reflect"

	"github.com/fredmessias43/car-hub/src/contracts"
	"github.com/fredmessias43/car-hub/src/models"
)

func getDynamicValue(model contracts.Model, key string) string {
	r := reflect.ValueOf(&model).FieldByName(key)
	return r.FieldByName(key).String()
}

func main() {
	manufacturer := models.Manufacturer{ID: 123, Name: "manufacturer 1"}

	// value := getDynamicValue(&manufacturer, "ID")

	fmt.Println(manufacturer.ToMap())

	// a := models.List[models.CarModelVersion]{}
	// fmt.Println(a.Value[0].Name)
}

/* manufacturer
brand
carModel
carModelVersion
car */

package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/contracts"
	"github.com/fredmessias43/car-hub/src/database"
	"github.com/fredmessias43/car-hub/src/models"
)

func getDynamicValue(model contracts.Model, key string) string {
	r := reflect.ValueOf(&model).FieldByName(key)
	return r.FieldByName(key).String()
}

func main() {
	var err error
	config.DB, err = database.NewDatabaseClient()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	/*
		manufacturer_ID, _ := strconv.Atoi("")
		manufacturer := models.Manufacturer{}
		if manufacturer_ID == 0 {
			manufacturer.Name = "Teste"
			manufacturer.CountryOrigin = "Pais Teste"
			config.DB.Save(&manufacturer)
			}
			if manufacturer_ID != 0 {
				config.DB.Find(&manufacturer, manufacturer_ID)
				} */

	manufacturer_ID, _ := strconv.Atoi("")
	manufacturer := models.Manufacturer{
		ID:            manufacturer_ID,
		Name:          "Test asdasde",
		CountryOrigin: "Pais Teste",
	}
	config.DB.Save(&manufacturer)

	fmt.Print(manufacturer.Name)
}

/* manufacturer
brand
carModel
carModelVersion
car */

package handlers

import (
	"net/http"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IntegratedHandler struct {
	DB *gorm.DB
}

func (h IntegratedHandler) Index(c *gin.Context) {
	manufacturerList := models.List[models.Manufacturer]{}
	config.DB.Find(&manufacturerList.Value)

	brandList := models.List[models.Brand]{}
	config.DB.Find(&brandList.Value)

	carModelList := models.List[models.CarModel]{}
	config.DB.Find(&carModelList.Value)

	carModelVersionList := models.List[models.CarModelVersion]{}
	config.DB.Find(&carModelVersionList.Value)

	carList := models.List[models.Car]{}
	config.DB.Find(&carList.Value)

	c.HTML(http.StatusOK, "", templates.IntegratedPage(manufacturerList, brandList, carModelList, carModelVersionList, carList))
}

func (h IntegratedHandler) Store(c *gin.Context) {

	/* Manufacturer.ID=
	Manufacturer.Name=
	Manufacturer.CountryOrigin=

	Brand.ID=
	Brand.Name=

	CarModel.ID=
	CarModel.Name=
	CarModel.ManufactureYear=

	CarModelVersion.ID=
	CarModelVersion.Name=
	CarModelVersion.Specifications=
	CarModelVersion.Features=

	Car.Color=
	Car.Mileage=
	Car.Price=
	Car.FuelType=
	Car.Transmission= */

	print(c.Request.Body)

	//

	manufacturerList := models.List[models.Manufacturer]{}
	config.DB.Find(&manufacturerList.Value)

	brandList := models.List[models.Brand]{}
	config.DB.Find(&brandList.Value)

	carModelList := models.List[models.CarModel]{}
	config.DB.Find(&carModelList.Value)

	carModelVersionList := models.List[models.CarModelVersion]{}
	config.DB.Find(&carModelVersionList.Value)

	carList := models.List[models.Car]{}
	config.DB.Find(&carList.Value)

	c.HTML(http.StatusOK, "", templates.IntegratedPage(manufacturerList, brandList, carModelList, carModelVersionList, carList))
}

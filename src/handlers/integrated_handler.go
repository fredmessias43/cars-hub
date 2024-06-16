package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/models"
	"github.com/fredmessias43/car-hub/src/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IntegratedHandler struct {
	DB *gorm.DB
}

func (h IntegratedHandler) Create(c *gin.Context) {
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

	c.HTML(http.StatusOK, "", templates.IntegratedPage(manufacturerList, brandList, carModelList, carModelVersionList, carList, models.Manufacturer{}, models.Brand{}, models.CarModel{}, models.CarModelVersion{}, models.Car{}))
}

func (h IntegratedHandler) Store(c *gin.Context) {
	var body IntegratedBody
	if err := c.BindJSON(&body); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "", err)
		return
	}

	// Manufacturer
	manufacturer_ID, _ := strconv.Atoi(body.Manufacturer_ID)
	manufacturer := models.Manufacturer{
		ID:            manufacturer_ID,
		Name:          body.Manufacturer_Name,
		CountryOrigin: body.Manufacturer_CountryOrigin,
	}
	config.DB.Save(&manufacturer)

	// Brand
	brand_ID, _ := strconv.Atoi(body.Brand_ID)
	brand := models.Brand{
		ID:             brand_ID,
		Name:           body.Brand_Name,
		ManufacturerID: manufacturer.ID,
	}
	config.DB.Save(&brand)

	// CarModel
	carModel_ID, _ := strconv.Atoi(body.CarModel_ID)
	carModel_ManufactureYear, _ := strconv.Atoi(body.CarModel_ManufactureYear)
	carModel := models.CarModel{
		ID:              carModel_ID,
		Name:            body.CarModel_Name,
		ManufactureYear: carModel_ManufactureYear,
		BrandID:         brand.ID,
	}
	config.DB.Save(&carModel)

	// CarModelVersion
	carModelVersion_ID, _ := strconv.Atoi(body.CarModelVersion_ID)
	carModelVersion := models.CarModelVersion{
		ID:             carModelVersion_ID,
		Name:           body.CarModelVersion_Name,
		Specifications: body.CarModelVersion_Specifications,
		Features:       body.CarModelVersion_Features,
		ModelID:        carModel.ID,
	}
	config.DB.Save(&carModelVersion)

	// Car
	car_Mileage, _ := strconv.Atoi(body.Car_Mileage)
	car_Price, _ := strconv.ParseFloat(body.Car_Price, 64)
	car := models.Car{
		Color:          body.Car_Color,
		Mileage:        car_Mileage,
		Price:          car_Price,
		FuelType:       body.Car_FuelType,
		Transmission:   body.Car_Transmission,
		ModelVersionID: carModel.ID,
	}
	config.DB.Save(&car)

	c.HTML(http.StatusOK, "", templates.IntegratedSuccessSaveMessage())
}

func (h IntegratedHandler) Edit(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param(("car")))

	fmt.Println(ID)

	car := models.Car{ID: ID}
	_ = config.DB.Find(&car)

	carModelVersion := models.CarModelVersion{ID: car.ModelVersionID}
	config.DB.Find(&carModelVersion)

	carModel := models.CarModel{ID: carModelVersion.ModelID}
	config.DB.Find(&carModel)

	brand := models.Brand{ID: carModel.BrandID}
	config.DB.Find(&brand)

	manufacturer := models.Manufacturer{ID: brand.ManufacturerID}
	config.DB.Find(&manufacturer)

	/////////////////////

	manufacturerList := models.List[models.Manufacturer]{}
	config.DB.Find(&manufacturerList.Value)

	brandList := models.List[models.Brand]{}
	config.DB.Find(&brandList.Value)

	carModelList := models.List[models.CarModel]{}
	config.DB.Find(&carModelList.Value)

	carModelVersionList := models.List[models.CarModelVersion]{}
	config.DB.Find(&carModelVersionList.Value)

	carList := models.List[models.Car]{}
	// config.DB.Find(&carList.Value)

	c.HTML(http.StatusOK, "", templates.IntegratedPage(manufacturerList, brandList, carModelList, carModelVersionList, carList, manufacturer, brand, carModel, carModelVersion, car))
}

func (h IntegratedHandler) Update(c *gin.Context) {
	var body IntegratedBody
	if err := c.Bind(&body); err != nil {
		fmt.Println("=======================> " + "aaaaaaaaaa" + " <======================")
		c.HTML(http.StatusUnprocessableEntity, "", templates.UnprocessableEntity(err))
		return
	}

	// Manufacturer
	manufacturer_ID, _ := strconv.Atoi(body.Manufacturer_ID)
	manufacturer := models.Manufacturer{
		ID:            manufacturer_ID,
		Name:          body.Manufacturer_Name,
		CountryOrigin: body.Manufacturer_CountryOrigin,
	}
	config.DB.Save(&manufacturer)

	// Brand
	brand_ID, _ := strconv.Atoi(body.Brand_ID)
	brand := models.Brand{
		ID:             brand_ID,
		Name:           body.Brand_Name,
		ManufacturerID: manufacturer.ID,
	}
	config.DB.Save(&brand)

	// CarModel
	carModel_ID, _ := strconv.Atoi(body.CarModel_ID)
	carModel_ManufactureYear, _ := strconv.Atoi(body.CarModel_ManufactureYear)
	carModel := models.CarModel{
		ID:              carModel_ID,
		Name:            body.CarModel_Name,
		ManufactureYear: carModel_ManufactureYear,
		BrandID:         brand.ID,
	}
	config.DB.Save(&carModel)

	// CarModelVersion
	carModelVersion_ID, _ := strconv.Atoi(body.CarModelVersion_ID)
	carModelVersion := models.CarModelVersion{
		ID:             carModelVersion_ID,
		Name:           body.CarModelVersion_Name,
		Specifications: body.CarModelVersion_Specifications,
		Features:       body.CarModelVersion_Features,
		ModelID:        carModel.ID,
	}
	config.DB.Save(&carModelVersion)

	// Car
	car_Mileage, _ := strconv.Atoi(body.Car_Mileage)
	car_Price, _ := strconv.ParseFloat(body.Car_Price, 64)
	car := models.Car{
		Color:          body.Car_Color,
		Mileage:        car_Mileage,
		Price:          car_Price,
		FuelType:       body.Car_FuelType,
		Transmission:   body.Car_Transmission,
		ModelVersionID: carModel.ID,
	}
	config.DB.Save(&car)

	c.HTML(http.StatusOK, "", templates.IntegratedSuccessSaveMessage())
}

type IntegratedBody struct {
	Manufacturer_ID            string `json:"Manufacturer_ID" binding:"required"`
	Manufacturer_Name          string `json:"Manufacturer_Name" binding:"required"`
	Manufacturer_CountryOrigin string `json:"Manufacturer_CountryOrigin" binding:"required"`
	//
	Brand_ID   string `json:"Brand_ID" binding:"required"`
	Brand_Name string `json:"Brand_Name" binding:"required"`
	//
	CarModel_ID              string `json:"CarModel_ID" binding:"required"`
	CarModel_Name            string `json:"CarModel_Name" binding:"required"`
	CarModel_ManufactureYear string `json:"CarModel_ManufactureYear" binding:"required"`
	//
	CarModelVersion_ID             string `json:"CarModelVersion_ID" binding:"required"`
	CarModelVersion_Name           string `json:"CarModelVersion_Name" binding:"required"`
	CarModelVersion_Specifications string `json:"CarModelVersion_Specifications" binding:"required"`
	CarModelVersion_Features       string `json:"CarModelVersion_Features" binding:"required"`
	//
	Car_Color        string `json:"Car_Color" binding:"required"`
	Car_Mileage      string `json:"Car_Mileage" binding:"required"`
	Car_Price        string `json:"Car_Price" binding:"required"`
	Car_FuelType     string `json:"Car_FuelType" binding:"required"`
	Car_Transmission string `json:"Car_Transmission" binding:"required"`
}

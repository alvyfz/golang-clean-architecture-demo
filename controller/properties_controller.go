package controller

import (
	"myproperty-clean-architecture/app/model"
	service "myproperty-clean-architecture/service/properties"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PropertiesController struct {
	PropertiesService service.PropertiesService
}

func NewPropertiesController(propertiesService *service.PropertiesService) PropertiesController {
	return PropertiesController{PropertiesService: *propertiesService}
}

func (controller *PropertiesController) Route(e *echo.Echo){


	e.POST("/property",controller.Create )
	e.GET("/property",controller.List )
}

func (controller *PropertiesController) Create(c echo.Context) error {

var request model.CreatePropertiesRequest
if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreatePropertyController",
			"error":   err.Error(),
		})
	}
	response := controller.PropertiesService.Create(request)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Create property success",
		"data":    response,
	})
}

func (controller *PropertiesController) List(c echo.Context) error {
	// properties := service.GetAllProperties()
	// return c.JSON(http.StatusOK, properties)
	responses := controller.PropertiesService.List()
	return c.JSON(http.StatusOK, responses)
}
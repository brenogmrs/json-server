package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"json-server/pkg/service"
	"json-server/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.IService
}

func New(service service.IService) controller {
	return controller{
		service: service,
	}
}

func (this *controller) Create(c *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	utils.ErrorChecker(err)

	jsonBody := string(bodyAsByteArray)

	var valuesMap map[string]interface{}
	json.Unmarshal([]byte(jsonBody), &valuesMap)

	createdBody := this.service.Create(valuesMap)

	c.JSON(http.StatusCreated, createdBody)

}

func (this *controller) GetAll(c *gin.Context) {

	requestedResource := this.service.GetAll()

	jsonStr, err := json.Marshal(requestedResource)

	data := string(jsonStr)

	if err != nil {
		fmt.Println("error: ", err)
	}

	c.Data(http.StatusOK, "application/json", []byte(data))
}

func (this *controller) GetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		customError := utils.NewBadRequestError("unprocessable id")

		c.JSON(customError.Code, customError)
		return
	}

	data, err := this.service.GetById(id)

	if err != nil {
		customError := utils.NewNotFoundError("Resource not found")

		c.JSON(customError.Code, customError)
		return
	}

	c.JSON(http.StatusOK, data)

}

func (this *controller) Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		customError := utils.NewBadRequestError("unprocessable id")

		c.JSON(customError.Code, customError)
		return
	}

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	utils.ErrorChecker(err)

	jsonBody := string(bodyAsByteArray)

	var valuesMap map[string]interface{}
	json.Unmarshal([]byte(jsonBody), &valuesMap)

	if _, ok := valuesMap["id"]; ok {
		delete(valuesMap, "id")
	}

	updatedBody := this.service.Update(id, valuesMap)

	c.JSON(http.StatusOK, updatedBody)
}

func (this *controller) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		customError := utils.NewBadRequestError("Unprocessable id")

		c.JSON(customError.Code, customError)
		return
	}

	if err := this.service.Delete(id); err != nil {
		customError := utils.NewNotFoundError("Resource not found")

		c.JSON(customError.Code, customError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

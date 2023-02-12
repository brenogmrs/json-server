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

func (c *controller) Create(ctx *gin.Context) {

	bodyAsByteArray, err := io.ReadAll(ctx.Request.Body)
	utils.ErrorChecker(err)

	jsonBody := string(bodyAsByteArray)

	var valuesMap map[string]interface{}
	json.Unmarshal([]byte(jsonBody), &valuesMap)

	createdBody := c.service.Create(valuesMap)

	ctx.JSON(http.StatusCreated, createdBody)

}

func (c *controller) GetAll(ctx *gin.Context) {

	requestedResource := c.service.GetAll()

	jsonStr, err := json.Marshal(requestedResource)

	data := string(jsonStr)

	if err != nil {
		fmt.Println("error: ", err)
	}

	ctx.Data(http.StatusOK, "application/json", []byte(data))
}

func (c *controller) GetById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		customError := utils.NewBadRequestError("unprocessable id")

		ctx.JSON(customError.Code, customError)
		return
	}

	data, err := c.service.GetById(id)

	if err != nil {
		customError := utils.NewNotFoundError("Resource not found")

		ctx.JSON(customError.Code, customError)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *controller) Update(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		customError := utils.NewBadRequestError("unprocessable id")

		ctx.JSON(customError.Code, customError)
		return
	}

	bodyAsByteArray, err := io.ReadAll(ctx.Request.Body)
	utils.ErrorChecker(err)

	jsonBody := string(bodyAsByteArray)

	var valuesMap map[string]interface{}
	json.Unmarshal([]byte(jsonBody), &valuesMap)

	if valuesMap["id"] != nil {
		delete(valuesMap, "id")
	}

	updatedBody := c.service.Update(id, valuesMap)

	ctx.JSON(http.StatusOK, updatedBody)
}

func (c *controller) Delete(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		customError := utils.NewBadRequestError("Unprocessable id")

		ctx.JSON(customError.Code, customError)
		return
	}

	if err := c.service.Delete(id); err != nil {
		customError := utils.NewNotFoundError("Resource not found")

		ctx.JSON(customError.Code, customError)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

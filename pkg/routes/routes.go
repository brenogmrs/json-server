package routes

import (
	"fmt"
	"json-server/pkg/controller"
	"json-server/pkg/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, resources []string) {

	for _, str := range resources {

		route := fmt.Sprintf("/%s/", str)
		idParam := ":id"

		serviceHandler := service.New(str)
		controller := controller.New(serviceHandler)

		r.GET(route, controller.GetAll)
		r.GET(route+idParam, controller.GetById)
		r.POST(route, controller.Create)
		r.PUT(route+idParam, controller.Update)
		r.DELETE(route+idParam, controller.Delete)

	}

	// field validation for Create and Update

}

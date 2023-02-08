package main

import (
	"json-server/pkg/middlewares"
	"json-server/pkg/providers"
	"json-server/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	JSONData := providers.GetDBData()

	resourcesNames := providers.GetResources(JSONData)
	
	router := gin.Default()
	router.Use(middlewares.CheckResources(resourcesNames))

	routes.InitRoutes(&router.RouterGroup, resourcesNames)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

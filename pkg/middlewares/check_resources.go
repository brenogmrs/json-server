package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckResources(resources []string) gin.HandlerFunc {
    return func (c *gin.Context) {
        
		requestUrl := strings.ReplaceAll(c.Request.URL.Path, "/", "")

		set := make(map[string]bool)

		for _, v := range resources {
			set[v] = true
		}

		if !set[requestUrl] {
			c.JSON(400, gin.H{"error": "non-existing-resource"})
		}
	
		c.Next()
    }
}


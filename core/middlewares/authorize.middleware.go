package middlewares

import (
	"github.com/bioyeneye/rest-gin-api/core/configs"
	"github.com/gin-gonic/gin"
)

type CurrentUser struct {
	Name string
	Email string
}

func AuthorizeMiddleware() gin.HandlerFunc {
	//errList := make(map[string]string)
	return func(c *gin.Context) {

		currentUser := CurrentUser{
			Name:  "Bolaji",
			Email: "bolaji@gmail.com",
		}

		c.Set(configs.ApplicationUserKey, currentUser)

		/*err := auth.TokenValid(c.Request)
		if err != nil {
			errList["unauthorized"] = "Unauthorized"
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  errList,
			})
			c.Abort()
			return
		}*/
		c.Next()
	}
}

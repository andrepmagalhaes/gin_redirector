package application

import (
	"fmt"

	"github.com/andrepmagalhaes/redirector/application/routes"
	"github.com/gin-gonic/gin"
)

func Init(port int) {
	app := gin.Default()

	app.POST("/login", routes.GetAuth)
	app.POST("/lockers", routes.GetLockers)
	app.POST("/lockers/slots", routes.GetLockersSlots)

	host := fmt.Sprintf("localhost:%d", port)

	app.Run(host)
}

package application

import (
	"github.com/andrepmagalhaes/redirector/application/routes"
	"github.com/gin-gonic/gin"
)

func Init(port string) {
	app := gin.Default()

	app.POST("/login", routes.PostAuth)
	app.POST("/lockers", routes.PostLockers)
	app.POST("/lockers/slots", routes.PostLockersSlots)
	app.POST("/lockers/reservation", routes.PostLockersReservation)

	//host := fmt.Sprintf("localhost:%d", port)

	app.Run(":" + port)
}

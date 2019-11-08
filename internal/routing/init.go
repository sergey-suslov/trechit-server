package routing

import (
	"github.com/labstack/echo"
	v1 "github.com/sergey-suslov/trechit-server/internal/routing/api/v1"
)

// Init routing
func Init() {
	e := echo.New()
	apiv1 := e.Group("/v1")
	v1.InitAPI(apiv1)

	e.Logger.Fatal(e.Start(":1323"))
}

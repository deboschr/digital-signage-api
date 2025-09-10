package routes

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api/v1")
    {
        RegisterDeviceRoutes(api) // load routes device
        // RegisterAirportRoutes(api)
        // RegisterChannelRoutes(api)
    }

    return r
}

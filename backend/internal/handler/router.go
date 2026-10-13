package handler

import "github.com/gin-gonic/gin"

type Dependencias struct {
	Salud map[string]ComprobadorSalud
}

func NuevoRouter(d Dependencias) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/salud", salud(d.Salud))

	return r
}

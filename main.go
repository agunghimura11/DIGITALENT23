package main

import (
	"DIGITALENT23/app/controlller"
	"github.com/gin-gonic/gin"

)

func  main()  {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controlller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controlller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controlller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id:idAntrian/delete", controlller.DeleteAntrianHandler)

	router.GET("/antrian", controlller.PageAntrianHandler)
	router.Run(":8080")
}




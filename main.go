package main

import "net/http"
import "github.com/gin-gonic/gin"

var data []Antrian

type Antrian struct {
	ID string `json:"id"`
	Status bool `json:"status"`
}

func  main()  {
	router := gin.Default()
	router.GET("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status",GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id:idAntrian/delete", DeleteAntrianHandler)
	router.Run(":8000")
}

func AddAntrianHandler(c *gin.Context) {
	return true, data, nil
}

func GetAntrianHandler(c *gin.Context) {

}

func UpdateAntrianHandler(c *gin.Context) {

}

func DeleteAntrianHandler(c *gin.Context) {

}


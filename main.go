package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
	"strings"
)

var client *db.Client
var ctx context.Context

func init(){
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://dts23-aae6f.firebaseio.com/",
	}

	opt := option.WithCredentialsFile("firebase-admin-sdk.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing databse client:", err)
	}
}

var data []Antrian

type Antrian struct {
	Id string `json:"id"`
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

func getAntrian() (bool, error, []Antrian) {
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database: ", err)
		return false, err, nil
	}
	return true, nil, data
}


func deleteAntrian(idAntrian string) (bool, error) {
	//for i := range data {
	//	if data[i].Id == idAntrian {
	//		data = append(data[:i], data[i+1:]...)
	//	}
	//}
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	if err := childRef.Delete(ctx); err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func addAntrian() (bool, error) {
	_, dataAntrian, _ := getAntrian()
	var Id string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")
	if dataAntrian == nil{
		Id = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		Id = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}

	antrian := Antrian {
		Id: Id,
		Status: false,
	}
	if err := antrianRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}

	//data = append(data, Antrian{
	//	Id: Id,
	//	Status: false,
	//})
	return true, nil
}

func AddAntrianHandler(c *gin.Context) {
	flag, err := addAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}

func GetAntrianHandler(c *gin.Context) {
	flag, err, resp := getAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data": resp,
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}
// update antrian
func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := updateAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}

func updateAntrian(idAntrian string) (bool, error)  {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian {
		Id: idAntrian,
		Status: true,
	}
	if err := childRef.Set(ctx, antrian); err != nil {
		log.Fatal(err)
		return false, err
	}
	//for i,_ := range data{
	//	if data[i].Id == idAntrian{
	//		data[i].Status = true
	//		break
	//	}
	//}

	return true, nil
}

// Delete antrian
func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Params("idAntrian")
	flag, err := deleteAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	}else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}

func PageAntrianHandler(c *gin.Context)  {
	flag, result, err := getAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}


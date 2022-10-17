package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/timalbpathirana/golangAPI/pkg/models"
	"github.com/timalbpathirana/golangAPI/pkg/utils"
	"net/http"
	"strconv"
)

/*	{
	"ProductID": 25,
	"ProductName": "Test",
	"ProductSize": "Test size",
	"ProductColour": "Green",
	"ProductPrice": 100,
	"ProductLocation": "Richmond",
	"ProductAvailability": "Y"
}*/

var newProduct models.Product

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	newProduct := models.GetProduct()
	// json marshal (byte to string)
	res, _ := json.Marshal(newProduct)
	setHeader(w)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["id"]
	if ID, err := strconv.ParseInt(productId, 0, 0); err == nil {
		productDetail, _ := models.GetProductById(ID)
		res, _ := json.Marshal(productDetail)
		setHeader(w)
		w.Write(res)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := &models.Product{}
	utils.ParseBody(r, newProduct)
	p := newProduct.CreateProduct()
	res, _ := json.Marshal(p)
	setHeader(w)
	w.Write(res)
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["id"]
	if ID, err := strconv.ParseInt(productId, 0, 0); err == nil {
		productToDelete := models.DeleteProductById(ID)
		res, _ := json.Marshal(productToDelete)
		setHeader(w)
		w.Write(res)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := &models.Product{}
	utils.ParseBody(r, newProduct)
	vars := mux.Vars(r)
	productID := vars["id"]

	ID, err := strconv.ParseInt(productID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	productDetails, db := models.GetProductById(ID)
	productDetails.ProductID = newProduct.ProductID
	productDetails.ProductName = newProduct.ProductName
	productDetails.ProductSize = newProduct.ProductSize
	productDetails.ProductColour = newProduct.ProductColour
	productDetails.ProductPrice = newProduct.ProductPrice
	productDetails.ProductLocation = newProduct.ProductLocation
	productDetails.ProductAvailability = newProduct.ProductAvailability

	//saving the db
	db.Save(&newProduct)

	//sending json response
	res, _ := json.Marshal(newProduct)
	setHeader(w)
	w.Write(res)
}

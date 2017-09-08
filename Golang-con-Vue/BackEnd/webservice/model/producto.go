package model

import (
	cx 	"Projects/Golang-con-Vue/BackEnd/webservice/conexion"
	"github.com/gin-gonic/gin"
	"fmt"
)


//Este es nuestro modelo de producto
type Producto struct {
	ID uint
	Nombre string `json:"nombre"`
	Costo string `json:"costo"`
	//Este dato lo pongo por si alguien sabe como manejar las fechas ojala nos pueda dar un ejemplito
	//Yo lo he manejado como cadena pero bueno funciona
	Caducidad string `json:"caducidad"`
}

func CreateDB() {
	db := cx.Connection()
	db.DropTableIfExists(&Producto{})
	db.CreateTable(&Producto{})
}

func CrearDatos(){

}












/*___________________________

 Funciones de gin y gorm

//___________________________*/


//Funciones para llamar a todos los productos
func ObtenerTodosLosProductos(c *gin.Context) {
	var producto []Producto
	db:=cx.Connection()
	//db.Close()
	if err := db.Find(&producto).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}else {
		c.JSON(200, producto)
	}

}

//Funciones para llamar a producto especifico
func ObtieneUnProducto(c *gin.Context) {

	//Es el id que se envia por la direccion
	id := c.Params.ByName("id")
	var producto Producto
	db:=cx.Connection()

	if err := db.Where("id = ?", id).First(&producto).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}else{
		c.JSON(200, producto)
	}

}

//Funcion para crear empleado
func CreaProducto(c *gin.Context) {
	var producto Producto

	db:=cx.Connection()


	c.BindJSON(&producto)
	fmt.Println(producto)
	err :=db.Create(&producto).Error
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
		return
	}else{

	}

}
//no implemente el update porque esta muy facil ejej
//No la verdad me dio flojerapero les juro que funciona lo pueden checar con el postman
func ActualizaProducto(c *gin.Context) {
	var producto Producto
	id := c.Params.ByName("id")



	db:=cx.Connection()
	//db.Close()
	if err := db.Where("id = ?", id).First(&producto).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}else{
		c.BindJSON(producto)
		db.Save(&producto)
		c.JSON(200, producto)
	}

}

func EliminaProducto(c *gin.Context) {
	id := c.Params.ByName("id")
	var producto Producto
	db := cx.Connection()
	//db.Close()
	d := db.Where("id = ?", id).Delete(&producto)
	err:=d.Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		return
	}else{
		fmt.Println(d)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}



}




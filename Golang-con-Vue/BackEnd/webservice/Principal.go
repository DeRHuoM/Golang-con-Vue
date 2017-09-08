package main

import (
	"github.com/gin-gonic/gin"
	m "Projects/Golang-con-Vue/BackEnd/webservice/model"
)


//Metodo donde creo la base de datos
//Con el DropTa... borro todo si existe
//Con el Create table creo la tabla jaja
//Bueno ahora creare unos datos basura jeje



func main()  {

	//Primero llamo la base de datos
	m.CreateDB()
	m.CrearDatos()

	///
	r := gin.Default()
	r.Use(Cors())

	// Definimos nuestras direcciones para nuestra api
	r.GET("/producto", m.ObtenerTodosLosProductos)
	r.GET("/producto/:id", m.ObtieneUnProducto)
	r.POST("/producto", m.CreaProducto)
	r.PUT("/producto/:id", m.ActualizaProducto)
	r.DELETE("/producto/:id", m.EliminaProducto)
	//Definimos el puerto donde se iniciara el servidor
	r.Run(":9060")



}
// Esto la verdad un no lo entiendo bien
//Si alguien nos ilustra con una explicacion seria de gran adyua
func Cors() gin.HandlerFunc{
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin","*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		context.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type")
		context.Writer.Header().Set("Access-Control-Allow-Methods","POST, OPTIONS, GET, PUT, UPDATE,DELETE")
		if context.Request.Method=="OPTIONS"{
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
}

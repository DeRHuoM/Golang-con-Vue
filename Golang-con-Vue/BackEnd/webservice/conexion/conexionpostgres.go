package conexion

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

)

type db struct {
	user string
	pass string
	port string
	namedb string
	host string
}
//Debe iniciar en mayusculas para que aparesca el metodo donde se importa
//Creanlo o no me tarde un mes en haces esto posible jeje \\\|||////---_____ baah
//Esta conexion de base de datos la hago con postgres, algo muy simple
//Si alguien me ayuda a mejorar este codigo se lo agradeceria ... Mi tesis depende de esto xD
func Connection() ( *gorm.DB) {
	dbCredential := db{
		user:"postgres",
		pass:"1234",
		port:"15432",
		namedb:"tienda",
		host:"localhost",
	}
	db,err := gorm.Open("postgres", "host= "+dbCredential.host+" user= "+dbCredential.user+ " dbname= "+dbCredential.namedb +" sslmode=disable password= "+dbCredential.pass)
	if err!=nil{
		panic(err)
	}

	return  db
}


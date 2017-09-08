# Golang-con-Vue
Implemetacion simple de Front end Vue js con Back End Golang Gin Gorm

El projecto se divide en dos partes Front End y Back End
Se debe tener instalado golang y postgres

# 0.Base de datos

En la base de datos postgres se debe tener creada una base de datos llamada tienda

En la carpera de golang usr/Projects copiar esta carpeta
Direccion relativa
../go/src/Projects/

# 1. Back end

Se debe instalar las librerias gin y gorm
Comandos para instalar paquetes:

go get "github.com/gin-gonic/gin"

go get "github.com/jinzhu/gorm"

Podre una direccion relativa como ejemplo donde se deben situar

../go/src/Projects/Golang-con-Vue/BackEnd/webservice


Compilar el archivo Principal.go

go build Principal.go

//En windows
webservice.exe

//En linux

./webservice

Verificar que el servidor funciona con el navegador
http://localhost:9060/producto




# 2. Front End
Se utiliza Vue js para la instalacion seguir los siguientes pasos

Para ejecutar el projecto se debe intalar node js

https://nodejs.org/es/

Instalar el paquete Vue-Cli

https://github.com/vuejs/vue-cli

Comando para instalacion

npm install -g vue-cli

Podre una direccion relativa como ejemplo donde se deben situar

../go/src/Projects/Golang-con-Vue/FrontEnd/paginaweb

en esta carpeta ejecutar 

npm install

npm run dev

Verifica su funcionamiento

http://localhost:8080


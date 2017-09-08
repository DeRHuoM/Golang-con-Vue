<template>
  <div class="hello">
    <h1>{{ msg }}</h1>


    <form >
        <!-- text -->
        <h3>nombre</h3>
        <input type="text" v-model="dato.nombre">
        <h3>costo</h3>
        <input type="number" v-model="dato.costo">
        <h3>caducidad</h3>
        <input type="date" v-model="dato.caducidad">
        <br>
        <button @click="limpiarDato"> Limpiar</button>
        <button @click="crearProducto"> Guardar</button>
        
    </form>

    <ul >
      <li v-for="item in items">
        {{item}}
         <button @click="eliminarProducto(item)">Eliminar</button>
      </li>
    </ul>

  </div>
</template>

<script>
export default {
  name: 'hello',
  data () {
    return {
      msg: 'Welcome to Your Vue.js with Golang App',
      items:[],
      dato:{
        nombre:'',
        costo:0,
        caducidad:''
      }
    }
  },
  methods :{

    cargarProductos: function () 
    {
      this.axios.get('http://localhost:9060/producto')
      .then((response)=>{
        this.items = response.data;
        console.log(JSON.stringify(this.items));
        //imprime lo que recibe del servidor
        console.log(JSON.stringify(response.data));
      });  
      
    },
    eliminarProducto: function(item)
    {
      console.log(JSON.stringify(item));
      this.axios.delete('http://localhost:9060/producto/'+item.ID)
      .then((response)=>{
        console.log("Poducto eliminado correctamente");
        this.cargarProductos(); 
      });
      //Recarga datos desde el web service yeaah
      
    },
    crearProducto: function()
    {
      //this.dato hace referencia al objeto creado en el formulario
      this.axios.post('http://localhost:9060/producto',JSON.stringify(this.dato))
      .then((response)=>{
        console.log(JSON.stringify(response.data));
         //Carga poductos del web service
        this.cargarProductos();
        //Limpia datos
        this.limpiarDato();
      });
     
    },
    limpiarDato: function()
    {
      this.dato.costo=0;
      this.dato.nombre='';
      this.dato.caducidad='';
    }
  },
  mounted(){
    this.cargarProductos();
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>

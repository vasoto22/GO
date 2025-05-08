// Paquete principal del programa
package main

// ------------------------
// IMPORTACIÓN DE PAQUETES
// ------------------------
// Aquí estamos importando el paquete 'fmt', que nos permite imprimir en la consola y formatear el texto.
import (
	"fmt"
)

// ------------------------
// DEFINICIÓN DE ESTRUCTURAS
// ------------------------
// Aquí definimos la estructura 'Persona', que tiene dos campos: 'Nombre' y 'Edad'.
type Persona struct {
	Nombre string
	Edad   int
}

// ------------------------
// FUNCIÓN PRINCIPAL: main
// ------------------------
// La función 'main' es el punto de entrada al programa. Cuando ejecutamos el programa, Go comienza a ejecutar desde aquí.
func main() {

	// ------------------------
	// DECLARACIÓN DE VARIABLES
	// ------------------------
	// En Go, podemos declarar variables de dos maneras: explícitamente (con el tipo) o mediante inferencia de tipo.
	var nombre string = "Miguel" // Declaración explícita de la variable 'nombre' como tipo string.
	edad := 25                   // Go infiere el tipo de la variable 'edad' como 'int', porque asignamos el valor 25.

	// ------------------------
	// IMPRESIÓN DE MENSAJES
	// ------------------------
	// Usamos fmt.Println para imprimir mensajes en la consola.
	fmt.Println("¡Hola, Mundo!")                                 // Imprime un mensaje estático.
	fmt.Println("Mi nombre es", nombre, "y tengo", edad, "años") // Imprime el nombre y la edad de la persona.

	// ------------------------
	// CREACIÓN Y USO DE ESTRUCTURAS
	// ------------------------
	// Creamos una instancia de la estructura 'Persona' y le asignamos valores a sus campos.
	p1 := Persona{
		Nombre: "Maria", // Asignamos el valor "Maria" al campo 'Nombre'
		Edad:   30,      // Asignamos el valor 30 al campo 'Edad'
	}
	// Imprimimos los detalles de la persona creada.
	fmt.Println("Persona:", p1.Nombre, "con", p1.Edad, "años")

	// ------------------------
	// CICLO 'FOR'
	// ------------------------
	// El ciclo 'for' en Go es una estructura que permite ejecutar un bloque de código repetidamente.
	// En este caso, estamos usando el ciclo para contar de 1 a 5.
	for i := 1; i <= 5; i++ {
		// Cada vez que el ciclo se ejecuta, imprime el número de la iteración.
		fmt.Println("Contando:", i)
	}

	// ------------------------
	// MAPAS (MAPS)
	// ------------------------
	// Los mapas son colecciones de pares clave-valor. En este caso, tenemos un mapa de nombres (claves) y edades (valores).
	mapa_personas := map[string]int{"Juan": 25, "Maria": 30}
	// Usamos 'range' para recorrer el mapa y obtener cada par clave-valor.
	for nombre, edad := range mapa_personas {
		// Imprimimos cada par clave-valor del mapa.
		fmt.Println(nombre, edad)
	}

	// ------------------------
	// CONDICIONAL 'IF-ELSE'
	// ------------------------
	// El condicional 'if-else' se utiliza para ejecutar diferentes bloques de código según si una condición es verdadera o falsa.
	if edad >= 18 {
		// Si la edad es mayor o igual a 18, se ejecuta este bloque.
		fmt.Println("Eres mayor de edad.")
	} else {
		// Si la edad es menor a 18, se ejecuta este bloque.
		fmt.Println("Eres menor de edad.")
	}

	// ------------------------
	// LLAMADA A UNA FUNCIÓN
	// ------------------------
	// Aquí estamos llamando a la función 'saludar' y pasando el argumento "Carlos".
	saludar("Carlos") // La función 'saludar' imprimirá un saludo con el nombre "Carlos".

	// ------------------------
	// MAPAS DE ESTRUCTURAS
	// ------------------------
	// En Go, los mapas pueden almacenar cualquier tipo de datos, incluso estructuras. Aquí, almacenamos instancias de la estructura 'Persona' en un mapa.
	personas := make(map[string]Persona) // Creamos un mapa vacío de tipo 'map[string]Persona'
	// Insertamos personas en el mapa usando claves de tipo string (nombres).
	personas["juan"] = Persona{Nombre: "Juan", Edad: 25}
	personas["maria"] = Persona{Nombre: "Maria", Edad: 30}
	// Imprimimos el mapa completo.
	fmt.Println("Mapa de personas:", personas)
	// Accedemos a una persona específica mediante su clave en el mapa (por ejemplo, "juan").
	fmt.Println("Persona de Juan:", personas["juan"])

	// ------------------------
	// ARREGLOS (ARRAYS)
	// ------------------------
	// Los arreglos en Go tienen un tamaño fijo y están compuestos por elementos del mismo tipo.
	var numeros [5]int // Definimos un arreglo de 5 enteros.
	// Asignamos valores a cada elemento del arreglo.
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50
	// Imprimimos el arreglo completo.
	fmt.Println("Arreglo de números:", numeros)

	// ------------------------
	// SLICES
	// ------------------------
	// Los slices son similares a los arreglos, pero tienen un tamaño dinámico y son más flexibles.
	var frutas = []string{"Manzana", "Banana", "Cereza"} // Creamos un slice de frutas.
	// Imprimimos el slice.
	fmt.Println("Slice de frutas:", frutas)
	// Agregamos un nuevo elemento al slice usando la función 'append'.
	frutas = append(frutas, "Naranja")
	// Imprimimos el slice después de agregar un nuevo elemento.
	fmt.Println("Slice de frutas después de agregar una más:", frutas)

	// ------------------------
	// PUNTEROS
	// ------------------------
	// Los punteros en Go almacenan la dirección de memoria de otra variable en lugar de almacenar el valor directamente.
	var numero int = 58        // Creamos una variable 'numero' con el valor 58.
	var puntero *int = &numero // Creamos un puntero que almacena la dirección de memoria de 'numero'.
	// Imprimimos el valor de la variable y la dirección de memoria donde se guarda.
	fmt.Println("Valor de numero:", numero)
	fmt.Println("Dirección de memoria de numero:", puntero)
	// Desreferenciamos el puntero para obtener el valor de la variable a la que apunta.
	fmt.Println("Valor al que apunta el puntero:", *puntero)
	// Modificamos el valor de la variable a través del puntero.
	*puntero = 100                                // Cambiamos el valor de 'numero' a través del puntero.
	fmt.Println("Nuevo valor de numero:", numero) // Imprimimos el nuevo valor de 'numero' (100).
}

// ------------------------
// FUNCIÓN 'saludar'
// ------------------------
// Aquí definimos una función simple que recibe un argumento (nombre) y lo imprime junto con un saludo.
func saludar(nombre string) {
	// Imprimimos un saludo con el valor recibido como parámetro.
	fmt.Println("Hola,", nombre)
}

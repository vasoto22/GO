package main // Importamos el paquete principal donde se ejecutará el programa

import (
	"fmt"       // Importamos el paquete fmt que nos permite trabajar con entradas/salidas de consola
	"math/rand" // Importamos el paquete rand que nos permite generar números aleatorios
	"time"      // Importamos el paquete time para trabajar con fechas y tiempos
)

func main() {
	// Crear una fuente de números aleatorios basada en el tiempo actual
	// El UnixNano devuelve el tiempo en nanosegundos desde la época UNIX (1 de enero de 1970)
	source := rand.NewSource(time.Now().UnixNano()) // Usamos time.Now().UnixNano() como semilla
	r := rand.New(source)                           // Creamos un nuevo generador de números aleatorios con esta fuente

	begin := time.Now()                          // Guardamos el tiempo actual para calcular la duración de la ejecución
	fmt.Println("==== Construction begins ====") // Imprimimos un mensaje indicando el inicio de la "construcción"

	// Iniciamos tres goroutines que simulan el trabajo concurrente de tres personas
	// 'go' indica que la función se ejecutará en paralelo de manera asíncrona (concurrente)
	go Worker("Bob", "installing a door", r)   // Goroutine para el trabajador Bob
	go Worker("Alice", "mowing the lawn", r)   // Goroutine para la trabajadora Alice
	go Worker("John", "painting the walls", r) // Goroutine para el trabajador John

	// Pausamos el programa durante 100 milisegundos para dar tiempo a que las goroutines terminen
	// En una aplicación real, se utilizaría un mecanismo para sincronizar las goroutines, como un `sync.WaitGroup`.
	time.Sleep(100 * time.Millisecond) // Tiempo de espera para las goroutines

	// Calculamos el tiempo transcurrido desde que se inició la construcción y lo imprimimos
	fmt.Printf("==== Construction ended in %.2f seconds ====\n", time.Since(begin).Seconds())
}

// Worker simula el trabajo de un trabajador, que toma un número aleatorio de segundos entre 3 y 5 para completar su tarea
// Recibe el nombre del trabajador, el tipo de trabajo y el generador de números aleatorios
func Worker(name, job string, r *rand.Rand) {
	fmt.Printf("Worker %s began to work on %s\n", name, job) // Imprimimos que el trabajador ha comenzado a trabajar

	// Generamos un tiempo aleatorio entre 3 y 5 segundos para el trabajo
	randomTime := r.Intn(3) + 3 // Intn(3) devuelve un número entre 0 y 2. Al sumarle 3, el rango es de 3 a 5

	// El ciclo for simula que el trabajador está trabajando durante un tiempo radom en segundos
	// El ciclo imprime el progreso del trabajador de 1 a tiempo random
	for i := 1; i <= randomTime; i++ {
		// Imprime el progreso del trabajador en cada iteración del ciclo
		fmt.Printf("Worker %s is working %d\n", name, i)
		time.Sleep(1 * time.Second) // Simula que el trabajador está trabajando durante 1 segundo por cada iteración
	}

	// Imprime cuando el trabajador termina su tarea
	fmt.Printf("Worker %s finished working on %s\n", name, job)
}

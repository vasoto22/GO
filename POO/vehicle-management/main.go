package main

import (
	"vehicle-management/models"
	"vehicle-management/services"
	"vehicle-management/utils"
)

func main() {
	// Crear instancias de vehículos
	car := models.Car{
		Vehicle: models.Vehicle{
			ID:    "1",
			Make:  "Toyota",
			Model: "Corolla",
			Year:  2020,
		},
		Doors: 4,
	}

	truck := models.Truck{
		Vehicle: models.Vehicle{
			ID:    "2",
			Make:  "Ford",
			Model: "F-150",
			Year:  2021,
		},
		PayloadCapacity: 5000,
	}

	// Crear un servicio de vehículos y agregar los vehículos
	vehicleService := services.VehicleService{}
	vehicleService.AddVehicle(car)
	vehicleService.AddVehicle(truck)

	// Iniciar todos los vehículos
	utils.LogMessage("Iniciando vehículos...")
	vehicleService.StartAllVehicles()
}

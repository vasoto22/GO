package services

import (
	"fmt"

	"vehicle-management/interfaces"
)

// Servicio que maneja la colección de vehículos
type VehicleService struct {
	Vehicles []interfaces.Vehicle
}

// Agregar un vehículo al servicio
func (vs *VehicleService) AddVehicle(v interfaces.Vehicle) {
	vs.Vehicles = append(vs.Vehicles, v)
}

// Iniciar todos los vehículos en el servicio
func (vs *VehicleService) StartAllVehicles() {
	for _, vehicle := range vs.Vehicles {
		fmt.Println(vehicle.Start())
	}
}

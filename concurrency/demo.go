package concurrency

import (
	"fmt"
	"time"
)

func GenerateReports() {

	reportChannel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		reportChannel <- "Reporte de Ventas generado"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		reportChannel <- "Reporte de Inventario generado"
	}()

	fmt.Println(<-reportChannel)
	fmt.Println(<-reportChannel)
}

package reports

import "fmt"

// ReportService genera reportes del sistema.
type ReportService struct{}

// NewReportService crea una instancia del servicio.
func NewReportService() *ReportService {
	return &ReportService{}
}

// GenerateSalesReport genera un reporte simple de ventas.
func (r *ReportService) GenerateSalesReport(total float64) {

	fmt.Println("=== REPORTE DE VENTAS ===")
	fmt.Printf("Ventas totales: %.2f\n", total)
}

// GenerateInventoryReport genera un reporte de inventario.
func (r *ReportService) GenerateInventoryReport(stock int) {

	fmt.Println("=== REPORTE DE INVENTARIO ===")
	fmt.Printf("Stock disponible: %d\n", stock)
}

package models

type Pago struct {
    ID       uint    `gorm:"primaryKey"`
    PedidoID uint    `gorm:"not null"`
    Monto    float64 `gorm:"not null"`
    Metodo   string  `gorm:"not null"` 
    Estado   string  `gorm:"not null"` 
}

package handler

type Metrics struct {
	ID    string   // имя метрики
	MType string   // параметр, принимающий значение gauge или counter
	Delta *int64   // значение метрики в случае передачи counter
	Value *float64 // значение метрики в случае передачи gauge
}

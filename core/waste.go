package core

type Waste struct {
	WasteType       string
	RecyclingMethod string
	Quantity        float64
	Location        string
	Status          string
}

var (
	WasteData       []Waste
	TriggerShowData bool   = true
	Version         string = "v1.2.2"
)

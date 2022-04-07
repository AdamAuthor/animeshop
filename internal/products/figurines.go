package products

type Figurines struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Manufacturer string  `json:"manufacturer"`
	Material     string  `json:"material"`
	WeightInG    float32 `json:"weight_in_g"`
	SizeInCm     int     `json:"size_in_cm"`
	Rating       float32 `json:"rating"`
}

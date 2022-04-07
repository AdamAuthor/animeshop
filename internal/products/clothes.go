package products

type Clothes struct {
	ID        int     `json:"id"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Source    string  `json:"source"`
	Color     string  `json:"color"`
	Gender    string  `json:"gender"`
	Size      string  `json:"size"`
	WeightInG float32 `json:"weight_in_g"`
	Rating    float32 `json:"rating"`
}

package products

type Manga struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Author     string   `json:"author"`
	Genres     []string `json:"genres"`
	NumOfPages int      `json:"num_of_pages"`
	Price      int      `json:"price"`
	WeightInG  float32  `json:"weight_in_g"`
	Rating     float32  `json:"rating"`
}

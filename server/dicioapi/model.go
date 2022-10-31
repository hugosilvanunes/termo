package dicioapi

type DicioResponse struct {
	ID        int    `json:"id,omitempty"`
	Word      string `json:"word,omitempty"`
	Count     int    `json:"count,omitempty"`
	Character string `json:"character,omitempty"`
}

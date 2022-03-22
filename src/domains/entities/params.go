package entities

type Param struct {
	Pairing string `json:"pairing"`
	G      []byte `json:"g"`
	U      []byte `json:"u"`
}
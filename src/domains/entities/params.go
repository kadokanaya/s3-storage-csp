package entities

type Param struct {
	Pairing string `json:"paring"`
	G      []byte `json:"g"`
	U      []byte `json:"u"`
}
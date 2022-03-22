package entities

type ArtLog struct {
	HashedData [][]byte `json:"hashed_data"`
	Owner      string   `json:"owner"`
}

type Chal struct {
	ContentId string `json:"art_id"`
	C         int    `json:"ck"`
	K1        []byte `json:"k1"`
	K2        []byte `json:"k2"`
}

type Proof struct {
	Myu       []byte `json:"myu"`
	Gamma     []byte `json:"gamma"`
	ContentId string `json:"content_id"`
}

type Proofs struct {
	DataList []Proof `json:"proofs"`
	Total    int     `json:"total"`
}

type Chals struct {
	DataList []Chal `json:"data"`
	Total    int    `json:"total"`
}

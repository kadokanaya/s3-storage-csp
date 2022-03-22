package entities

type User struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	PubKey   []byte `json:"pub_key"`
	PrivKey  []byte `json:"priv_key"`
	Password string `json:"password"`
}

type UserRequest struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

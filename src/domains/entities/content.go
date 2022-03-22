package entities

type Content struct {
	Content     []byte   `json:"content"`
	MetaData    [][]byte `json:"meta_data"`
	HashedData  [][]byte `json:"hashed_data"`
	ContentName string   `json:"name"`
	SplitCount  int      `json:"split_count"`
	Owner       string   `json:"owner"`
	Id          string   `json:"id"`
	UserId      string   `json:"user_id"`
	ContentId   string   `json:"content_id"`
}

type Receipt struct {
	Id           string `gorm:"primary_key" json:"id"`
	ContentLogId string `json:"content_log_id"`
	ContentURL   string `json:"content_url"`
	FileName     string `json:"name"`
	UserId       string `json:"user_id"`
	ContentId    string `json:"content_id"`
}

type ReceiptFromBC struct {
	ContentName string   `json:"name"`
	Owner       string   `json:"owner"`
	ContentId   string   `json:"content_id"`
	HashedData  [][]byte `json:"hashed_data"`
}

func NewContent() *Content {
	return &Content{}
}

type ContentInStorage struct {
	Id       string   `json:"id"`
	File     []byte   `json:"file"`
	MetaData [][]byte `json:"meta_data"`
	FileName string   `json:"name"`
}

type ContentInBlockChain struct {
	MetaData   [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName   string   `json:"name"`
	SplitCount int      `json:"split_count"`
	Owner      string   `json:"owner"`
	ContentId  string   `json:"content_id"`
}

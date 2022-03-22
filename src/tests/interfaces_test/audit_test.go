package interfaces_test

import (
	"sp/src/asserts"
	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/interfaces/crypt"
	"testing"

	"github.com/Nik-U/pbc"
)

// AuditProofGenテスト
func TestAuditProofGen(t *testing.T) {
		Pairing: "",
		G:       []byte{},
		U:       []byte{},
	}
	Pg := crypt.NewAuditCrypt(param)
	testByte := []byte{1}
	contentA := &entities.Content{
		Content:     []byte{1, 2, 3, 4},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "",
		UserId:      "",
		ContentId:   "",
	}
	contentB := &entities.Content{
		Content:     []byte{},
		MetaData:    [][]byte{testByte},
		HashedData:  [][]byte{testByte},
		ContentName: "",
		SplitCount:  0,
		Owner:       "",
		Id:          "",
		UserId:      "",
		ContentId:   "",
	}
	chal := &entities.Chal{
		ContentId: "",
		C:         0,
		K1:        []byte{},
		K2:        []byte{},
	}
	_, err := Pg.AuditProofGen(chal, contentA, contentB)
	if err != nil {
		t.Fatal(err)
	}
}

// メタデータ作成テスト
func TestMakeMetaData(t *testing.T) {
	//* パラメータ生成
	params := pbc.GenerateA(uint32(160), uint32(512))
	pairing := params.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	privKey := pairing.NewZr().Rand()
	pubKey := pairing.NewG1().MulZn(g, privKey)
	p := &contracts.Param{
		Paring: params.String(),
		G:      g.Bytes(),
		U:      u.Bytes(),
	}

	//* メタデータ作成
	f, err := core.UseFileRead("./linux_logo.jpg")
	if err != nil {
		t.Fatal(err)
	}
	uc := &entities.Content{
		Content:     f,
		ContentName: "testName",
		SplitCount:  3,
	}

	user := &entities.User{
		PubKey:  pubKey.Bytes(),
		PrivKey: privKey.Bytes(),
	}

	content, err := crypt.CreateMetaData(uc, user, p)
	if err != nil {
		t.Fatal(err)
	}

	Pg := crypt.NewAuditCrypt(p)

	chal := &entities.Chal{
		ContentId: "",
		C:         2,
		K1:        []byte("K2VldQ4QswyO8FVWyLe6Nb+j+2o="),
		K2:        []byte("WJxJcgDXVLHJAN31lr/YzmrJ7U0="),
	}
	proof, err := Pg.AuditProofGen(chal, content, content)
	if err != nil {
		t.Fatal(err)
	}
	l, r, err := crypt.AuditVerify(p, user.PubKey, content, proof, chal)
	if err != nil {
		t.Fatal(l, r)
	}
	asserts.AssertEqual(t,l,r,"ペアリングの一致の確認")
}

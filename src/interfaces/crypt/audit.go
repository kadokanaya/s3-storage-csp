package crypt

import (
	"fmt"
	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/interfaces/contracts"
	"sp/src/usecases/port"

	"github.com/Nik-U/pbc"
)

type auditCrypt struct {
	Param *contracts.Param
}

func NewAuditCrypt(param *contracts.Param) port.AuditCrypt {
	return &auditCrypt{
		Param: param,
	}
}

func (pr *auditCrypt) AuditProofGen(
	chal *entities.Chal,
	content *entities.Content,
	contentLog *entities.Content,
) (*entities.Proof, error) {
	var myu *pbc.Element
	var gamma *pbc.Element
	pairing, err := pbc.NewPairingFromString(pr.Param.Paring)
	if err != nil {
		return nil, err
	}
	splitedFile, err := core.SplitSlice(content.Content, content.SplitCount)
	if err != nil {
		return nil, err
	}
	aTable, vTable := core.HashChallen(contentLog.SplitCount, int(chal.C), chal.K1, chal.K2, pairing)

	var MSum *pbc.Element
	if chal.C < 1 {
		return nil, fmt.Errorf("challengeの形が良くない")
	}
	for cIndex := 0; cIndex < int(chal.C); cIndex++ {
		meta := pairing.NewG1().SetBytes(content.MetaData[aTable[cIndex]])
		m := pairing.NewG1().SetFromHash(splitedFile[aTable[cIndex]])
		mm := core.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetBytes(mm)
		if cIndex == 0 {
			myu = pairing.NewZr().MulBig(vTable[cIndex], m.X())
			gamma = pairing.NewG1().PowZn(meta, vTable[cIndex])
		} else {
			myu = pairing.NewZr().Add(myu, pairing.NewZr().MulBig(vTable[cIndex], m.X()))
			gamma.Mul(gamma, pairing.NewG1().PowZn(meta, vTable[cIndex]))
		}
		if cIndex == 0 {
			MSum = pairing.NewG1().PowZn(M, vTable[cIndex])
		} else {
			MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[cIndex]))
		}
	}
	proof := &entities.Proof{
		Myu:       myu.Bytes(),
		Gamma:     gamma.Bytes(),
		ContentId: contentLog.ContentId,
	}
	return proof, nil

}

func CreateMetaData(uc *entities.Content, user *entities.User, param *contracts.Param) (*entities.Content, error) {
	pairing, err := pbc.NewPairingFromString(param.Paring)
	if err != nil {
		return nil, err
	}
	u := pairing.NewG1().SetBytes(param.U)
	splitedFile, err := core.SplitSlice(uc.Content, uc.SplitCount)
	if err != nil {
		return nil, err
	}
	privKey := pairing.NewZr().SetBytes(user.PrivKey)

	// メタデータの作成
	var metaData [][]byte
	var hashData [][]byte
	metaToHash := ""
	for i := 0; i < len(splitedFile); i++ {
		m := pairing.NewG1().SetFromHash(splitedFile[i])

		mm := core.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetBytes(mm)

		um := pairing.NewG1().PowBig(u, m.X())
		temp := pairing.NewG1().Mul(um, M)
		meta := pairing.NewG1().MulZn(temp, privKey)

		metaData = append(metaData, meta.Bytes())
		metaToHash = metaToHash + meta.String()
		hashData = append(hashData, mm)
	}

	return &entities.Content{
		Content:     uc.Content,
		MetaData:    metaData,
		HashedData:  hashData,
		ContentName: uc.ContentName,
		SplitCount:  uc.SplitCount,
		Owner:       uc.Owner,
		Id:          "",
		UserId:      uc.UserId,
		ContentId:   "",
	}, nil
}

func AuditVerify(params *contracts.Param, pubKeyByte []byte, content *entities.Content, proof *entities.Proof, chal *entities.Chal) ([]byte, []byte, error) {
	pairing, _ := pbc.NewPairingFromString(params.Paring)
	aTable, vTable := core.HashChallen(content.SplitCount, chal.C, chal.K1, chal.K2, pairing)
	g := pairing.NewG1().SetBytes(params.G)
	pubKey := pairing.NewG1().SetBytes(pubKeyByte)
	u := pairing.NewG1().SetBytes(params.U)
	myuT := pairing.NewZr().SetBytes(proof.Myu)
	gammaT := pairing.NewG1().SetBytes(proof.Gamma)
	var MSum *pbc.Element
	for c := 0; c < chal.C; c++ {
		M := pairing.NewG1().SetBytes(content.HashedData[aTable[c]])
		if c == 0 {
			MSum = pairing.NewG1().PowZn(M, vTable[c])
		} else {
			MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[c]))
		}
	}
	uProof := pairing.NewG1().PowZn(u, myuT)
	right_hand := pairing.NewG1().Mul(uProof, MSum)
	pairingLeft := pairing.NewGT().Pair(gammaT, g)
	pairingRight := pairing.NewGT().Pair(right_hand, pubKey)
	return pairingLeft.Bytes(), pairingRight.Bytes(), nil

}

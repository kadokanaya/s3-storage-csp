package drivers

// func CreateMetaData(uploadFile *structure.UploadFile, para *structure.Params, user *structure.User) gin.HandlerFunc {
// 	// ペアリングを変換
// 	pairing, _ := pbc.NewPairingFromString(para.Pairing)
// 	u := pairing.NewG1().SetBytes(para.U)
// 	n := 3

// 	//ファイルの分割
// 	splitedFile, _ := tool.SplitSlice(inputFile, n)

// 	// メタデータの作成
// 	var metaData [][]byte
// 	var MData [][]byte
// 	metaToHash := ""
// 	for i := 0; i < len(splitedFile); i++ {
// 		m := pairing.NewG1().SetFromHash(splitedFile[i])

// 		mm := tool.GetBinaryBySHA256(m.X().String())
// 		log.Print("mm\n")
// 		log.Print(mm)
// 		M := pairing.NewG1().SetBytes(mm) 

// 		um := pairing.NewG1().PowBig(u, m.X())
// 		temp := pairing.NewG1().Mul(um, M)
// 		meta := pairing.NewG1().MulZn(temp, privKey)

// 		metaData = append(metaData, meta.Bytes())
// 		metaToHash = metaToHash + meta.String()
// 		MData = append(MData, mm)
// 	}

// 	return metaData, MData, string(tool.MD5(metaToHash))
// }
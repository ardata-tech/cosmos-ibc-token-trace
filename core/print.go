package core

import "fmt"

type PrintService struct {
}

func NewPrintService() *PrintService {
	return &PrintService{}
}

func (p PrintService) PrintTokenInfo(tokenInfos []TokenInfo, totalAmounts TotalAmounts) {
	for _, info := range tokenInfos {
		fmt.Printf("%s:\n", info.ChainName)
		fmt.Printf("%s, %s, %f, %v\n", info.Denom, info.OriginalDenom, info.Amount, info.Path)
	}

	fmt.Println("TOTAL AMOUNTS:")
	fmt.Printf("%s, %f\n", originalDenoms["dAsset"], totalAmounts.DAssetAmount)
	fmt.Printf("%s, %f\n", originalDenoms["lAsset"], totalAmounts.LAssetAmount)
}

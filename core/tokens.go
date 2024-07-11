package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TokenInfo struct {
	ChainName     string
	Denom         string
	OriginalDenom string
	Amount        float64
	Path          []string
}

type TotalAmounts struct {
	DAssetAmount float64
	LAssetAmount float64
}

type DenomTrace struct {
	BaseDenom string `json:"base_denom"`
	Path      string `json:"path"`
}

type TokenService struct {
	ChainEndpoints map[string]string
	OriginalDenoms map[string]string
}

func NewTokenService(chainEndpoints map[string]string, originalDenoms map[string]string) *TokenService {
	return &TokenService{
		ChainEndpoints: chainEndpoints,
		OriginalDenoms: originalDenoms,
	}
}
func (t TokenService) fetchBalances(chain, address string) (map[string]float64, error) {
	url := fmt.Sprintf("%s/bank/v1beta1/balances/%s", t.ChainEndpoints[chain], address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Balances []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balances"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	balances := make(map[string]float64)
	for _, balance := range result.Balances {
		var amount float64
		fmt.Sscanf(balance.Amount, "%f", &amount)
		balances[balance.Denom] = amount
	}
	return balances, nil
}

func (t TokenService) fetchDenomTrace(chain, denom string) (DenomTrace, error) {
	var denomTrace DenomTrace
	url := fmt.Sprintf("%s/ibc/apps/transfer/v1beta1/denom_traces/%s", chainEndpoints[chain], denom)
	resp, err := http.Get(url)
	if err != nil {
		return denomTrace, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&denomTrace); err != nil {
		return denomTrace, err
	}
	return denomTrace, nil
}

func (t TokenService) GetTokenInfo(addresses map[string]string) ([]TokenInfo, TotalAmounts) {
	var tokenInfos []TokenInfo
	totalAmounts := TotalAmounts{}

	for chain, addr := range addresses {
		balances, err := t.fetchBalances(chain, addr)
		if err != nil {
			fmt.Printf("Error fetching balances for %s: %v\n", chain, err)
			continue
		}

		for denom, amount := range balances {
			tokenInfo := TokenInfo{
				ChainName: chain,
				Denom:     denom,
				Amount:    amount,
				Path:      []string{chain},
			}

			if strings.HasPrefix(denom, "ibc/") {
				denomTrace, err := t.fetchDenomTrace(chain, strings.TrimPrefix(denom, "ibc/"))
				if err != nil {
					fmt.Printf("Error fetching denom trace for %s: %v\n", denom, err)
					continue
				}
				tokenInfo.OriginalDenom = denomTrace.BaseDenom
				tokenInfo.Path = append(tokenInfo.Path, strings.Split(denomTrace.Path, "/")...)
				totalAmounts.DAssetAmount += amount
			} else if denom == originalDenoms["dAsset"] {
				tokenInfo.OriginalDenom = denom
				totalAmounts.DAssetAmount += amount
			} else if denom == originalDenoms["lAsset"] {
				tokenInfo.OriginalDenom = denom
				totalAmounts.LAssetAmount += amount
			}

			tokenInfos = append(tokenInfos, tokenInfo)
		}
	}

	return tokenInfos, totalAmounts
}

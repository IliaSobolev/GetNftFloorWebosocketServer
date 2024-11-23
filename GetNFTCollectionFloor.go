package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FloorPriceResponse struct {
	Data struct {
		AlphaNftCollectionStats struct {
			FloorPrice float64 `json:"floorPrice"`
		} `json:"alphaNftCollectionStats"`
	} `json:"data"`
}

func GetNFTCollectionFloor(nftCollectionAddress string) (float64, error) {
	if nftCollectionAddress == "" {
		return 0, fmt.Errorf("nftCollectionAddress is empty")
	}

	query := `query AlphaNftCollectionStats($address: String!) { alphaNftCollectionStats(address: $address) { floorPrice } }`

	reqBody := map[string]interface{}{
		"query": query,
		"variables": map[string]interface{}{
			"address": nftCollectionAddress,
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post("https://api.getgems.io/graphql", "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	var floorPriceResp FloorPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&floorPriceResp); err != nil {
		return 0, err
	}
	return floorPriceResp.Data.AlphaNftCollectionStats.FloorPrice, nil
}

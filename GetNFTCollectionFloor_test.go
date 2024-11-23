package main

import (
	"testing"
)

func TestGetNFTCollectionFloorWithWallet(t *testing.T) {
	walletAddress := "EQAOQdwdw8kGftJCSFgOErM1mBjYPe4DBPq8-AhF6vr9si5N"
	floor, err := GetNFTCollectionFloor(walletAddress)
	if floor == 0 || err != nil {
		t.Fatalf("GetNFTCollectionFloor failed: %v", err)
	}
}

func TestGetNFTCollectionFloorWithoutWallet(t *testing.T) {
	floor, err := GetNFTCollectionFloor("")
	if floor != 0 || err == nil {
		t.Fatalf("GetNFTCollectionFloorWithoutWallet failed because error is nil or floor is 0")
	}
}

//func TestGetNFTCollectionFloorWithInvalidAddress(t *testing.T) {
//	walletAddress := "EQAOQdwdw8kGftJcSFgOErM1mBjYPe4DRPq7-AhF6vr9si5N"
//	floor, err := GetNFTCollectionFloor(walletAddress)
//	if floor != 0 || err == nil {
//		t.Fatalf("GetNFTCollectionFloorWithInvalidAddress failed: %v, floor is %f", err, floor)
//	}
//}

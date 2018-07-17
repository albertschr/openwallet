/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package merchant

import (
	"testing"
	"path/filepath"
	"log"
	"github.com/blocktree/OpenWallet/openwallet"
)


func init() {
	nodeConfig = NodeConfig{
		NodeKey:         "",
		PublicKey:       "dajosidjaiosjdioajsdioajsdiowhefi",
		PrivateKey:      "",
		MerchantNodeURL: "ws://192.168.30.4:8084/websocket",
		NodeID:          1,
		CacheFile:       filepath.Join(merchantDir, cacheFile),
	}

}

func TestGetImportAddress(t *testing.T) {

	m, err := NewMerchantNode(nodeConfig)
	if err != nil {
		t.Errorf("GetChargeAddressVersion failed unexpected error: %v", err)
	}

	walletID := "sss"
	wallet, err := m.GetMerchantWalletByID(walletID)
	if err != nil {
		log.Printf("unexpected error: %v", err)
		return
	}

	db, err := wallet.OpenDB()
	if err != nil {
		log.Printf("unexpected error: %v", err)
		return
	}
	defer db.Close()

	var addresses []*openwallet.Address
	err = db.All(&addresses)
	if err != nil {
		log.Printf("unexpected error: %v", err)
		return
	}

	for _, a := range addresses {
		log.Printf("address = %s\n", a.Address)
	}

}
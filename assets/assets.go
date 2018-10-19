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

package assets

import (
	"strings"

	"github.com/blocktree/OpenWallet/log"

	"github.com/blocktree/OpenWallet/assets/bitcoin"
	"github.com/blocktree/OpenWallet/assets/bitcoincash"
	"github.com/blocktree/OpenWallet/assets/bopo"
	"github.com/blocktree/OpenWallet/assets/bytom"
	"github.com/blocktree/OpenWallet/assets/cardano"
	"github.com/blocktree/OpenWallet/assets/decred"
	"github.com/blocktree/OpenWallet/assets/ethereum"
	"github.com/blocktree/OpenWallet/assets/hypercash"
	"github.com/blocktree/OpenWallet/assets/iota"
	"github.com/blocktree/OpenWallet/assets/litecoin"
	"github.com/blocktree/OpenWallet/assets/qtum"
	"github.com/blocktree/OpenWallet/assets/sia"
	"github.com/blocktree/OpenWallet/assets/stc2345"
	"github.com/blocktree/OpenWallet/assets/tezos"
	"github.com/blocktree/OpenWallet/assets/tron"
	"github.com/blocktree/OpenWallet/assets/nebulasio"
)

//WalletManager 钱包管理器
type WalletManager interface {
	//初始化配置流程
	InitConfigFlow() error
	//查看配置信息
	ShowConfig() error
	//创建钱包流程
	CreateWalletFlow() error
	//创建地址流程
	CreateAddressFlow() error
	//汇总钱包流程
	SummaryFollow() error
	//备份钱包流程
	BackupWalletFlow() error
	//查看钱包列表，显示信息
	GetWalletList() error
	//发送交易
	TransferFlow() error
	//恢复钱包
	RestoreWalletFlow() error
}

// 节点管理接口
type NodeManager interface {
	// GetNodeStatus 节点状态
	GetNodeStatus(string) error
	// StartNodeFlow 创建节点
	CreateNodeFlow(string) error
	// StartNodeFlow 开启节点
	StartNodeFlow(string) error
	// StopNodeFlow 关闭节点
	StopNodeFlow(string) error
	// RestartNodeFlow 重启节点
	RestartNodeFlow(string) error
	// RemoveNodeFlow 移除节点
	RemoveNodeFlow(string) error
	// LogsNodeFlow 日志
	LogsNodeFlow(string) error

	// //LoginNode 登陆节点
	// LoginNode() error
	// //ShowNodeInfo 显示节点信息
	// ShowNodeInfo() error
}

// 配置管理接口
type ConfigManager interface {
	//SetConfigFlow 初始化配置流程
	SetConfigFlow(subCmd string) error
	//ShowConfigInfo 查看配置信息
	ShowConfigInfo(subCmd string) error
}

//钱包管理器组
var Managers = make(map[string]interface{})

// RegWMD makes a WalletManager available by the name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func RegWMD(name string, manager interface{}) {
	if manager == nil {
		panic("WalletManager: Register adapter is nil")
	}
	if _, ok := Managers[name]; ok {
		log.Error("WalletManager: Register called twice for adapter ", name)
		return
	}
	Managers[name] = manager
}

// GetWMD 根据币种类型获取已注册的管理者
func GetWMD(symbol string) interface{} {
	manager, ok := Managers[symbol]
	if !ok {
		return nil
	}
	return manager
}

//注册钱包管理工具
func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	RegWMD(strings.ToLower(cardano.Symbol), &cardano.WalletManager{})
	RegWMD(strings.ToLower(bytom.Symbol), &bytom.WalletManager{})
	RegWMD(strings.ToLower(bopo.Symbol), &bopo.WalletManager{})
	RegWMD(strings.ToLower(bitcoincash.Symbol), &bitcoincash.WalletManager{})
	RegWMD(strings.ToLower(sia.Symbol), &sia.WalletManager{})
	RegWMD(strings.ToLower(ethereum.Symbol), ethereum.NewWalletManager())
	RegWMD(strings.ToLower(stc2345.Symbol), stc2345.NewWalletManager())
	RegWMD(strings.ToLower(bitcoin.Symbol), bitcoin.NewWalletManager())
	RegWMD(strings.ToLower(hypercash.Symbol), hypercash.NewWalletManager())
	RegWMD(strings.ToLower(iota.Symbol), &iota.WalletManager{})
	RegWMD(strings.ToLower(tezos.Symbol), tezos.NewWalletManager())
	RegWMD(strings.ToLower(litecoin.Symbol), litecoin.NewWalletManager())
	RegWMD(strings.ToLower(qtum.Symbol), qtum.NewWalletManager())
	RegWMD(strings.ToLower(decred.Symbol), decred.NewWalletManager())
	RegWMD(strings.ToLower(tron.Symbol), tron.NewWalletManager())
	RegWMD(strings.ToLower(nebulasio.Symbol), nebulasio.NewWalletManager())
}

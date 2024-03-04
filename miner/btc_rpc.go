package miner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/rpcclient"
)

type BtcRpc struct {
	Network  string `json:"network" mapstructure:"network"`
	Host     string `json:"host" mapstructure:"host"`
	Endpoint string `json:"endpoint" mapstructure:"endpoint"`
	User     string `json:"user" mapstructure:"user"`
	Pass     string `json:"pass" mapstructure:"pass"`

	client      *rpcclient.Client
	blockHeight int64
}

func (br *BtcRpc) Connect() error {

	disableTLS := false
	if br.Endpoint == "http" || br.Endpoint == "ws" {
		disableTLS = true
	}
	httpPostMode := false
	if br.Endpoint == "http" || br.Endpoint == "https" {
		httpPostMode = true
	}

	connCfg := &rpcclient.ConnConfig{
		Host:         br.Host,
		Endpoint:     br.Endpoint,
		User:         br.User,
		Pass:         br.Pass,
		Params:       br.Network,
		HTTPPostMode: httpPostMode,
		DisableTLS:   disableTLS,
	}

	if !disableTLS {
		btcdHomeDir := btcutil.AppDataDir("btcd", false)
		certs, err := os.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
		if err != nil {
			log.Fatal(err)
		}

		connCfg.Certificates = certs
	}

	fmt.Println("btc rpc client connecting...")
	// fmt.Println(connCfg)

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		fmt.Println("btc rpc client error", err)
		return err
	}

	br.client = client
	fmt.Println("btc rpc client connected")
	return nil
}

func (br *BtcRpc) Disconnect() {
	if br.client != nil {
		br.client.Shutdown()
	}
}

func (br *BtcRpc) GetBlockHeight() (int64, error) {
	blockCount, err := br.client.GetBlockCount()
	if err != nil {
		return -1, err
	}

	br.blockHeight = blockCount
	// fmt.Printf("block count: %d\n", blockCount)

	return blockCount, nil
}

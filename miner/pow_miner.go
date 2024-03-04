package miner

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type PowMiner struct {
	Address    string `json:"address" mapstructure:"address"`
	TxID       string `json:"txid" mapstructure:"txid"`
	P          string `json:"protocol" mapstructure:"protocol"`
	Tick       string `json:"tick" mapstructure:"tick"`
	Amt        string `json:"amt" mapstructure:"amt"`
	Difficulty uint8  `json:"difficulty" mapstructure:"difficulty"`

	AutoSend bool   `json:"auto_send" mapstructure:"auto_send"`
	GasPrice uint64 `json:"gas_price" mapstructure:"gas_price"`
}

func (pm *PowMiner) Transfer(amt string, to string) (*InscribeTransfer, []byte) {
	inscription := InscribeTransfer{
		P:    pm.P,
		Op:   "transfer",
		Tick: pm.Tick,
		Amt:  amt,
	}
	if to != "" {
		inscription.To = to
	}

	inscriptionText, err := json.Marshal(inscription)
	if err != nil {
		return nil, []byte{}
	}

	return &inscription, inscriptionText
}

func (pm *PowMiner) Mint() (string, *InscribeMint, []byte) {
	nonce := uuid.New().ID()

	inscription := InscribeMint{
		P:        pm.P,
		Op:       "mint",
		Tick:     pm.Tick,
		Amt:      pm.Amt,
		Solution: fmt.Sprintf("%s:%s:%s:%d", pm.Tick, pm.Address, pm.TxID, nonce),
	}

	inscriptionText, err := json.Marshal(inscription)
	if err != nil {
		return "", nil, []byte{}
	}

	h := sha256.New()
	h.Write([]byte(inscriptionText))
	solution := fmt.Sprintf("%x", h.Sum(nil))

	prefix := strings.Repeat("0", int(pm.Difficulty))
	if !strings.HasPrefix(solution, prefix) {
		// fmt.Printf("%s unmatched difficulty, continue...\n", solution)
		return "", nil, []byte{}
	}

	return solution, &inscription, inscriptionText
}

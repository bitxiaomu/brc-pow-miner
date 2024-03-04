package miner

import (
	"fmt"
	"strings"
	"time"
)

type Miner struct {
	PowMiner PowMiner `json:"pow_miner" mapstructure:"pow_miner"`
	BtcRpc   BtcRpc   `json:"btc_rpc" mapstructure:"btc_rpc"`

	ExitMainLoop chan bool
}

func (mn *Miner) MainLoop() {
	ticker := time.NewTicker(time.Microsecond * 20)

	updateTms := 0.0
	updateTm := time.Now()

	for {
		select {
		case <-mn.ExitMainLoop:
			fmt.Println("exit main loop")
			return

		case <-ticker.C:
			updateTms++
			solution, inscription, inscriptionText := mn.PowMiner.Mint()

			if inscription != nil {
				deltaTm := time.Since(updateTm).Seconds()
				// fmt.Println(updateTms, deltaTm, int(updateTms/deltaTm))

				fmt.Println(strings.Repeat("-", 80))
				fmt.Printf("HashRate: %d c/s %d runs %d sec\n", uint64(updateTms/deltaTm), uint64(updateTms), uint64(deltaTm))
				fmt.Printf("Inscription: %v\n", inscription)
				fmt.Printf("Found Solution: %s\n", solution)
				fmt.Println("Inscription Text:")
				fmt.Println("")
				fmt.Println(string(inscriptionText))
				fmt.Println("")
				fmt.Println(strings.Repeat("-", 80))

				updateTms = 0.0
				updateTm = time.Now()
			}
		}
	}
}

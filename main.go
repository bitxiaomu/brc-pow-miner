package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"brc-pow-miner/miner"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var minerIns miner.Miner

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./etc")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("read config failed with %s\n", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&minerIns); err != nil {
		log.Printf("unmarshal config failed with %s\n", err)
		os.Exit(1)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}

func main() {
	initConfig()

	minerIns.BtcRpc.Connect()
	defer minerIns.BtcRpc.Disconnect()

	minerIns.ExitMainLoop = make(chan bool, 1)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	// ticker := time.NewTicker(time.Duration(cfg.MonitorInterval) * time.Second)
	ticker := time.NewTicker(time.Second * 5)

	startTm := time.Now()

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("miner running %ds\n", uint64(math.Ceil(time.Since(startTm).Seconds())))
				minerIns.BtcRpc.GetBlockHeight()

			case sig := <-sigs:
				fmt.Println("receive signal:", sig)
				done <- true
			}
		}
	}()

	go minerIns.MainLoop()

	fmt.Printf("miner started at %v\n", startTm)
	<-done
	fmt.Printf("miner exiting, run %v\n", time.Since(startTm))
}

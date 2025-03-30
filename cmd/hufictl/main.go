package main

import (
	hq2 "github.com/conplementag/cops-hq/v2/pkg/hq"
	"github.com/huf1/hufictl/internal/hufictl/orchestration/lottery"
)

func main() {
	hq := hq2.NewQuiet("hufictl", "0.1.0", "hufictl.log")
	createCommands(hq)

	err := hq.Run()

	if err != nil {
		panic(err)
	}
}

func createCommands(hq hq2.HQ) {
	lotteryCommandGroup := hq.GetCli().AddBaseCommand("lottery", "", "", nil)

	lotteryCommandGroup.AddCommand("spin", "", "", func() {
		lottery.NewOrchestrator().Spin()
	})
}

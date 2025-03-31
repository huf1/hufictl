package main

import (
	"fmt"
	"github.com/conplementag/cops-hq/v2/pkg/hq"
	"github.com/huf1/hufictl/internal/hufictl/orchestration/lottery"
)

// see .goreleaser.yaml ldflags
var version = "local"
var commit = "local"
var date = "local"

func main() {
	copshq := hq.NewQuiet("hufictl", version, "hufictl.log")
	createCommands(copshq)

	err := copshq.Run()

	if err != nil {
		panic(err)
	}
}

func createCommands(hq hq.HQ) {
	lotteryCommandGroup := hq.GetCli().AddBaseCommand("lottery", "", "", nil)

	lotteryCommandGroup.AddCommand("spin", "", "", func() {
		lottery.NewOrchestrator().Spin()
	})

	hq.GetCli().AddBaseCommand("version-full", "", "", func() {
		fmt.Printf("Version: %s\nCommit: %s\nDate: %s\n", version, commit, date)
	})

}

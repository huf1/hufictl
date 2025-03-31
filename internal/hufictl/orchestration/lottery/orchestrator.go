package lottery

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/huf1/hufictl/internal/hufictl/common/helpers"
	"math/rand/v2"
	"slices"
)

type Orchestrator struct{}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (orchestrator *Orchestrator) Spin() {
	blueColor := color.New(color.FgBlue)
	yellowColor := color.New(color.FgHiYellow)

	fortyNine := 49
	nine := 9

	var mainNumbers []int
	var additionalNumber []int

	for i := 0; i < 6; i++ {
		for {
			nextRandom := rand.IntN(fortyNine)
			if !slices.Contains(mainNumbers, nextRandom) {
				mainNumbers = append(mainNumbers, rand.IntN(fortyNine-1)+1)
				break
			}
		}
	}

	additionalNumber = append(additionalNumber, rand.IntN(nine))

	fmt.Println("Your next winning lottery numbers are:")
	_, err := blueColor.Printf("%d ", mainNumbers)
	helpers.FatalOnError(err)
	_, err = yellowColor.Println(additionalNumber)
	helpers.FatalOnError(err)
}

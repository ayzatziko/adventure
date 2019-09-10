package adventure

import (
	"fmt"
	"io"
	"log"

	"github.com/ayzatziko/algos"
)

type Adventure struct {
	playerPos int
	exitLoc   int
	g         *algos.Graph
}

func Read(r io.Reader) (*Adventure, error) {
	g := algos.ReadGraph(r)

	adv := Adventure{g: g}
	fmt.Fscanf(r, "%d %d\n", &adv.playerPos, &adv.exitLoc)
	return &adv, nil
}

func (adv *Adventure) Move(r io.Reader) error {
	dst := 0

	fmt.Fscanf(r, "%d\n", &dst)

	if !algos.HasPath(adv.g, adv.playerPos, dst) {
		return fmt.Errorf("adventure: cannot move from %d to %d", adv.playerPos, dst)
	}

	adv.playerPos = dst
	return nil
}

func (adv *Adventure) IsExit() bool {
	return adv.playerPos == adv.exitLoc
}

func (adv *Adventure) Run(r io.Reader) {
	for {
		if err := adv.Move(r); err != nil {
			log.Println(err)
		}

		if adv.IsExit() {
			fmt.Println("END")
			break
		}
	}
}

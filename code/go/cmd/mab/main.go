package main

import (
	"context"
	"fmt"
	"log"
	"github.com/stitchfix/mab"
	"github.com/stitchfix/mab/numint"
)

func main() {

	rewards := map[string][]mab.Dist{
		"us": {
			mab.Beta(40, 474),
			mab.Beta(64, 730),
			mab.Beta(71, 818),
		},
		"uk": {
			mab.Beta(25, 254),
			mab.Beta(100, 430),
			mab.Beta(30, 503),
		},
	}

	bandit := mab.Bandit{
		RewardSource: &mab.ContextualRewardStub{rewards},
		Strategy:     mab.NewThompson(numint.NewQuadrature()),
		Sampler:      mab.NewSha1Sampler(),
	}

	result, err := bandit.SelectArm(context.Background(), "user_id:12345", "us")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
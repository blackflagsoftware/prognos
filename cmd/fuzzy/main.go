package main

import (
	"fmt"
	"strings"

	th "github.com/blackflagsoftware/prognos/internal/entities/transactionhistory"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	categories := []string{"THE HOT TIN ROOF", "CULVERS", "THE MEGAPLEX", "MCDONALD'S", "TACO BELL", "GOOGLE*CLOUD SOMETHING", "WAL-MART 2306 2306"}
	for _, c := range categories {
		search(c)
	}
}

func search(categoryStr string) {
	transactionHistory := make(map[string]int)
	ths := th.InitStorage()
	thm := th.NewTransactionHistoryManager(ths)
	thm.Read(transactionHistory)

	categoryId := 0 // set to unknown
	matchSplit := strings.Split(categoryStr, " ")
	match := matchSplit[0] // get the first word
	if len(match) < 5 {
		// if there is another word add it
		if len(matchSplit) > 1 {
			match = match + " " + matchSplit[1]
		}
	}
	keys := []string{}
	for k := range transactionHistory {
		keys = append(keys, k)
	}
	ranks := fuzzy.RankFindFold(match, keys)
	key := ""
	minDistance := 0
	if len(ranks) > 0 {
		minDistance = ranks[0].Distance
		key = ranks[0].Target
		categoryId = transactionHistory[key]
	}
	// if len(distances) > 0 {
	// avg := int(sumDistances / len(distances))
	// if avg < 15 {
	// 	categoryId = transactionHistory[key]
	// 	return
	// }
	// }
	// 	}
	// }
	fmt.Printf("catIn: %s, key: %s; catId: %d; minDistance: %d\n", match, key, categoryId, minDistance)
	return
}

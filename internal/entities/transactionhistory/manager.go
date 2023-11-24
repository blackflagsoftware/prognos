package transactionhistory

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type (
	TransactionHistoryDataAdapter interface {
		Read(map[string]int) error
		Create(string, int) error
	}

	TransactionHistoryManager struct {
		transactionHistoryDataAdapter TransactionHistoryDataAdapter
	}
)

func NewTransactionHistoryManager(ta TransactionHistoryDataAdapter) *TransactionHistoryManager {
	return &TransactionHistoryManager{transactionHistoryDataAdapter: ta}
}

func (t *TransactionHistoryManager) Read(transactionHistory map[string]int) error {
	return t.transactionHistoryDataAdapter.Read(transactionHistory)
}

func (t *TransactionHistoryManager) Create(text string, categoryId int) error {
	text = strings.TrimSpace(text)
	minLen := 20
	if len(text) < 20 {
		minLen = len(text)
	}
	text = string(text[:minLen])
	return t.transactionHistoryDataAdapter.Create(text, categoryId)
}

func (t *TransactionHistoryManager) FindCategory(categoryStr string) (categoryId int) {
	transactionHistory := make(map[string]int)
	t.Read(transactionHistory)

	categoryId = 0 // set to unknown
	compareLen := 20
	if len(categoryStr) < 20 {
		compareLen = len(categoryStr)
	}
	compare := categoryStr[:compareLen]
	// compare whole, just in case it is a common transaction
	// by using RankMatch if the categoryStr (coming in) matches the transaction history, 1 for 1
	// send back the saved transaction history's saved category id
	for k, v := range transactionHistory {
		rank := fuzzy.RankMatch(k, compare)
		if rank == 0 {
			categoryId = v
			return
		}
	}
	// load all the keys of transactionHistory into an array for comparison
	keys := []string{}
	for k := range transactionHistory {
		keys = append(keys, k)
	}
	// split all the words from the category coming in, if the first word is less than 5 characters long
	// add the next word if applicable
	matchSplit := strings.Split(categoryStr, " ")
	match := matchSplit[0] // get the first word
	if len(match) < 5 {
		// if there is another word add it
		if len(matchSplit) > 1 {
			match = match + " " + matchSplit[1]
		}
	}
	// this will use the Levenshtein algorithm to see if any of the transactionHistory "keys" match, if so,
	// take the first one
	ranks := fuzzy.RankFindFold(match, keys)
	if len(ranks) > 0 {
		categoryId = transactionHistory[ranks[0].Target]
	}
	return
}

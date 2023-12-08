package main

import (
	"bufio"
	"fmt"
	"sort"
	"os"
	"strconv"
	"strings"
)

type HandRanking int

const (
	HighCard HandRanking = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardRanks = map[rune]int{  
	'A': 13,  
	'K': 12,  
	'Q': 11,  
	'J': 10,  
	'T': 9,  
	'9': 8,  
	'8': 7,  
	'7': 6,  
	'6': 5,  
	'5': 4,  
	'4': 3,  
	'3': 2,  
	'2': 1,  
}  

type PokerHand struct {
	Ranking HandRanking
	Hand    string
	Bid     int
}

func (p PokerHand) compare(other PokerHand) int {
	if p.Ranking < other.Ranking {
		return -1
	} else if p.Ranking > other.Ranking {
		return 1
	} else {
		for i := 0; i < len(p.Hand); i++ {
			if cardRanks[rune(p.Hand[i])] < cardRanks[rune(other.Hand[i])] {  
				return -1  
			} else if cardRanks[rune(p.Hand[i])] > cardRanks[rune(other.Hand[i])] {  
				return 1  
			}  
		}
		return 0
	}
}


func (p *PokerHand) setHand(handStr string, bid int) {
	ranks := make(map[rune]int)
	for _, card := range handStr {
		ranks[card]++
	}

	counts := make([]int, 0, len(ranks))
	for _, count := range ranks {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	var ranking HandRanking
	switch {
	case len(counts) == 1:
		ranking = FiveOfAKind
	case len(counts) == 2:
		if counts[0] == 4 {
			ranking = FourOfAKind
		} else {
			ranking = FullHouse
		}
	case len(counts) == 3:
		if counts[0] == 3 {
			ranking = ThreeOfAKind
		} else {
			ranking = TwoPair
		}
	case len(counts) == 4:
		ranking = OnePair
	default:
		ranking = HighCard
	}

	p.Ranking = ranking
	p.Hand = handStr
	p.Bid = bid
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hands_strings := make([]string, 0)	
	bids := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fields := strings.Fields(line)
		hands_strings = append(hands_strings, fields[0])
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Invalid input:", fields[1])
			os.Exit(1)
		}
		bids = append(bids, bid)
	}

	hands := make([]PokerHand, len(hands_strings))
	for i := 0; i < len(hands_strings); i++ {
		hands[i] = PokerHand{}
		hands[i].setHand(hands_strings[i], bids[i])
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compare(hands[j]) < 0
	})

	total_winnings := 0
	for rank, hand := range hands {
		total_winnings += (rank + 1) * hand.Bid
	}
	
	fmt.Println(total_winnings)
}

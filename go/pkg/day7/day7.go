package day7

import (
	"fmt"
	utils "rospierski/aocgo/pkg/aocutils"
    "strings"
    "strconv"
    "sort"
)

const (
    HIGH_CARD = iota
    ONE_PAIR = iota
    TWO_PAIR = iota
    THREE_KIND = iota
    FULL_HOUSE = iota
    FOUR_KIND = iota
    FIVE_KIND = iota
)

type Hand struct {
    cards int
    cards_str string
    kind int
    bid int
}

func (h *Hand) String() string {
    return fmt.Sprintf("%d %d %d", h.cards, h.kind, h.bid)
}

func newHand(line string) (Hand, error) {
    str_split := strings.Split(line, " ")
    cards_str, bid_str := str_split[0], str_split[1]
    card_bin := make(map[rune]int)
    for _, c := range cards_str {
        card_bin[c]++
    }

    og_cards_str := cards_str
    cards_str = strings.ReplaceAll(cards_str, "A", "E")
    cards_str = strings.ReplaceAll(cards_str, "K", "D")
    cards_str = strings.ReplaceAll(cards_str, "Q", "C")
    cards_str = strings.ReplaceAll(cards_str, "J", "B")
    cards_str = strings.ReplaceAll(cards_str, "T", "A")

    decimal_from_hex, err := strconv.ParseInt(cards_str, 16, 64)
    if err != nil {
        return Hand{}, err
    }
    cards := int(decimal_from_hex)
    bid_num, err := strconv.Atoi(bid_str)
    kind := -1
    switch len(card_bin) {
    case 5:
        kind = HIGH_CARD
    case 4:
        kind = ONE_PAIR
    case 3:
        for _, v := range card_bin {
            if v == 3 {
                kind = THREE_KIND
                break
            }
            if v == 2 {
                kind = TWO_PAIR
                break
            }
        }
    case 2:
        for _, v := range card_bin {
            if v == 4 {
                kind = FOUR_KIND
                break
            }
            if v == 3 {
                kind = FULL_HOUSE
                break
            }
        }
    case 1:
        kind = FIVE_KIND
    }
    return Hand{
        cards,
        og_cards_str,
        kind,
        bid_num,
    }, nil 
}

func getHands(lines []string) ([]Hand, error) {
    hands := make([]Hand, len(lines))
    for i, line := range lines {
        hand, err := newHand(line)
        if err != nil {
            return nil, err
        }
        hands[i] = hand
    }
    return hands, nil
}

func Day7Part1(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return -1, err
	}

    hands, err := getHands(lines)
    if err != nil {
        return -1, err
    }

    // sort hands
    sort.Slice(hands, func(i, j int) bool {
        if hands[i].kind == hands[j].kind {
            return hands[i].cards < hands[j].cards
        }
        return hands[i].kind < hands[j].kind
    })

    sum := 0
    for i, hand := range hands {
        sum += hand.bid * (i + 1)
    }

	return sum, nil
}

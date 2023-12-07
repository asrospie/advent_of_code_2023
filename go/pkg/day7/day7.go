package day7

import (
	"errors"
	"fmt"
	utils "rospierski/aocgo/pkg/aocutils"
	"sort"
	"strconv"
	"strings"
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

func hexHandToInt(hand string) (int, error) {
    decimal_from_hex, err := strconv.ParseInt(hand, 16, 64)
    if err != nil {
        return -1, err
    }
    return int(decimal_from_hex), nil
}

func getHandKind(card_bin map[rune]int) int {
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
    return kind
}

func newHand(line string, face_card_map map[string]string) (Hand, error) {
    str_split := strings.Split(line, " ")
    cards_str, bid_str := str_split[0], str_split[1]
    card_bin := make(map[rune]int)
    for _, c := range cards_str {
        card_bin[c]++
    }

    og_cards_str := cards_str
    cards_str = strings.ReplaceAll(cards_str, "A", face_card_map["A"])
    cards_str = strings.ReplaceAll(cards_str, "K", face_card_map["K"])
    cards_str = strings.ReplaceAll(cards_str, "Q", face_card_map["Q"])
    cards_str = strings.ReplaceAll(cards_str, "J", face_card_map["J"])
    cards_str = strings.ReplaceAll(cards_str, "T", face_card_map["T"])

    decimal_from_hex, err := hexHandToInt(cards_str)
    if err != nil {
        return Hand{}, err
    }
    cards := int(decimal_from_hex)
    bid_num, err := strconv.Atoi(bid_str)
    kind := getHandKind(card_bin)
    if kind == -1 {
        return Hand{}, errors.New("invalid hand")
    }
    return Hand{
        cards,
        og_cards_str,
        kind,
        bid_num,
    }, nil 
}

func getHands(lines []string) ([]Hand, error) {
    face_card_map := map[string]string{
        "A": "E",
        "K": "D",
        "Q": "C",
        "J": "B",
        "T": "A",
    }
    hands := make([]Hand, len(lines))
    for i, line := range lines {
        hand, err := newHand(line, face_card_map)
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

func getHandsJoker(lines []string) ([]Hand, error) {
    hands := make([]Hand, len(lines))
    for i, line := range lines {
        hand, err := newHandJoker(line)
        if err != nil {
            return nil, err
        }
        hands[i] = hand
    }
    return hands, nil
}

func newHandJoker(line string) (Hand, error) {
    face_card_map := map[string]string{
        "A": "E",
        "K": "D",
        "Q": "C",
        "J": "1",
        "T": "A",
    }
    hand, err := newHand(line, face_card_map) 
    if err != nil {
        return Hand{}, err
    }

    joker_count := strings.Count(hand.cards_str, "J")
    new_kind := hand.kind
    if joker_count > 0 {
        switch hand.kind {
        case HIGH_CARD:
            new_kind = ONE_PAIR
        case ONE_PAIR:
            new_kind = THREE_KIND
        case TWO_PAIR:
            if joker_count == 1 {
                new_kind = FULL_HOUSE
            } else {
                new_kind = FOUR_KIND
            }
        case THREE_KIND:
            new_kind = FOUR_KIND
        case FULL_HOUSE:
            new_kind = FIVE_KIND
        case FOUR_KIND:
            new_kind = FIVE_KIND
        default:
            new_kind = hand.kind
        }
    }
    if new_kind != hand.kind {
        fmt.Printf("%s %s %s\n", hand.cards_str, kindToString(hand.kind), kindToString(new_kind))
    }
    hand.kind = new_kind
    return hand, nil
}

func kindToString(kind int) string {
    switch kind {
    case HIGH_CARD:
        return "HIGH_CARD"
    case ONE_PAIR:
        return "ONE_PAIR"
    case TWO_PAIR:
        return "TWO_PAIR"
    case THREE_KIND:
        return "THREE_KIND"
    case FULL_HOUSE:
        return "FULL_HOUSE"
    case FOUR_KIND:
        return "FOUR_KIND"
    case FIVE_KIND:
        return "FIVE_KIND"
    }
    return "INVALID"
}

func Day7Part2(filename string) (int, error) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		return -1, err
	}

    hands, err := getHandsJoker(lines)
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

    for _, hand := range hands {
        fmt.Println(hand)
    }

	return sum, nil
}

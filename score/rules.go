package score 

import (
	"fmt"
	"sort"
)

type Rule int

const (
	Aces Rule = iota
	Twos
	Threes
	Fours
	Fives
	Sixes
	ThreeOfAKind
	FourOfAKind
	FullHouse
	SmallStraight
	LargeStraight
	Yahtzee
	Chance
	LastRule
)

type Scoring struct {
	Hand *Hand
	Rule Rule

	Score int
	Name  string
}

func New(rule Rule, hand *Hand) (*Scoring, error) {
	s := &Scoring{
		Rule: rule,
		Hand: hand,
	}
	switch rule {
	case 0:
		s.Aces()
		s.Name = "Aces"
	case 1:
		s.Twos()
		s.Name = "Twos"
	case 2:
		s.Threes()
		s.Name = "Threes"
	case 3:
		s.Fours()
		s.Name = "Fours"
	case 4:
		s.Fives()
		s.Name = "Fives"
	case 5:
		s.Sixes()
		s.Name = "Sixes"
	case 6:
		s.ThreeOfAKind()
		s.Name = "Three of a Kind"
	case 7:
		s.FourOfAKind()
		s.Name = "Four of a Kind"
	case 8:
		s.FullHouse()
		s.Name = "Full House"
	case 9:
		s.SmallStraight()
		s.Name = "Small Straight"
	case 10:
		s.LargeStraight()
		s.Name = "Large Straight"
	case 11:
		s.Yahtzee()
		s.Name = "Yahtzee"
	case 12:
		s.Chance()
		s.Name = "Chance"
	default:
		return nil, fmt.Errorf("invalid scoring")
	}
	return s, nil
}

func (s *Scoring) String() string {
	return s.Name
}

func (s *Scoring) Aces() {
	s.UpperSectionScore(1)
}

func (s *Scoring) Twos() {
	s.UpperSectionScore(2)
}

func (s *Scoring) Threes() {
	s.UpperSectionScore(3)
}

func (s *Scoring) Fours() {
	s.UpperSectionScore(4)
}

func (s *Scoring) Fives() {
	s.UpperSectionScore(5)
}

func (s *Scoring) Sixes() {
	s.UpperSectionScore(6)
}

func (s *Scoring) ThreeOfAKind() {
	s.OfAKind(3)
}

func (s *Scoring) FourOfAKind() {
	s.OfAKind(4)
}

func (s *Scoring) FullHouse() {
	m := make(map[int]int)
	for i := 1; i <= 6; i++ {
		m[i]++
	}
	two := false
	three := false
	for _, v := range m {
		if v == 2 {
			two = true
		}
		if v == 3 {
			three = true
		}
	}
	if two && three {
		s.Score = 25
	}
}

func (s *Scoring) SmallStraight() {
	l := s.Hand.GetHand()
	sort.Ints(l)
	if s.IsStraight(l[:4]) || s.IsStraight(l[1:]) {
		s.Score = 30
	}
}

func (s *Scoring) LargeStraight() {
	l := s.Hand.GetHand()
	sort.Ints(l)
	if s.IsStraight(l) {
		s.Score = 40
	}
}

func (s *Scoring) Yahtzee() {
	s.OfAKind(5)
}

func (s *Scoring) Chance() {
	s.Score = s.Hand.Sum()
}

// Calculation helpers of Upper and Lower Section scorings.
func (s *Scoring) UpperSectionScore(v int) {
	s.Score = s.Hand.Count(v) * v
}

func (s *Scoring) OfAKind(v int) {
	for i := 1; i <= 6; i++ {
		if s.Hand.Count(i) >= v {
			if v != 5 {
				s.Score = s.Hand.Sum()
			} else {
				s.Score = 50
			}
		}
	}
}

func (s *Scoring) IsStraight(v []int) bool {
	rep := make(map[int]bool)
	sum := 0
	for _, f := range v {
		if _, value := rep[f]; !value {
			rep[f] = true
			sum += f
		} else {
			return false
		}
	}
	consecutiveSum := len(v) / 2 * (v[0] + v[len(v)-1])

	return sum == consecutiveSum
}

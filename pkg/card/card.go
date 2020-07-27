package card

import (
	"strconv"
	"strings"
)

type Service struct {
	BankName   string
	Cards      []*Card
	BankPrefix string
}

func (s *Service) SearchByNumber(number string) (*Card, bool) {
	if !strings.HasPrefix(number, s.BankPrefix) {
		return nil, false
	}
	for _, card := range s.Cards {
		if card.Number == number {
			return card, true
		}
	}
	return nil, false
}

func (s *Service) AddCards(cards ...*Card) {
	s.Cards = append(s.Cards, cards...)
}

func IsValid(number string) bool {
	sum := 0
	number = strings.ReplaceAll(number, " ", "")
	numbers := strings.Split(number, "")
	digits := make([]int, len(numbers))

	for i, v := range numbers {
		d, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		digits[i] = d
	}

	for i, num := range digits {
		if (len(digits)-i)%2 == 0 {
			num = num * 2

			if num > 9 {
				num = num - 9
			}
		}
		sum += num
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

func NewService(bankName string, bankPrefix string) *Service {
	return &Service{BankName: bankName}
}

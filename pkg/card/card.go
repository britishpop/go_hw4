package card

import "strings"

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

func (s *Service) AddCards(cards []*Card) {
	for _, card := range cards {
		s.Cards = append(s.Cards, card)
	}
}

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

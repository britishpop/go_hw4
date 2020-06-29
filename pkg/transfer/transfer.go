package transfer

import "go_hw4/pkg/card"

type Service struct {
	CardSvc           *card.Service
	PercentCommission float64
	MinSumCommission  int64
}

func NewService(cardSvc *card.Service, percent float64, minSum int64) *Service {
	return &Service{
		CardSvc:           cardSvc,
		PercentCommission: percent,
		MinSumCommission:  minSum,
	}
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {

	fromCard := s.CardSvc.SearchByNumber(from)
	toCard := s.CardSvc.SearchByNumber(to)

	commission := s.PercentCommission / 100
	transferSum := int64(float64(amount) + commission*float64(amount))
	total = transferSum

	if transferSum < s.MinSumCommission {
		transferSum = s.MinSumCommission
	}

	if fromCard == nil && toCard == nil {
		ok = true
		return
	}

	if fromCard != nil && toCard != nil {
		total = amount

		if amount < fromCard.Balance {
			toCard.Balance = toCard.Balance + amount
			fromCard.Balance = fromCard.Balance - amount
			ok = true
		}
	}

	if fromCard != nil && transferSum >= fromCard.Balance {
		ok = false
		return
	}

	if fromCard != nil {
		fromCard.Balance = fromCard.Balance - transferSum
		ok = true
	}

	if toCard != nil {
		toCard.Balance = toCard.Balance + total
		ok = true
	}
	return
}

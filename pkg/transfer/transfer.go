package transfer

import (
	"errors"
	"go_hw4/pkg/card"
)

var ErrSourceCardBalance = errors.New("card balance is not enough to transfer")

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

func (s *Service) Card2Card(from, to string, amount int64) (total int64, err error) {

	fromCard := s.CardSvc.SearchByNumber(from)
	toCard := s.CardSvc.SearchByNumber(to)

	commission := s.PercentCommission / 100
	transferSum := float64(amount) + commission*float64(amount)
	total = int64(transferSum)

	if transferSum < float64(s.MinSumCommission) {
		transferSum = float64(s.MinSumCommission)
	}

	if fromCard == nil && toCard == nil {
		err = nil
		return
	}

	if fromCard != nil && toCard != nil {
		total = amount

		if float64(amount) >= float64(fromCard.Balance) {
			err = ErrSourceCardBalance
		} else {
			toCard.Balance = toCard.Balance + amount
			fromCard.Balance = fromCard.Balance - amount
			err = nil
		}
		return
	}

	if fromCard != nil && transferSum >= float64(fromCard.Balance) {
		err = ErrSourceCardBalance
		return
	}

	if fromCard != nil {
		fromCard.Balance = fromCard.Balance - int64(transferSum)
		err = nil
	}

	if toCard != nil {
		toCard.Balance = toCard.Balance + total
		err = nil
	}
	return
}

package transfer

import (
	"errors"
	"go_hw4/pkg/card"
)

var (
	ErrSourceCardBalanceNotEnough = errors.New("card balance is not enough to transfer")
	ErrSourceCardNotFound         = errors.New("source card not found")
	ErrTargetCardNotFound         = errors.New("target card not found")
)

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

	fromCard, okFrom := s.CardSvc.SearchByNumber(from)
	toCard, okTo := s.CardSvc.SearchByNumber(to)

	if !okFrom {
		return 0, ErrSourceCardNotFound
	}

	if !okTo {
		return 0, ErrTargetCardNotFound
	}

	commission := s.PercentCommission / 100
	transferSum := int64(float64(amount) + commission*float64(amount))
	total = transferSum

	if transferSum < s.MinSumCommission {
		transferSum = s.MinSumCommission
	}

	if amount < fromCard.Balance {
		toCard.Balance = toCard.Balance + amount
		fromCard.Balance = fromCard.Balance - amount
		err = nil
	}

	if fromCard != nil && transferSum >= fromCard.Balance {
		err = ErrSourceCardBalanceNotEnough
		return
	}

	if fromCard != nil {
		fromCard.Balance = fromCard.Balance - transferSum
		err = nil
	}

	if toCard != nil {
		toCard.Balance = toCard.Balance + total
		err = nil
	}
	return
}

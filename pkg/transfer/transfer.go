package transfer

import (
	"errors"
	"go_hw4/pkg/card"
)

var (
	ErrSourceCardBalanceNotEnough = errors.New("card balance is not enough to transfer")
	ErrSourceCardNotFound         = errors.New("source card not found")
	ErrTargetCardNotFound         = errors.New("target card not found")
	ErrInvalidSourceCardNumber    = errors.New("invalid source card number")
	ErrInvalidTargetCardNumber    = errors.New("invalid target card number")
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

func Transfer(from, to string) error {
	if !card.IsValid(from) {
		return ErrInvalidSourceCardNumber
	}
	if !card.IsValid(to) {
		return ErrInvalidTargetCardNumber
	}
	return nil
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, err error) {

	if err := Transfer(from, to); err != nil {
		return 0, err
	}

	fromCard, okFrom := s.CardSvc.SearchByNumber(from)
	toCard, okTo := s.CardSvc.SearchByNumber(to)

	commission := s.PercentCommission / 100
	transferSum := int64(float64(amount) + commission*float64(amount))
	total = amount

	if transferSum < s.MinSumCommission {
		transferSum = s.MinSumCommission
	}

	if !okTo {
		return transferSum, ErrTargetCardNotFound
	}

	if !okFrom {
		return transferSum, ErrSourceCardNotFound
	}

	if amount < fromCard.Balance {
		toCard.Balance = toCard.Balance + amount
		fromCard.Balance = fromCard.Balance - amount
		err = nil
	} else {
		err = ErrSourceCardBalanceNotEnough
	}
	return
}

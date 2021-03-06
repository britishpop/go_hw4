package transfer

import (
	"go_hw4/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc           *card.Service
		PercentCommission float64
		MinSumCommission  int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}

	cardSvc := card.NewService("Tinkoff", "510621")

	cardSvc.AddCards(&card.Card{Balance: 100_00, Number: "5106217896580000"},
		&card.Card{Balance: 64700_00, Number: "5106217896540004"},
		&card.Card{Balance: 3950_00, Number: "5106218888990009"})

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantErr   error
	}{
		{
			name:      "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "5106218888990009", to: "5106217896540004", amount: 3000_00},
			wantTotal: 3000_00,
			wantErr:   nil,
		},
		{
			name:      "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "5106217896580000", to: "5106217896540004", amount: 600_00},
			wantTotal: 600_00,
			wantErr:   ErrSourceCardBalanceNotEnough,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "5106217896540004", to: "48969999654387", amount: 100_00},
			wantTotal: 110_00,
			wantErr:   ErrTargetCardNotFound,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "5106217896580000", to: "987673923923982", amount: 400_00},
			wantTotal: 440_00,
			wantErr:   ErrTargetCardNotFound,
		},
		{
			name:      "Карта чужого банка -> Карта своего банка",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "3854912828238235", to: "5106218888990009", amount: 50_00},
			wantTotal: 55_00,
			wantErr:   ErrSourceCardNotFound,
		},
		{
			name:      "Карта чужого банка -> Карта чужого банка",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "753292992392094", to: "3854912828238235", amount: 500_00},
			wantTotal: 550_00,
			wantErr:   ErrTargetCardNotFound,
		},
	}
	for _, tt := range tests {
		s := &Service{
			CardSvc:           tt.fields.CardSvc,
			PercentCommission: tt.fields.PercentCommission,
			MinSumCommission:  tt.fields.MinSumCommission,
		}
		gotTotal, gotErr := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
		if gotTotal != tt.wantTotal {
			t.Errorf("Service.Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
		}
		if gotErr != tt.wantErr {
			t.Errorf("Service.Card2Card() gotErr = %v, want %v", gotErr, tt.wantErr)
		}
	}
}

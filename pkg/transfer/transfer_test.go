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

	cardSvc := card.NewService("VTB")
	cards := []*card.Card{&card.Card{Balance: 100_00, Number: "0003"},
		&card.Card{Balance: 64700_00, Number: "0004"},
		&card.Card{Balance: 4000_00, Number: "0005"},
		&card.Card{Balance: 3950_00, Number: "0006"}}

	cardSvc.AddCards(cards)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		{
			name:      "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "0004", to: "0005", amount: 3000_00},
			wantTotal: 3000_00,
			wantOk:    true,
		},
		{
			name:      "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "0003", to: "0005", amount: 400_00},
			wantTotal: 400_00,
			wantOk:    false,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "0006", to: "4896", amount: 100_00},
			wantTotal: 110_00,
			wantOk:    true,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "0003", to: "9876", amount: 400_00},
			wantTotal: 440_00,
			wantOk:    false,
		},
		{
			name:      "Карта чужого банка -> Карта своего банка",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "3854", to: "0003", amount: 50_00},
			wantTotal: 55_00,
			wantOk:    true,
		},
		{
			name:      "Карта чужого банка -> Карта чужого банка",
			fields:    fields{CardSvc: cardSvc, PercentCommission: 10, MinSumCommission: 10_00},
			args:      args{from: "7532", to: "9678", amount: 500_00},
			wantTotal: 550_00,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		s := &Service{
			CardSvc:           tt.fields.CardSvc,
			PercentCommission: tt.fields.PercentCommission,
			MinSumCommission:  tt.fields.MinSumCommission,
		}
		gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
		if gotTotal != tt.wantTotal {
			t.Errorf("Service.Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
		}
		if gotOk != tt.wantOk {
			t.Errorf("Service.Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
		}
	}
}

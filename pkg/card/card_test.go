package card

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Card number is valid",
			args: args{number: "5106218888990009"},
			want: true,
		},
		{
			name: "Card number is invalid",
			args: args{number: "5106218888997609"},
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsValid(tt.args.number)
		if got != tt.want {
			t.Errorf("IsValid() = %v, want %v", got, tt.want)
		}
	}
}

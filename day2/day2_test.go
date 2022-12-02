package day2

import (
	"testing"
)

func Test_winner(t *testing.T) {
	type args struct {
		o string
		u string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"A beats Z", args{o: "A", u: "Z"}, 1, false},
		{"B beats X", args{o: "B", u: "X"}, 1, false},
		{"C beats Y", args{o: "C", u: "Y"}, 1, false},
		{"X beats C", args{o: "C", u: "X"}, 2, false},
		{"Y beats A", args{o: "A", u: "Y"}, 2, false},
		{"Z beats B", args{o: "B", u: "Z"}, 2, false},
		{"A draws X", args{o: "A", u: "X"}, 0, false},
		{"B draws Y", args{o: "B", u: "Y"}, 0, false},
		{"C draws Z", args{o: "C", u: "Z"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := winner(tt.args.o, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("winner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("winner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOpponent(t *testing.T) {
	type args struct {
		o string
	}
	tests := []struct {
		name    string
		args    args
		want    Hand
		wantErr bool
	}{
		{"A is R", args{o: "A"}, Rock, false},
		{"B is P", args{o: "B"}, Paper, false},
		{"C is S", args{o: "C"}, Scissors, false},
		{"T is U", args{o: "T"}, Unknown, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOpponent(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOpponent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseOpponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseYou(t *testing.T) {
	type args struct {
		o string
	}
	tests := []struct {
		name    string
		args    args
		want    Hand
		wantErr bool
	}{
		{"X is R", args{o: "X"}, Rock, false},
		{"Y is P", args{o: "Y"}, Paper, false},
		{"Z is S", args{o: "Z"}, Scissors, false},
		{"T is U", args{o: "T"}, Unknown, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseYou(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseYou() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseYou() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_score(t *testing.T) {
	type args struct {
		u      Hand
		result int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "R and Lose", args{Rock, 1}, 1 },
		{ "P and Lose", args{Paper, 1}, 2 },
		{ "S and Lose", args{Scissors, 1}, 3 },
		{ "R and Draw", args{Rock, 0}, 4 },
		{ "P and Draw", args{Paper, 0}, 5 },
		{ "S and Draw", args{Scissors, 0}, 6 },
		{ "R and Win", args{Rock, 2}, 7 },
		{ "P and Win", args{Paper, 2}, 8 },
		{ "S and Win", args{Scissors, 2}, 9 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.args.u, tt.args.result); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}

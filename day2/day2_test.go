package day2

import (
	"testing"
)

func Test_winner(t *testing.T) {
	type args struct {
		o Hand
		u Hand
	}
	tests := []struct {
		name string
		args args
		want Result
	}{
		{"Rock beats Scissors", args{o: Rock, u: Scissors}, YouLose},
		{"Paper beats Rock", args{o: Paper, u: Rock}, YouLose},
		{"Scissors beats Paper", args{o: Scissors, u: Paper}, YouLose},
		{"Rock beats Scissors", args{o: Scissors, u: Rock}, YouWin},
		{"Paper beats Rock", args{o: Rock, u: Paper}, YouWin},
		{"Scissors beats Paper", args{o: Paper, u: Scissors}, YouWin},
		{"Rock draws Rock", args{o: Rock, u: Rock}, YouDraw},
		{"Paper draws Paper", args{o: Paper, u: Paper}, YouDraw},
		{"Scissors draws Scissors", args{o: Scissors, u: Scissors}, YouDraw},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := winner(tt.args.o, tt.args.u)
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
		result Result
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"R and Lose", args{Rock, YouLose}, 1},
		{"P and Lose", args{Paper, YouLose}, 2},
		{"S and Lose", args{Scissors, YouLose}, 3},
		{"R and Draw", args{Rock, YouDraw}, 4},
		{"P and Draw", args{Paper, YouDraw}, 5},
		{"S and Draw", args{Scissors, YouDraw}, 6},
		{"R and Win", args{Rock, YouWin}, 7},
		{"P and Win", args{Paper, YouWin}, 8},
		{"S and Win", args{Scissors, YouWin}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.args.u, tt.args.result); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}

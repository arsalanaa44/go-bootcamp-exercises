package entity

import "time"

type Game struct {
	ID          int
	CategoryID  int
	QuestionIDs []int
	//Players   []Player
	//homogenization
	PlayerIDs []int
	StartTime time.Time
}
type Player struct {
	ID      int
	UserID  int
	GameID  int
	Score   int
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         int
	PlayerID   int
	QuestionID int
	choice     PossibleAnswerChoice
}

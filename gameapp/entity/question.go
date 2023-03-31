package entity

type Question struct {
	ID              int
	Text            string
	PossibleAnswers []PossibleAnswer
	CorrectAnswerID int
	Difficulty      QuestionDifficulty
	CategoryID      int
}

type PossibleAnswer struct {
	Text   string
	Choice PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	return p >= PossibleAnswerA && p <= PossibleAnswerD
}

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (qd QuestionDifficulty) IsValid() bool {
	return qd >= QuestionDifficultyEasy && qd <= QuestionDifficultyHard
}

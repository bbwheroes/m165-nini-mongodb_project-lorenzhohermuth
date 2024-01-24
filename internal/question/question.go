package question

type Question struct {
	questionText string
	whatValue string
}

var values = [4]string{"weight", "base_experience", "height", "is_base_form"}

func GenerateQuestion() Question {
	return Question {
		questionText: "?",
		whatValue: values[0],
	}
}

func (q Question) GetWhatValue() string {
	return q.whatValue
}

func (q Question) GetQuestionText() string {
	return q.questionText
}

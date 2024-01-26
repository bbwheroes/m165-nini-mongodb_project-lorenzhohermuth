package question

import "math/rand"

type Question struct {
	questionText string
	whatValue string
}

var values = [4]string{"weight", "base_experience", "height", "is_base_form"}

func GenerateQuestion() Question {
	index := rand.Intn(len(values))
	return Question {
		questionText: getQuestionText(index),
		whatValue: values[index],
	}
}

func (q Question) GetWhatValue() string {
	return q.whatValue
}

func (q Question) GetQuestionText() string {
	return q.questionText
}

func getQuestionText(i int) string {
	if i == 0 {
		return "With Listed Pokemon weighs %v kg ?"
	}
	if i == 1 {
		return "With Listed Pokemon has a Base Experience of %v exp ?"
	}
	if i == 2 {
		return "With Listed Pokemon is %v cm tall ?"
	}
	if i == 3 {
		return "This Pokemon %v in its Base Form ?"
	}
	return ""
}

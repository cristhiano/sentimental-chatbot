package nlp

import "github.com/cdipaolo/sentiment"

func MachineLearningScore(text string) int {
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	analysis := model.SentimentAnalysis(text, sentiment.English)

	return int(analysis.Score)
}

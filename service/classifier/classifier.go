//go:generate mockgen -package classifier -source=classifier.go -destination classifier_mocks.go

package classifier


type Classifier struct {
}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Classify(input string) string  {
	//list of classification functions
	//iterate through applying each to input
	//break when function returns true and return classification
	return "filler"
}

func iAmFeeling(input string) string {
	//rule for spotting "I am feeling [adjective]" sentence
	//sentence that starts "I am feeling", contains an adjective, and the next character is not a space
}
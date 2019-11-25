package classifier


type Classifier struct {
}

func NewClassifier() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Classify(input string) string  {
	return "filler"
}


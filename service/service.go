package service


type Classifier interface {
	Classify(input string) string
}

type Reponder interface {
	Respond(input string, classification string) string
}

type Service struct {
	classifier Classifier
	responder Reponder
}

func NewService(classifier Classifier, responder Reponder) *Service{
	return &Service{classifier:classifier, responder:responder}
}

func (s *Service) ListenAndRespond(input string) string {
	class := s.classifier.Classify(input)
	response := s.responder.Respond(input, class)
	return response
}


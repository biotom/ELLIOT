package responder

import "math/rand"

const (
	default1 = "Can you elaborate on that?"
	default2 = "Does that trouble you?"
	default3 = "What are your feelings now?"
	default4 = "What do you mean by that?"
	default5 = "How does that make you feel?"

)
func fillerResponder() func() string {
	i := 0
	responses := []string{default1, default2, default3, default4, default5}
	responseIndeces := rand.Perm(len(responses))
	if i > len(responses){
		i = 0
	}
	return func() string {
		responseIndex := responseIndeces[i]
		i += 1
		return responses[responseIndex]
	}
}



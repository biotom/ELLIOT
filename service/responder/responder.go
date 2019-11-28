//go:generate mockgen -package responder -source=responder.go -destination responder_mocks.go

package responder

type Responder struct {}

func NewResponder() *Responder {
	return &Responder{}
}

func (r *Responder) Respond(input string, classification string) string {
	response := fillerResponder()
	return response()
}
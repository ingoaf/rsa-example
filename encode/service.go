package encode

// Service represents a rune <-> string encoding technique
type Service interface {
	EncodeMessage(string) []rune
	DecodeMessage([]rune) string
}

type service struct {
}

// NewService generates a new incoding instance
func NewService() Service {
	return &service{}
}

func (s *service) EncodeMessage(m string) []rune {
	return []rune(m)
}

func (s *service) DecodeMessage(r []rune) string {
	return string(r)
}

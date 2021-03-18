package creational_patterns

// We need some single, shared value of some particular type.
// We need to restrict object creation of some type to a single unit along the entire program.
// this implementation is not thread safe!
type Singleton interface {
	AddoOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

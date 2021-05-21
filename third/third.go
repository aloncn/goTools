package gotool

var Third = third{}
type third struct {}

func (t *third)SayHi() string {
	return "hi"
}

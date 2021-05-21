package T

var T = third{}
type third struct {}

func (t *third)SayHi() string {
	return "hi"
}

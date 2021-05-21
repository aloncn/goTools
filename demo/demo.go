package gotool

var Demo = demo{}
type demo struct {}

func (d *demo) SayHi (s string) string {
	return "hello, " + s
}

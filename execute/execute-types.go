package execute

type Result struct {
	output []byte
}

type ConvertResult interface {
	String() string
}

func (r Result) String() string {
	return string(r.output[:])
}

package tracker

type Input interface {
	Get() string
}

type Output interface {
	Out(text string)
}

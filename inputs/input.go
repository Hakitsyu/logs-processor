package inputs

type Input interface {
	Load() [][]byte
	name() string
}

type InputConfig struct {
	name string
}

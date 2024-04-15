package transport

type Transport interface {
	Input(string) string
}

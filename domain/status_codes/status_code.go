package status_codes

type StatusCode interface {
	String() string
	Int() int
}

package typeconv

type FormatProvider interface {
	CanFormat(v interface{}) bool
	Format(v interface{}) interface{}
}

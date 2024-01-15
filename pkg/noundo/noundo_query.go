package noundo

type OrderIface[T any] interface {
	Less(t1 T, t2 T) bool
}

type FilterIface[T any] interface {
	Keep(t T) bool
}

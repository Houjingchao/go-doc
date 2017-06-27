package singleton

type cat struct {
	name string
}

func New() *cat {
	return new(cat)
}

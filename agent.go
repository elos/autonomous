package autonomous

type Core struct {
	Life
	Managed
	Stopper
}

func NewCore() *Core {
	return &Core{
		Life:    NewLife(),
		Stopper: make(Stopper),
	}
}

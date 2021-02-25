package hash

//Driver hash加密驱动接口
type Driver interface {
	Make(password []byte) ([]byte, error)
	Check(hashedPassword, password []byte) error
}

type hash struct {
	driver Driver
}

func (h *hash)Make(password []byte) ([]byte, error) {
	return h.driver.Make(password)
}

func (h *hash)Check(hashedPassword, password []byte) error {
	return h.driver.Check(hashedPassword, password)
}

func NewHash(handle... OptionHandle) *hash {
	options := Options{
		Type: BCRYPT,
		Cost: DEFAULTCOST,
	}
	for _, o := range handle{
		o(&options)
	}
	driver := getDriver(options)
	return &hash{
		driver: driver,
	}
}

func getDriver(options Options) Driver {
	var driver Driver

	switch options.Type {
	case BCRYPT:
		driver = &Bcrypt{cost: options.Cost}
		break
	default:
		panic("hash driver not found")
	}

	return driver
}


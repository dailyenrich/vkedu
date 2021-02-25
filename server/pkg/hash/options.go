package hash

import "golang.org/x/crypto/bcrypt"

const (
	BCRYPT = 1
	DEFAULTCOST = bcrypt.DefaultCost
)

type OptionHandle func(options *Options)

type Options struct {
	Type int
	Cost int
}

func WithDriver(driver int) OptionHandle {
	return func(options *Options) {
		options.Type = driver
	}
}

func WithConst(cost int) OptionHandle {
	return func(options *Options) {
		options.Cost = cost
	}
}

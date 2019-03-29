package ifmock

type MyInterface interface {
	NoReturnAndParamater()
	NoReturnAndOneParamater(int)
	NoReturnAndTwoParamater(int, string)
	OneReturnAndOneParamater(int) error
	OneReturnAndTwoParamater(int, string) error
	TwoReturnAndOneParamater(int) (string, error)
	TwoReturnAndTwoParamater(int, string) (int, error)
}

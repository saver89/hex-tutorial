package ports

type ArithmeticPort interface {
	Addition(a, b int32) (int32, error)
	Substruction(a, b int32) (int32, error)
	Mutliplication(a, b int32) (int32, error)
	Division(a, b int32) (int32, error)
}

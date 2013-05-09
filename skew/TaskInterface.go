package skew

type TaskInterface interface {
	Name() string
	Run(*Logger, interface{}) string
	CanReschedule(error) bool
}

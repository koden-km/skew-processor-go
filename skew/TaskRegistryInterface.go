package skew

type TaskRegistryInterface interface {
	Get(name string) *TaskInterface
	Capabilities() []string
}

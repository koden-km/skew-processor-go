package skewd

// TODO:
// This naming is not Go style?

type TaskRegistryInterface interface {
	Get(name string) *TaskInterface
	Capabilities() []string
}

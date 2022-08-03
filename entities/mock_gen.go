package entities

//go:generate mockery --name=TaskHandler --inpackage --case=underscore --testonly

// TaskHandler ...
//
// It's used only for mock generating.
//
type TaskHandler interface {
	HandleTask(task Task)
}

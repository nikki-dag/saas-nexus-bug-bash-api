package service

const (
	BugBashServiceName = "nexus-bug-bash-service"

	EchoOperationName  = "nexus-bug-bash-echo-operation"
	HelloOperationName = "nexus-bug-bash-hello-operation"
)

type (
	EchoInput struct {
		Message string
	}
	EchoOutput EchoInput

	HelloInput struct {
		Name     string
		Language string
	}
	HelloOutput struct {
		Message string
	}
)

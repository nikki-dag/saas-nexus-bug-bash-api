package service

const (
	BugBashServiceName = "nexus-bug-bash-service"

	EchoOperationName  = "nexus-bug-bash-echo-operation"
	HelloOperationName = "nexus-bug-bash-hello-operation"

	EN Language = "en"
	FR Language = "fr"
	DE Language = "de"
	ES Language = "es"
	TR Language = "tr"
)

type (
	EchoInput struct {
		Message string
	}
	EchoOutput EchoInput

	Language   string
	HelloInput struct {
		Name     string
		Language Language
	}
	HelloOutput struct {
		Message string
	}
)

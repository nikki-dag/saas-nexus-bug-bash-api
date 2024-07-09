package service

const (
	EchoServiceName   = "nexus-echo-service"
	EchoOperationName = "nexus-echo-operation"
	EchoEndpointName  = "nexus_echo_endpoint"

	HelloServiceName   = "nexus-hello-service"
	HelloOperationName = "nexus-hello-operation"
	HelloEndpointName  = "nexus_hello_endpoint"
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

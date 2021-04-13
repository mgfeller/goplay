package test_files

import "fmt"

const (
	ErrGetNameCode            = "1000"
	ErrCreateInstanceCode     = "1001"
	ErrMeshConfigCode         = "1002"
	ErrValidateKubeconfigCode = "1003"
	ErrClientConfigCode       = "1004"
	ErrClientSetCode          = "1005"
	ErrStreamEventCode        = "1006"
	ErrOpInvalidCode          = "1007"
	ErrApplyOperationCode     = "1008"
	ErrListOperationsCode     = "1009"
	ErrNewSmiCode             = "1010"
	ErrRunSmiCode             = "1011"
	ErrNoResponseCode         = "1011"
)

var (

	// GRPC server specific codes
	// Range 600-699
	ErrPanic        = "600"
	ErrGrpcListener = "601"
	ErrGrpcServer   = "602"

	// Config specific codes
	// Range 700-799
	ErrEmptyConfig = "700"
	ErrInMem       = "701"
	ErrViper       = "702"

	// Tracing specific codes
	// Range 800-899

	// Meshery server specific codes
	// Range 10000-10099

	// Meshery adapter specific codes
	// Range 10100 to 10199
	ErrGetName        = "1000"
	ErrInstallMesh    = "1001"
	ErrMeshConfig     = "1002"
	ErrPortForward    = "1003"
	ErrClientConfig   = "1004"
	ErrClientSet      = "1005"
	ErrStreamEvent    = "1006"
	ErrOpInvalid      = "1007"
	ErrApplyOperation = "1008"
	ErrListOperations = "1009"

	// Meshkit specific codes
	// Range 10200 to 10299
	ErrSmiInit          = "kit_10200"
	ErrInstallSmi       = "kit_10201"
	ErrConnectSmi       = "kit_10202"
	ErrRunSmi           = "kit_10203"
	ErrDeleteSmi        = "kit_10204"
	ErrUnmarshal        = "kit_10205"
	ErrMarshal          = "kit_10205"
	ErrGetBool          = "kit_10205"
	ErrApplyManifest    = "kit_10206"
	ErrServiceDiscovery = "kit_10207"
	ErrLoadFile         = "kit_10208"
	ErrApplyHelmChart   = "kit_10209"

	// Istio Service mesh specific codes
	// Range 11000 to 11099

	// Linkerd Service mesh specific codes
	// Range 11100 to 11199

	// Open Service mesh specific codes
	// Range 11200 to 11299

	// Kuma Service mesh specific codes
	// Range 11300 to 11399

	// Citrix Service mesh specific codes
	// Range 11400 to 11499

	// Network Service mesh specific codes
	// Range 11500 to 11599

	// ErrInvalidReplicaSetNoSelectors is the error for replicaset (v1) with no selectors
	ErrInvalidReplicaSetNoSelectors = fmt.Errorf("invalid replicaset: no selectors, therefore cannot be exposed")

	// ErrNoPortsFoundForHeadlessResource is the error when no ports are found for non headless resource
	ErrNoPortsFoundForHeadlessResource = fmt.Errorf("no ports found for the non headless resource")

	// Consul Service mesh specific codes
	// Range 11600 to 11699

	// Octarine Service mesh specific codes
	// Range 11700 to 11799

	// Nginx Service mesh specific codes
	// Range 11800 to 11899

)

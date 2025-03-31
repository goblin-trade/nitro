// To activate live tracer with flags --node.vmtrace.tracername and --node.vmtrace.jsonconfig
package arbostypes

type VMTraceConfig struct {
	TracerName string `koanf:"tracername"`
	JSONConfig string `koanf:"jsonconfig"`
}

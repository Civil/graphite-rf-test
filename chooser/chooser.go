package chooser

// Result represents result of Simulation
type Result struct {
	MetricsTotal int
	MetricsLost  int
}

// Chooser maps keys to shards
type Chooser interface {
	// LoadMetrics loads metrics to the memory
	LoadMetrics([]string) error
	// KillServer - kills random 'number' servers
	Simulate(int) Result
}

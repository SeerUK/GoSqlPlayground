package database

// Version represents a database version, and provides the SQL to go "up" to this version. Down
// migrations are not supported, as it's generally better to have a sensible backup / recover
// strategy. Down migrations often become stale.
type Version interface {
	// Name returns the name of the version. @todo: Can we make this automatic?
	Name() string
	// Returns SQL to run for the "up" operation.
	Migration() string
	// Register this migration as available to run.
	Register()
}

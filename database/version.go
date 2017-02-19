package database

// Version represents a database version, and provides the SQL to go "up" to this version. Down
// migrations are not supported, as it's generally better to have a sensible backup / recover
// strategy. Down migrations often become stale.
type Version interface {
	// Migration returns SQL to run for the "up" operation.
	Migration() string
	// Number returns the version number of this migration. This is used to ensure the versions are
	// applied in the correct order.
	Number() int
}

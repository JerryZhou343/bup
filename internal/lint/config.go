package lint

// LintConfig is the lint config.
type LintConfig struct {
	// Group is the specific group of linters to use.
	// The default group is the "default" lint group, which is equal
	// to the "uber1" lint group.
	// Setting this value will result in NoDefault being ignored.
	Group string
	// NoDefault is set to exclude the default set of linters.
	// This value is ignored if Group is set.
	// Deprecated: Use group "empty" instead.
	NoDefault bool
	// IncludeIDs are the list of linter IDs to use in addition to the defaults.
	// Expected to be all uppercase.
	// Expected to be unique.
	// Expected to have no overlap with ExcludeIDs.
	IncludeIDs []string
	// ExcludeIDs are the list of linter IDs to exclude from the defaults.
	// Expected to be all uppercase.
	// Expected to be unique.
	// Expected to have no overlap with IncludeIDs.
	ExcludeIDs []string
	// IgnoreIDToFilePaths is the map of ID to absolute file path to ignore.
	// IDs expected to be all upper-case.
	// File paths expected to be absolute paths.
	IgnoreIDToFilePaths map[string][]string
	// FileHeader is contents of the file that contains the header for all
	// Protobuf files, typically a license header. If this is set and the
	// FILE_HEADER linter is turned on, files will be checked to begin
	// with the contents of this file, and format --fix will place this
	// header before the syntax declaration. Note that format --fix will delete
	// anything before the syntax declaration if this is set.
	FileHeader string
	// JavaPackagePrefix is the prefix for java packages. This only has an
	// effect if the linter FILE_OPTIONS_EQUAL_JAVA_PACKAGE_PREFIX is turned on.
	// This also affects create and format --fix.
	// The default behavior is to use "com".
	JavaPackagePrefix string
	// AllowSuppression says to honor @suppresswarnings annotations.
	AllowSuppression bool
}

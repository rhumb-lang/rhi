package loader

// FindShelfEntry returns the path to the +file.rh if it exists.
// Returns "", nil if no entry point (pure library).
// Returns error if multiple +files exist.
func FindShelfEntry(dir string) (string, []string, error) {
	// 1. Glob for all .rh files
	// 2. Identify the one starting with "+"
	// 3. Return (entryPointPath, listOfAllSourceFiles, err)
}

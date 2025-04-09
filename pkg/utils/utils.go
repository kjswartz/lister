// Package utils is a collection of utility functions and typings that are used in multiple places in the codebase.
package utils

import (
	"os"
	"path/filepath"
)

// DBPath is the path to the SQLite database file.
var DBPath = filepath.Join(os.Getenv("HOME"), "go", "data", "lister.db")

// +build !windows

package filesystem

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func markHidden(path string) error {
	// POSIX platforms don't have the notion of a hidden attribute, they only
	// hide dot-prefixed paths, so ensure that the path begins with a dot.
	if strings.IndexRune(filepath.Base(path), '.') != 0 {
		return errors.New("only dot-prefixed files are hidden on POSIX")
	}

	// Success.
	return nil
}

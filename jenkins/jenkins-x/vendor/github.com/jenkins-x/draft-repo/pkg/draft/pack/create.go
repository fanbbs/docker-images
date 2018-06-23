package pack

import (
	"fmt"
)

// CreateFrom scaffolds a directory with the src pack.
func CreateFrom(dest, src string) error {
	// first do some validation that we are copying from a valid pack directory
	pack, err := FromDir(src)
	if err != nil {
		return fmt.Errorf("could not load %s: %s", src, err)
	}

	return pack.SaveDir(dest)
}

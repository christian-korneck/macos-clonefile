//go:build darwin

package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

// Darwin clonefile function:
// https://github.com/apple-oss-distributions/xnu/blob/main/bsd/man/man2/clonefile.2

func cloneFile(source, target string) error {
	if err := unix.Clonefile(source, target, unix.CLONE_NOFOLLOW); err != nil {
		if !errors.Is(err, unix.ENOTSUP) && !errors.Is(err, unix.EXDEV) {
			return fmt.Errorf("clonefile failed: %w", err)
		}

		return nil
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: clonefile [source] [target]")
		os.Exit(1)
	}

	err := cloneFile(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

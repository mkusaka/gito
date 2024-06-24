package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/browser"
)

func main() {
	if err := run(); err != nil {
		_, err = os.Stderr.WriteString(err.Error() + "\n")
		if err != nil {
			fmt.Printf("unexpected error with: %+v", err)
		}
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	remoteName := flag.String("remoteName", "origin", "target remote name")
	targetPath := "."
	if len(os.Args) == 2 {
		targetPath = os.Args[1]
	}

	r, err := git.PlainOpenWithOptions(targetPath, &git.PlainOpenOptions{
		DetectDotGit: true,
	})

	if err != nil {
		return err
	}

	remotes, err := r.Remotes()

	if err != nil {
		return err
	}

	var remote *git.Remote

	if len(remotes) == 1 {
		remote = remotes[0]
	} else {
		for _, r2 := range remotes {
			if r2.Config().Name == *remoteName {
				remote = r2
			}
		}
	}

	if remote == nil {
		if len(remotes) == 0 {
			return errors.New("no remote find")
		} else {
			var remoteNames []string

			for _, r2 := range remotes {
				remoteNames = append(remoteNames, r2.Config().Name)
			}
			return errors.New(fmt.Sprintf("cannot find target remote name: %s, current remotes: %s", *remoteName, strings.Join(remoteNames, ", ")))
		}
	}

	u := remote.Config().URLs[0]

	err = browser.OpenURL(u)
	if err != nil {
		return fmt.Errorf("cannot open browser with error: %+v", err)
	}
	return nil
}

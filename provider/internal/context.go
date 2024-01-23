package internal

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/moby/patternmatcher/ignorefile"
	"github.com/spf13/afero"
	"github.com/tonistiigi/fsutil"
)

type contextHashAccumulator struct {
	dockerContextPath string
	input             bytes.Buffer // This will hold the file info and content bytes to pass to a hash object
}

// hashPath accumulates hashes for files in a directory. If the file is a symlink, the location it
// points to is hashed. If it is a regular file, we hash the contents of the file. In order to
// detect file renames and mode changes, we also write to the accumulator a relative name and file
// mode.
func (accumulator *contextHashAccumulator) hashPath(
	filePath string,
	relativeNameOfFile string,
	fileMode fs.FileMode,
) error {
	hash := sha256.New()

	if fileMode.Type() == fs.ModeSymlink {
		// For symlinks, we hash the symlink _path_ instead of the file content.
		// This will allow us to:
		// a) ignore changes at the symlink target
		// b) detect if the symlink _itself_ changes
		// c) avoid a panic on io.Copy if the symlink target is a directory
		symLinkPath, err := filepath.EvalSymlinks(filePath)
		if err != nil {
			return fmt.Errorf("could not evaluate symlink at %s: %w", filePath, err)
		}
		// Hashed content is the clean, os-agnostic file path:
		_, err = io.Copy(hash, strings.NewReader(filepath.ToSlash(filepath.Clean(symLinkPath))))
		if err != nil {
			return fmt.Errorf("could not copy symlink path %s to hash: %w", filePath, err)
		}
	} else {
		// For regular files, we can hash their content.
		// TODO: consider only hashing file metadata to improve performance
		f, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("could not open file %s: %w", filePath, err)
		}
		defer f.Close()
		_, err = io.Copy(hash, f)
		if err != nil {
			return fmt.Errorf("could not copy file %s to hash: %w", filePath, err)
		}
	}

	// We use "filepath.ToSlash" to return an OS-agnostic path, which uses "/".
	accumulator.input.Write([]byte(filepath.ToSlash(path.Clean(relativeNameOfFile))))
	accumulator.input.Write([]byte(fileMode.String()))
	accumulator.input.Write(hash.Sum(nil))
	accumulator.input.WriteByte(0)
	return nil
}

func (accumulator *contextHashAccumulator) hexSumContext() string {
	h := sha256.New()
	_, err := accumulator.input.WriteTo(h)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func HashContext(dockerContextPath string, dockerfilePath string) (string, error) {
	// exclude all files listed in dockerignore
	ignorePatterns, err := GetIgnorePatterns(afero.NewOsFs(), dockerfilePath, dockerContextPath)
	if err != nil {
		return "", err
	}

	accumulator := contextHashAccumulator{dockerContextPath: dockerContextPath}
	// The dockerfile is always hashed into the digest with the same "name", regardless of its actual
	// name.
	//
	// If the dockerfile is outside the build context, this matches Docker's behavior. Whether it's
	// "foo.Dockerfile" or "bar.Dockerfile", the builder only cares about its contents, not its name.
	//
	// If the dockerfile is inside the build context, we will hash it twice, but that is OK. We hash
	// it here the first time with the name "Dockerfile", and then in the WalkDir loop on we hash it
	// again with its actual name.
	err = accumulator.hashPath(dockerfilePath, "Dockerfile", 0)
	if err != nil {
		return "", fmt.Errorf("error hashing dockerfile %q: %w", dockerfilePath, err)
	}
	err = fsutil.Walk(context.Background(), dockerContextPath, &fsutil.FilterOpt{
		ExcludePatterns: ignorePatterns,
	}, func(filePath string, fileInfo fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}
		// fsutil.Walk makes filePath relative to the root, we join it back to get an absolute path to
		// the file to hash.
		err = accumulator.hashPath(filepath.Join(dockerContextPath, filePath), filePath, fileInfo.Mode())
		if err != nil {
			return fmt.Errorf("error while hashing %q: %w", filePath, err)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("unable to hash build context: %w", err)
	}
	// create a hash of the entire input of the hash accumulator
	return accumulator.hexSumContext(), nil
}

// GetIgnorePatterns returns all patterns to ignore when constructing a build
// context for the given Dockerfile, if any such patterns exist.
//
// Precedence is given to Dockerfile-specific ignore-files as per
// https://docs.docker.com/build/building/context/#filename-and-location.
func GetIgnorePatterns(fs afero.Fs, dockerfilePath, contextRoot string) ([]string, error) {
	paths := []string{
		// Prefer <Dockerfile>.dockerignore if it's present.
		dockerfilePath + ".dockerignore",
		// Otherwise fall back to the ignore-file at the root of our build context.
		filepath.Join(contextRoot, ".dockerignore"),
	}

	// Attempt to parse our candidate ignore-files, skipping any that don't
	// exist.
	for _, p := range paths {
		f, err := fs.Open(p)
		if errors.Is(err, afero.ErrFileNotFound) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("reading %q: %w", p, err)
		}
		defer f.Close()

		ignorePatterns, err := ignorefile.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("unable to parse %q: %w", p, err)
		}
		return ignorePatterns, nil
	}

	return nil, nil
}
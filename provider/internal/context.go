package internal

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	gofs "io/fs"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"github.com/moby/patternmatcher/ignorefile"
	"github.com/muesli/reflow/dedent"
	"github.com/spf13/afero"
	"github.com/tonistiigi/fsutil"

	"github.com/pulumi/pulumi-go-provider/infer"
)

var (
	_ = (infer.Annotated)((*Context)(nil))
	_ = (infer.Annotated)((*BuildContext)(nil))
)

type Context struct {
	Location string `pulumi:"location"`
}

type BuildContext struct {
	Context
	Named NamedContexts `pulumi:"named,optional"`
}

type NamedContexts map[string]Context

func (nc NamedContexts) Map() map[string]string {
	m := map[string]string{}
	for k, v := range nc {
		m[k] = v.Location
	}
	return m
}

func (c *Context) Annotate(a infer.Annotator) {
	a.Describe(&c.Location, dedent.String(`
		Path to use for build context. If omitted, an empty context is used.`,
	))
}

func (bc *BuildContext) Annotate(a infer.Annotator) {
	a.Describe(&bc.Named, dedent.String(`
		Additional build contexts which can be accessed with "FROM name" or
		"--from=name" statements when using Dockerfile 1.4 syntax. Values can
		be local paths, HTTP URLs, or  "docker-image://" images.`,
	))
}

func hashFile(
	h hash.Hash,
	fs fsutil.FS,
	relativePath string,
	fileMode gofs.FileMode,
) error {
	if fileMode.IsDir() {
		return nil
	}
	if !(fileMode.IsRegular() || fileMode.Type() == os.ModeSymlink) {
		return nil
	}

	f, err := fs.Open(relativePath)
	if err != nil {
		return fmt.Errorf("could not open %q: %w", relativePath, err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if errors.Is(err, syscall.EISDIR) {
		// Ignore symlinks to directories.
		return nil
	}
	if err != nil {
		return fmt.Errorf("could not copy %q to hash: %w", relativePath, err)
	}

	h.Write([]byte(filepath.ToSlash(path.Clean(relativePath))))
	h.Write([]byte(fileMode.String()))

	return nil
}

// BuildxContext accumulates hashes for files in a directory. If the file is a symlink, the location it
// points to is hashed. If it is a regular file, we hash the contents of the file. In order to
// detect file renames and mode changes, we also write to the accumulator a relative name and file
// mode.
func BuildxContext(contextPath, dockerfilePath string, namedContexts map[string]string) (string, error) {
	h := sha256.New()
	fs := afero.NewOsFs()

	// Grab .dockerignore if our context and/or Dockerfile is on-disk.
	excludes := []string{}
	if isLocalDir(fs, contextPath) || isLocalFile(fs, dockerfilePath) {
		e, err := GetIgnorePatterns(fs, dockerfilePath, contextPath)
		if err != nil {
			return "", err
		}
		excludes = e
	}

	if isLocalFile(fs, dockerfilePath) {
		err := hashDockerfile(h, dockerfilePath)
		if err != nil {
			return "", nil
		}
	}

	if isLocalDir(fs, contextPath) {
		// Hash our context if it's on-disk.
		fs, err := rootFS(contextPath, excludes)
		if err != nil {
			return "", err
		}
		if _, err := hashPath(h, fs); err != nil {
			return "", err
		}
	}

	// Hash any local named contexts.
	for _, namedContext := range namedContexts {
		if isLocalDir(fs, namedContext) {
			fs, err := rootFS(namedContext, excludes)
			if err != nil {
				return "", err
			}
			if _, err := hashPath(h, fs); err != nil {
				return "", err
			}
		}
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// HashContext is used by the old provider.
func HashContext(contextPath, dockerfilePath string) (string, error) {
	return BuildxContext(contextPath, dockerfilePath, nil)
}

func hashPath(h hash.Hash, fs fsutil.FS) (string, error) {
	err := fs.Walk(context.Background(), "/", func(filePath string, dir gofs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if dir.IsDir() {
			return nil
		}
		// fsutil.Walk makes filePath relative to the root, we join it back to get an absolute path to
		// the file to hash.
		fi, err := dir.Info()
		if err != nil {
			return err
		}
		return hashFile(h, fs, filePath, fi.Mode())
	})
	if err != nil {
		return "", fmt.Errorf("unable to hash build context: %w", err)
	}
	// create a hash of the entire input of the hash accumulator
	return hex.EncodeToString(h.Sum(nil)), nil
}

func hashDockerfile(h hash.Hash, path string) error {
	// The Dockerfile might be capture by .dockerignore, so we explicitly hash
	// its content (but not filename -- to match Docker) in order to detect
	// changes in it.
	df, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading dockerfile %q: %w", path, err)
	}
	_, err = h.Write(df)
	if err != nil {
		return fmt.Errorf("error hashing dockerfile %q: %w", path, err)
	}
	return nil
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
	}

	if isLocalDir(fs, contextRoot) {
		// Otherwise fall back to the ignore-file at the root of our build context.
		paths = append(paths, filepath.Join(contextRoot, ".dockerignore"))
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

func isLocalDir(fs afero.Fs, path string) bool {
	stat, err := fs.Stat(path)
	return err == nil && stat.IsDir()
}

func isLocalFile(fs afero.Fs, path string) bool {
	stat, err := fs.Stat(path)
	return err == nil && !stat.IsDir()
}

// rootFS returns a new fsutil.FS scoped to the given root and with the given
// exclusions.
func rootFS(root string, excludes []string) (fsutil.FS, error) {
	fs, err := fsutil.NewFS(root)
	if err != nil {
		return nil, err
	}
	return fsutil.NewFilterFS(fs, &fsutil.FilterOpt{ExcludePatterns: excludes})
}

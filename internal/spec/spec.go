package spec

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	defaultToolName       = "cherry"
	defaultVersion        = "1.0"
	defaultLanguage       = "go"
	defaultMainFile       = "main.go"
	defaultVersionPackage = "./cmd/version"
	defaultModel          = "master"
)

var (
	specFiles = []string{"cherry.yml", "cherry.yaml", "cherry.json"}

	defaultGoVersions = []string{"1.13"}
	defaultPlatforms  = []string{"linux-386", "linux-amd64", "linux-arm", "linux-arm64", "darwin-386", "darwin-amd64", "windows-386", "windows-amd64"}
)

// Error is the custom error type for spec package.
type Error struct {
	err          error
	SpecNotFound bool
}

func (e *Error) Error() string {
	return e.err.Error()
}

// Unwrap returns the next error in the error chain.
func (e *Error) Unwrap() error {
	return e.err
}

// Build has the specifications for build command.
type Build struct {
	CrossCompile   bool     `json:"crossCompile" yaml:"cross_compile"`
	MainFile       string   `json:"mainFile" yaml:"main_file"`
	BinaryFile     string   `json:"binaryFile" yaml:"binary_file"`
	VersionPackage string   `json:"versionPackage" yaml:"version_package"`
	GoVersions     []string `json:"goVersions" yaml:"go_versions"`
	Platforms      []string `json:"platforms" yaml:"platforms"`
}

// SetDefaults sets default values for empty fields.
func (b *Build) SetDefaults() {
	defaultBinaryFile := "bin/app"
	if wd, err := os.Getwd(); err == nil {
		defaultBinaryFile = "bin/" + filepath.Base(wd)
	}

	if b.MainFile == "" {
		b.MainFile = defaultMainFile
	}

	if b.BinaryFile == "" {
		b.BinaryFile = defaultBinaryFile
	}

	if b.VersionPackage == "" {
		b.VersionPackage = defaultVersionPackage
	}

	if len(b.GoVersions) == 0 {
		b.GoVersions = defaultGoVersions
	}

	if len(b.Platforms) == 0 {
		b.Platforms = defaultPlatforms
	}
}

// FlagSet returns a flag set for input arguments for build command.
func (b *Build) FlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("build", flag.ContinueOnError)
	fs.BoolVar(&b.CrossCompile, "cross-compile", b.CrossCompile, "")
	fs.StringVar(&b.MainFile, "main-file", b.MainFile, "")
	fs.StringVar(&b.BinaryFile, "binary-file", b.BinaryFile, "")
	fs.StringVar(&b.VersionPackage, "version-package", b.VersionPackage, "")

	return fs
}

// Release has the specifications for release command.
type Release struct {
	Model string `json:"model" yaml:"model"`
	Build bool   `json:"build" yaml:"build"`
}

// SetDefaults sets default values for empty fields.
func (r *Release) SetDefaults() {
	if r.Model == "" {
		r.Model = defaultModel
	}
}

// FlagSet returns a flag set for input arguments for release command.
func (r *Release) FlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("release", flag.ContinueOnError)
	fs.StringVar(&r.Model, "model", r.Model, "")
	fs.BoolVar(&r.Build, "build", r.Build, "")

	return fs
}

// Spec has all the specifications for Cherry.
type Spec struct {
	ToolName    string `json:"-" yaml:"-"`
	ToolVersion string `json:"-" yaml:"-"`

	Version     string  `json:"version" yaml:"version"`
	Language    string  `json:"language" yaml:"language"`
	VersionFile string  `json:"versionFile" yaml:"version_file"`
	Build       Build   `json:"build" yaml:"build"`
	Release     Release `json:"release" yaml:"release"`
}

// SetDefaults sets default values for empty fields.
func (s *Spec) SetDefaults() {
	if s.ToolName == "" {
		s.ToolName = defaultToolName
	}

	if s.Version == "" {
		s.Version = defaultVersion
	}

	if s.Language == "" {
		s.Language = defaultLanguage
	}

	s.Build.SetDefaults()
	s.Release.SetDefaults()
}

// Read reads and returns specifications from a file.
func Read() (*Spec, error) {
	for _, file := range specFiles {
		ext := filepath.Ext(file)
		path := filepath.Clean(file)

		f, err := os.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		defer f.Close()

		spec := new(Spec)

		if ext == ".yml" || ext == ".yaml" {
			err = yaml.NewDecoder(f).Decode(spec)
		} else if ext == ".json" {
			err = json.NewDecoder(f).Decode(spec)
		} else {
			return nil, errors.New("unknown spec file")
		}

		if err != nil {
			return nil, err
		}

		return spec, nil
	}

	return nil, &Error{
		err:          errors.New("no spec file found"),
		SpecNotFound: true,
	}
}

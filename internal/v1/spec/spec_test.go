package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		expectedError string
		expectedSpec  *Spec
	}{
		{
			name:          "NoFile",
			path:          "test/null",
			expectedError: "no such file or directory",
		},
		{
			name:          "EmptyYAML",
			path:          "test/empty.yaml",
			expectedError: "EOF",
		},
		{
			name:          "InvalidYAML",
			path:          "test/error.yaml",
			expectedError: "cannot unmarshal",
		},
		{
			name: "Minimum",
			path: "test/min.yaml",
			expectedSpec: &Spec{
				Version:  "1.0",
				Language: "go",
				Test:     Test{},
				Build:    Build{},
				Release: Release{
					Build: true,
				},
			},
		},
		{
			name: "Maximum",
			path: "test/max.yaml",
			expectedSpec: &Spec{
				Version:     "1.0",
				Language:    "go",
				VersionFile: "VERSION",
				Test: Test{
					ReportPath: "coverage",
				},
				Build: Build{
					CrossCompile:   true,
					MainFile:       "main.go",
					BinaryFile:     "bin/cherry",
					VersionPackage: "./cmd/version",
					GoVersions:     []string{"1.10", "1.11.4"},
					Platforms:      []string{"linux-386", "linux-amd64", "darwin-386", "darwin-amd64", "windows-386", "windows-amd64"},
				},
				Release: Release{
					Build: true,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			spec, err := Read(tc.path)

			if tc.expectedError != "" {
				assert.Contains(t, err.Error(), tc.expectedError)
				assert.Empty(t, spec)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedSpec, spec)
			}
		})
	}
}

func TestSpecSetDefaults(t *testing.T) {
	tests := []struct {
		spec         Spec
		expectedSpec Spec
	}{
		{
			Spec{},
			Spec{
				ToolName:    defaultToolName,
				ToolVersion: defaultToolVersion,
				Version:     defaultVersion,
				Language:    defaultLanguage,
				VersionFile: defaultVersionFile,
				Test: Test{
					ReportPath: defaultReportPath,
				},
				Build: Build{
					CrossCompile:   defaultCrossCompile,
					MainFile:       defaultMainFile,
					BinaryFile:     defaultBinaryFile,
					VersionPackage: defaultVersionPackage,
					GoVersions:     defaultGoVersions,
					Platforms:      defaultPlatforms,
				},
				Release: Release{
					Build: defaultBuild,
				},
			},
		},
		{
			Spec{
				Version:     "2.0",
				Language:    "go",
				VersionFile: "version.yaml",
				Test: Test{
					ReportPath: "report",
				},
				Build: Build{
					CrossCompile:   true,
					MainFile:       "cmd/main.go",
					BinaryFile:     "build/app",
					VersionPackage: "./cmd/version",
					GoVersions:     []string{"1.10", "1.11"},
					Platforms:      []string{"linux-amd64", "darwin-amd64", "windows-amd64"},
				},
				Release: Release{
					Build: true,
				},
			},
			Spec{
				ToolName:    defaultToolName,
				ToolVersion: defaultToolVersion,
				Version:     "2.0",
				Language:    "go",
				VersionFile: "version.yaml",
				Test: Test{
					ReportPath: "report",
				},
				Build: Build{
					CrossCompile:   true,
					MainFile:       "cmd/main.go",
					BinaryFile:     "build/app",
					VersionPackage: "./cmd/version",
					GoVersions:     []string{"1.10", "1.11"},
					Platforms:      []string{"linux-amd64", "darwin-amd64", "windows-amd64"},
				},
				Release: Release{
					Build: true,
				},
			},
		},
	}

	for _, tc := range tests {
		tc.spec.SetDefaults()
		assert.Equal(t, tc.expectedSpec, tc.spec)
	}
}

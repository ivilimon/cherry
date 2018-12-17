package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsClean(t *testing.T) {
	tests := []struct {
		name          string
		workDir       string
		expectedError bool
	}{
		{
			name:          "Error",
			workDir:       os.TempDir(),
			expectedError: true,
		},
		{
			name:          "Success",
			workDir:       ".",
			expectedError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			git := New(tc.workDir)
			_, err := git.IsClean()

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetRepoName(t *testing.T) {
	tests := []struct {
		name          string
		workDir       string
		expectedError bool
	}{
		{
			name:          "Error",
			workDir:       os.TempDir(),
			expectedError: true,
		},
		{
			name:          "Success",
			workDir:       ".",
			expectedError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			git := New(tc.workDir)
			_, _, err := git.GetRepoName()

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetBranchName(t *testing.T) {
	tests := []struct {
		name          string
		workDir       string
		expectedError bool
	}{
		{
			name:          "Error",
			workDir:       os.TempDir(),
			expectedError: true,
		},
		{
			name:          "Success",
			workDir:       ".",
			expectedError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			git := New(tc.workDir)
			_, err := git.GetBranchName()

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCommitSHA(t *testing.T) {
	tests := []struct {
		name          string
		workDir       string
		short         bool
		expectedError bool
	}{
		{
			name:          "Error",
			workDir:       os.TempDir(),
			short:         false,
			expectedError: true,
		},
		{
			name:          "FullSHA",
			workDir:       ".",
			short:         false,
			expectedError: false,
		},
		{
			name:          "ShortSHA",
			workDir:       ".",
			short:         true,
			expectedError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			git := New(tc.workDir)
			_, err := git.GetCommitSHA(tc.short)

			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

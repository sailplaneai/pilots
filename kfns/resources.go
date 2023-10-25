package kfns

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

// LoadResourcesFromDir loads all resources in the specified directory.
func LoadResourcesFromDir(resourcesDir string) ([]*yaml.RNode, error) {
	reader := kio.LocalPackageReader{
		PackagePath:           resourcesDir,
		OmitReaderAnnotations: true,
	}

	return reader.Read()
}

// LoadResourcesFromWD loads all resources from a resources directory in the current working directory.
func LoadResourcesFromWD() ([]*yaml.RNode, error) { // Load any resources in the resources directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current working directory")
	}

	resourcesDir := filepath.Join(cwd, "resources")

	return LoadResourcesFromDir(resourcesDir)
}

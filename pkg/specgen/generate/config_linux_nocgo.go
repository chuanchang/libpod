// +build linux,!cgo

package generate

import (
	"errors"

	"github.com/containers/libpod/v2/libpod/image"
	"github.com/containers/libpod/v2/pkg/specgen"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

func getSeccompConfig(s *specgen.SpecGenerator, configSpec *spec.Spec, img *image.Image) (*spec.LinuxSeccomp, error) {
	return nil, errors.New("not implemented")
}

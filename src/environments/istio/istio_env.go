package istio

import (
	"fmt"
	"github.com/fenixsoft/fenix-cli/src/environments"
	"github.com/fenixsoft/fenix-cli/src/environments/kubernetes"
	"github.com/fenixsoft/fenix-cli/src/internal/krew"
	"github.com/pkg/errors"
)

// MUST register AFTER Kubernetes environment
func RegisterEnv() (*environments.Runtime, error) {
	if c, err := New(); err != nil {
		return nil, err
	} else if !krew.IsIstiocltAvailable() {
		return nil, errors.New("istio is not available")
	} else {
		c.Runtime = &environments.Runtime{
			Prefix:         "istioctl",
			Completer:      c.GenericCompleter,
			Executor:       environments.GetDefaultExecutor("istioctl", nil),
			MainSuggestion: options,
			LivePrefix: func() (prefix string, useLivePrefix bool) {
				return fmt.Sprintf("%v > istioctl ", kubernetes.Completer.Namespace), true
			},
		}
		return c.Runtime, nil
	}
}

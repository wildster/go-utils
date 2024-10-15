package goutils

import (
	"github.com/sercand/kuberesolver/v5"
)

func RegisterCore() {
	kuberesolver.RegisterInCluster()
}

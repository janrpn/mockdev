package configuration

import (
	"github.com/thorsager/mockdev/mockhttp"
	"github.com/thorsager/mockdev/mocksnmp"
	"github.com/thorsager/mockdev/mockssh"
)

type Config struct {
	Snmp []*mocksnmp.Configuration `yaml:"snmp"`
	Http []*mockhttp.Configuration `yaml:"http"`
	Ssh  []*mockssh.Configuration  `yaml:"ssh"`
}

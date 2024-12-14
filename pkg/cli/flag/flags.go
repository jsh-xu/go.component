package flag

import (
	"strings"
	goflag "flag"
	"github.com/spf13/pflag"
)


func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name,"_"){
		return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
	}

	return pflag.NormalizedName(name)
}

func InitFlags(flags *pflag.FlagSet){
	flags.SetNormalizeFunc(WordSepNormalizeFunc)
	flags.AddGoFlagSet(goflag.CommandLine)
}
package sha512sum

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/shirou/toybox/common"
)

const binaryName = "sha512sum"

type Option struct {
	helpFlag     bool
	compareFlag  bool
	suppressFlag bool
}

func NewFlagSet() (*flag.FlagSet, *Option) {
	ret := flag.NewFlagSet(binaryName, flag.ExitOnError)

	ret.Usage = func() {
		fmt.Println("sha512sum [-cs] FILE")
		ret.PrintDefaults()
	}
	var opt Option

	ret.BoolVar(&opt.helpFlag, "help", false, "show this message")
	ret.BoolVar(&opt.compareFlag, "c", false, "Compare the digest of the file against this string")
	ret.BoolVar(&opt.suppressFlag, "s", false, "Show nothing. only status code")

	return ret, &opt
}

func Main(stdout io.Writer, args []string) error {
	flagSet, opt := NewFlagSet()
	flagSet.Parse(args)

	if opt.helpFlag {
		flagSet.Usage()
		return nil
	}

	if opt.compareFlag {
		return common.CheckSumCompare("sha512", os.Stdout, flagSet, opt.suppressFlag)
	}
	return common.CheckSumMain("sha512", os.Stdout, flagSet, opt.suppressFlag)
}

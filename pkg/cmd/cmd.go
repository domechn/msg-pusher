/*
 * @Author: domchan
 * @Date: 2018-12-28 15:30:41
 * @Last Modified by: domchan
 * @Last Modified time: 2018-12-28 15:31:43
 */
package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// GratefulQuit returns a chan for blocking, which is closed if the program exits
func GratefulQuit() <-chan struct{} {
	stopCh := make(chan struct{})
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill, os.Interrupt)
	go func() {
		s := <-ch
		log.Infof("Exit signal %s received, clearing memory...", s.String())
		close(stopCh)
	}()
	return stopCh
}

// AddFlags adds all command line flags to the given command.
func AddFlags(rootCmd *cobra.Command) {
	flag.CommandLine.VisitAll(func(gf *flag.Flag) {
		rootCmd.PersistentFlags().AddGoFlag(gf)
	})
}

// PrintFlags logs the flags in the flagset
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Infof("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}

// WordSepNormalizeFunc change all "_" to "-"
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	fmt.Println(2342342)
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}

// PrintSections prints information for all startup parameters
func PrintSections(w io.Writer, name string, flags *pflag.FlagSet, cols int) {
	var wideFS *pflag.FlagSet
	wideFS = pflag.NewFlagSet("", pflag.ExitOnError)
	flags.VisitAll(func(fs *pflag.Flag) {
		wideFS.AddFlag(fs)
	})

	var zzz string
	if cols > 24 {
		zzz = strings.Repeat("z", cols-24)
		wideFS.Int(zzz, 0, strings.Repeat("z", cols-24))
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\n%s flags:\n\n%s", name, wideFS.FlagUsagesWrapped(cols))

	if cols > 24 {
		i := strings.Index(buf.String(), zzz)
		lines := strings.Split(buf.String()[:i], "\n")
		fmt.Fprint(w, strings.Join(lines[:len(lines)-1], "\n"))
		fmt.Fprintln(w)
	} else {
		fmt.Fprint(w, buf.String())
	}
}

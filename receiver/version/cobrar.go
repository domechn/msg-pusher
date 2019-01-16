/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : cobrar.go
#   Created       : 2019/1/8 15:57
#   Last Modified : 2019/1/8 15:57
#   Describe      :
#
# ====================================================*/
package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Command returns a command used to print version information.
func Command() *cobra.Command {
	var short bool
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints out build version information",
		Run: func(cmd *cobra.Command, args []string) {
			if short {
				fmt.Println(Info)
			} else {
				fmt.Println(Info.LongForm())
			}
		},
	}
	cmd.PersistentFlags().BoolVarP(&short, "short", "s", short, "Displays a short form of the version information")
	return cmd
}

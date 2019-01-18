/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : main.go
#   Created       : 2019/1/16 13:26
#   Last Modified : 2019/1/16 13:26
#   Describe      :
#
# ====================================================*/
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"time"

	"uuabc.com/sendmsg/config"
	"uuabc.com/sendmsg/pkg/cmd"
	"uuabc.com/sendmsg/pkg/log"
	"uuabc.com/sendmsg/receiver/version"
	"uuabc.com/sendmsg/sender"
)

type Options struct {
	host       string
	port       int
	configPath string
	logPath    string
	logLevel   string
}

var (
	opts = &Options{
		host:       "0.0.0.0",
		port:       8991,
		configPath: "/app/sendmsg/conf/config.yaml",
		logPath:    "/app/sendmsg/log/log.log",
		logLevel:   "info",
	}

	rootCmd = &cobra.Command{
		Use:          "sender-server",
		Short:        "Used to send messages.",
		SilenceUsage: true,
	}

	startCmd = &cobra.Command{
		Use: "start",
		Long: `Start the service to listen for message queue information 
				and send the received information to the specified client.`,
		RunE: start,
	}
)

func init() {
	startCmd.PersistentFlags().StringVarP(&opts.host, "host", "s", opts.host, "host for service startup")
	startCmd.PersistentFlags().IntVarP(&opts.port, "port", "p", opts.port, "port for service startup")
	startCmd.PersistentFlags().StringVarP(&opts.configPath, "config-path", "f", opts.configPath, "the path of the config file")
	startCmd.PersistentFlags().StringVar(&opts.logPath, "log-path", opts.logPath, "the location of the log file output")
	startCmd.PersistentFlags().StringVar(&opts.logLevel, "log-level", opts.logLevel, "log file output level")

	cmd.AddFlags(rootCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(version.Command())
}

func start(_ *cobra.Command, _ []string) error {
	stopC := cmd.GratefulQuit()
	var err error

	printFlags()
	rand.Seed(time.Now().UnixNano())

	// init log
	log.Init("sender", opts.logPath, opts.logLevel)
	if err = config.Init(opts.configPath); err != nil {
		return err
	}

	if err = sender.Init(); err != nil {
		return err
	}

	go func() {
		<-stopC
		logrus.Info("stopping server now")
		if err := sender.Close(); err != nil {
			logrus.Errorf("Server Close:", err)
		}
		os.Exit(0)
	}()

	return sender.Start()
}

func printFlags() {
	logrus.WithField("Host", opts.host).Info()
	logrus.WithField("Post", opts.port).Info()
	logrus.WithField("Config-Path", opts.configPath).Info()
	logrus.WithField("Log-Path", opts.logPath).Info()
	logrus.WithField("Log-Level", opts.logLevel).Info()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf("serve failed ,error: %v\n", err)
		os.Exit(-1)
	}
}

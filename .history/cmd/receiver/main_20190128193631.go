/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : main.go
#   Created       : 2019-01-08 14:26:16
#   Last Modified : 2019-01-08 14:26:16
#   Describe      :
#
# ====================================================*/
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hiruok/msg-pusher/config"
	"github.com/hiruok/msg-pusher/corn"
	"github.com/hiruok/msg-pusher/pkg/log"
	"github.com/hiruok/msg-pusher/pkg/opentracing"
	"github.com/hiruok/msg-pusher/receiver"
	"github.com/hiruok/msg-pusher/receiver/version"
	"github.com/sirupsen/logrus"

	"github.com/hiruok/msg-pusher/pkg/cmd"
	"github.com/spf13/cobra"
)

var (
	opts = &Options{
		host:       "0.0.0.0",
		port:       8990,
		configPath: "/app/msg-pusher/conf/conf.yaml",
		logPath:    "/app/msg-pusher/log",
		logLevel:   "info",
	}

	rootCmd = &cobra.Command{
		Use:          "receiver-server",
		Short:        "The producer of the message service.",
		SilenceUsage: true,
	}

	startCmd = &cobra.Command{
		Use: "start",
		Long: `Start the service to receive the parameters from the user 
		and send the parameters to mq for consumption by the consumer`,
		RunE: start,
	}
)

const (
	defaultTimeout = time.Second * 10
)

// Options 启动参数
type Options struct {
	host        string
	port        int
	configPath  string
	logPath     string
	logLevel    string
	addrJaeger  string
	addrMonitor string
}

func init() {
	startCmd.PersistentFlags().StringVarP(&opts.host, "host", "s", opts.host, "host for service startup")
	startCmd.PersistentFlags().IntVarP(&opts.port, "port", "p", opts.port, "port for service startup")
	startCmd.PersistentFlags().StringVarP(&opts.configPath, "config-path", "f", opts.configPath, "the path of the config file")
	startCmd.PersistentFlags().StringVar(&opts.logPath, "log-path", opts.logPath, "the location of the log file output")
	startCmd.PersistentFlags().StringVar(&opts.logLevel, "log-level", opts.logLevel, "log file output level")
	startCmd.PersistentFlags().StringVar(&opts.addrJaeger, "addr-jaeger", opts.addrJaeger, "the address of jaeger")
	startCmd.PersistentFlags().StringVar(&opts.addrMonitor, "addr-monitor", opts.addrMonitor, "the address of monitor(prometheus)")

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
	log.Init("receiver", opts.logPath, opts.logLevel)

	if err = config.Init(opts.configPath); err != nil {
		return err
	}

	// init opentracing
	if err = opentracing.New(opentracing.InitConfig(opts.addrJaeger)).Setup(); err != nil {
		return err
	}

	r := mux.NewRouter()

	if err := receiver.Init(r, opts.addrMonitor); err != nil {
		return err
	}

	// 启动定时存数据库的任务
	corn.Start()

	svr := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", opts.host, opts.port),
		Handler:      r,
		ReadTimeout:  defaultTimeout,
		WriteTimeout: defaultTimeout,
		IdleTimeout:  defaultTimeout,
	}

	// grateful quit
	go func() {
		<-stopC
		logrus.Info("stopping server now")
		receiver.Close()
		if err := svr.Close(); err != nil {
			logrus.Errorf("Server Close:", err)
		}
	}()
	// start server
	if err = svr.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logrus.Info("Server closed under request\n")
			return nil
		} 
			logrus.Infof("Server closed unexpect, %s\n", err.Error())
	}
	return err
}

func printFlags() {
	logrus.WithField("Host", opts.host).Info()
	logrus.WithField("Post", opts.port).Info()
	logrus.WithField("Config-Path", opts.configPath).Info()
	logrus.WithField("Log-Path", opts.logPath).Info()
	logrus.WithField("Log-Level", opts.logLevel).Info()
	logrus.WithField("Addr-Jaeger", opts.addrJaeger).Info()
	logrus.WithField("Addr-Monitor", opts.addrMonitor).Info()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf("serve failed ,error: %v\n", err)
		os.Exit(-1)
	}
}

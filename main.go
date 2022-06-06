package main

import (
	"github.com/ClareChu/proxy-istio-demo/pkg/proxy/router"
	"k8s.io/klog/v2"
	"os"
)

func init() {
	klog.InitFlags(nil)
	//pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

func main() {
	rootCmd := router.GetRootCmd(os.Args[1:])
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

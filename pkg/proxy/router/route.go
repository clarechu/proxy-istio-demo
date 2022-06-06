package router

import (
	"github.com/spf13/cobra"
)

type Root struct {
	Port    int32 `json:"port"`
	Timeout int32 `json:"timeout"`
}

// GetRootCmd returns the root of the cobra command-tree.
func GetRootCmd(args []string) *cobra.Command {
	ag := &Root{}
	rootCmd := &cobra.Command{
		Use:   "http-proxy",
		Short: "http-proxy ...",
		Long: `
Tips  Find more information at: https://github.com/clarechu/proxy-istio-demo
Example:

`,
		Run: func(cmd *cobra.Command, args []string) {
			server := NewServer(ag)
			server.Run()
		},
	}
	addFlag(rootCmd, ag)
	rootCmd.AddCommand(VersionCommand())
	return rootCmd
}

func addFlag(rootCmd *cobra.Command, root *Root) {
	rootCmd.PersistentFlags().Int32Var(&root.Port, "port", 7777, "proxy server ports")
	rootCmd.PersistentFlags().Int32Var(&root.Timeout, "timeout", 5, "proxy server timeout")

}

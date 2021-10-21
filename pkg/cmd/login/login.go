// Package get is for the get command
package login

import (
	"github.com/brevdev/brev-cli/pkg/auth"

	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)



type LoginOptions struct{}

var (
	BrevAuthDomain = "https://brev.dev/oauth/device/code"
	BrevClientID   = "foo"
)

func NewCmdLogin() *cobra.Command {
	opts := LoginOptions{}

	cmd := &cobra.Command{
		Use:                   "login",
		DisableFlagsInUseLine: true,
		Short:                 "log into brev",
		Long:                  "log into brev",
		Example:               "brev login",
		Args:                  cobra.NoArgs,
		// ValidArgsFunction: util.ResourceNameCompletionFunc(f, "pod"),
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(opts.Complete(cmd, args))
			cmdutil.CheckErr(opts.Validate(cmd, args))
			cmdutil.CheckErr(opts.RunLogin(cmd, args))
		},
	}
	return cmd
}

func (o *LoginOptions) Complete(cmd *cobra.Command, args []string) error {
	// return fmt.Errorf("not implemented")
	return nil
}

func (o *LoginOptions) Validate(cmd *cobra.Command, args []string) error {
	// return fmt.Errorf("not implemented")
	return nil
}

func (o *LoginOptions) RunLogin(cmd *cobra.Command, args []string) error {
	return auth.Login()
}

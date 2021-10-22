// Package get is for the get command
package get

import (
	"github.com/brevdev/brev-cli/pkg/brev_api"
	"github.com/brevdev/brev-cli/pkg/cmdcontext"
	"github.com/brevdev/brev-cli/pkg/terminal"
	"github.com/spf13/cobra"
)

func getOrgs() []brev_api.Organization {
	client, _ := brev_api.NewClient()
	orgs, _ := client.GetOrgs()

	return orgs
}

func getWorkspaces(orgID string) []brev_api.Workspace {
	// orgID := getOrgID(orgName)

	client, _ := brev_api.NewClient()
	workspaces, _ := client.GetWorkspaces(orgID)

	return workspaces
}

func getMe() brev_api.User {
	client, _ := brev_api.NewClient()
	user, _ := client.GetMe()
	return *user
}

func NewCmdGet(t *terminal.Terminal) *cobra.Command {
	// opts := SshOptions{}

	cmd := &cobra.Command{
		Use: "get",
		// Annotations: map[string]string{"project": ""},
		Short:   "Get stuff",
		Long:    "Get stuff but longer.",
		Example: `brev get [stuff]`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := cmdcontext.InvokeParentPersistentPreRun(cmd, args)
			if err != nil {
				return err
			}

			// _, err = brev_api.CheckOutsideBrevErrorMessage(t)
			return err
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return []string{}, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveDefault
		},
	}

	cmd.AddCommand(newCmdOrg(t))
	cmd.AddCommand(newCmdWorkspace(t))
	cmd.AddCommand(newCmdMe(t))

	return cmd
}

func newCmdOrg(t *terminal.Terminal) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:     "orgs",
		Short:   "List your Brev orgs.",
		Long:    "List your Brev orgs.",
		Example: `  brev get orgs`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listOrgs(t)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name of the endpoint")

	return cmd
}

func listOrgs(t *terminal.Terminal) error {
	orgs := getOrgs()
	for _, v := range orgs {
		t.Vprint(v.Name + " id:" + t.Yellow(v.Id))
	}
	return nil
}

func newCmdWorkspace(t *terminal.Terminal) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:     "workspace",
		Short:   "List your Brev workspaces.",
		Long:    "List your Brev workspaces.",
		Example: `  brev get workspaces`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listWorkspaces(t)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name of the endpoint")

	return cmd
}

func listWorkspaces(t *terminal.Terminal) error {
	orgs := getOrgs()
	// var workspaces map[string]interface{};

	for _, v := range orgs {

		ws := getWorkspaces(v.Id)

		if len(ws) == 0 {
			t.Vprint("0 Workspaces in Org: " + v.Name + " id:" + t.Yellow(v.Id))
		} else {
			t.Vprint("Workspaces in Org: " + v.Name + " id:" + t.Yellow(v.Id) + ":")
		}

		for _, w := range ws {
			t.Vprint("\t" + w.Name + " id: " + t.Yellow(w.Id))
		}

	}
	return nil
}
func newCmdMe(t *terminal.Terminal) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "me",
		Short: "return info about the current authenticated user",
		Long:  "return info about the current authenticated user",
		Example: `brev get me

		User ID: c0wj3ro`,
		RunE: func(cmd *cobra.Command, args []string) error {
			me := getMe()
			t.Vprintf("User ID: %s", me.Id)
			return nil
		},
	}

	return cmd
}

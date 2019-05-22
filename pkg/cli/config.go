package cli

import (
	gpflag "github.com/octago/sflags/gen/gpflag"
	cobra "github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var config = Config{}
	var cmd = &cobra.Command{
		Use: "selfcare",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	if err := gpflag.ParseTo(&config, cmd.PersistentFlags()); err != nil {
		panic(err)
	}
	return cmd
}

package cmd

import (
	"github.com/go-xorm/xorm"
	"github.com/spf13/cobra"
)

type Options struct {
	db *xorm.Engine
}

func NewCreateDummyUserCommand(opts *Options) *cobra.Command {
	return &cobra.Command{
		Use: "create_dummy_user",
		Run: func(cmd *cobra.Command, args []string) {
			println("create user")
		},
	}
}

func NewDeleteDummyUserCommand(opts *Options) *cobra.Command {
	return &cobra.Command{
		Use: "delete_dummy_user",
		Run: func(cmd *cobra.Command, args []string) {
			println("delete user")
		},
	}
}

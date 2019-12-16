package main

import (
	"fmt"
	"github.com/KouT127/Attendance-management/backend/attendance/cmd"
	"github.com/KouT127/Attendance-management/backend/config"
	"github.com/KouT127/Attendance-management/backend/database"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "Example",
	}
	configs.Init(configs.Development)
	c := configs.NewConfig()
	database.Init(c)
	db := database.NewDB()
	opts := &cmd.Options{
		db,
	}

	rootCmd.AddCommand(cmd.NewCreateDummyUserCommand(opts))
	rootCmd.AddCommand(cmd.NewDeleteDummyUserCommand(opts))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

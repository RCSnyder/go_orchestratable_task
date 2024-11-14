package main

import (
    "os"
    "github.com/spf13/cobra"
)

func main() {
    rootCmd := &cobra.Command{Use: "ot"}
    rootCmd.AddCommand(newRunCommand())
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

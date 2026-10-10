/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-26 17:05:18
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-10 11:13:20
 * @FilePath: /WindLiberity/cmd/root.go
 * @Description:
 *
 */
package cmd

import (
	"WindLiberity/cmd/migrate"
	"WindLiberity/cmd/server"
	"WindLiberity/cmd/version"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	"WindLiberity",
	Short:  "Anime app",
	Long:	"An anime app for anime viewers",
	Args:	args,
	Run:	func(cmd *cobra.Command, args []string) {
		log.Fatalf("unrecongnized command cmd: %v args: %v", cmd.Name(), args)
	},
}


func args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("please select option")
	}
	return nil
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
}

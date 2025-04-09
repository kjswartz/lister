/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/kjswartz/lister/pkg/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add list",
	Long: `Add a list to the lists. Pass in the -n name flag to provide a name for the list.
	For example: 
		lister add -n "read a book".`,
	Run: addFunc,
}

var name string

func init() {
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Set the name of the list")
}

func addFunc(cmd *cobra.Command, args []string) {
	if name == "" {
		fmt.Println("No name provided.")
		return
	}

	db, err := sql.Open("sqlite3", utils.DBPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	insertSQL := `INSERT INTO lists (name) VALUES (?);`
	_, err = db.Exec(insertSQL, name)
	if err != nil {
		fmt.Println("Error inserting item:", err)
		return
	}
	fmt.Println("Added list:", name)
}

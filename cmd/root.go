/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"

	"github.com/kjswartz/lister/cmd/list"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"

	"github.com/kjswartz/lister/pkg/utils"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lister",
	Short: "Simple command line list app",
	Long:  `A simple command line list app that allows you to manage lists of items.`,
	Run:   rootFunc,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// var showLists bool

func init() {
	cobra.OnInitialize(initDB)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initDB() {
	db, err := sql.Open("sqlite3", utils.DBPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	createTableListsSQL := `CREATE TABLE IF NOT EXISTS lists (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL
  );`

	_, err = db.Exec(createTableListsSQL)
	if err != nil {
		fmt.Println("Error creating lists table:", err)
		return
	}

	createTableListItemsSQL := `CREATE TABLE IF NOT EXISTS list_items (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"list_id" INTEGER NOT NULL,
		"position" INTEGER NOT NULL,
		"description" TEXT NOT NULL
  );`

	_, err = db.Exec(createTableListItemsSQL)
	if err != nil {
		fmt.Println("Error creating list_items table:", err)
		return
	}
}

func rootFunc(cmd *cobra.Command, args []string) {
	db, err := sql.Open("sqlite3", utils.DBPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	var query string
	var rows *sql.Rows

	query = `SELECT id, name FROM lists ORDER BY name ASC`
	rows, err = db.Query(query)

	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		printList(id, name)
	}
}

// printList prints a list with its ID and name.
func printList(id int, name string) {
	fmt.Printf("%d | %s\n", id, name)
}

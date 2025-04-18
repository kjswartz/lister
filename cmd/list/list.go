/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/kjswartz/lister/pkg/utils"
	"github.com/spf13/cobra"
)

type ListItem struct {
	ID          int
	Description string
}

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Long: `List commands for managing list items.
		
	Available Flags:
		--add    |-a 	Add items to a list.
		--remove |-r 	Remove items from a list.

	For example:
		lister list 123 -a "read a book".
		lister list 123 -r 456.
`,
	Run: listFunc,
}

var (
	itemDescription string
	itemID          int
)

func init() {
	ListCmd.Flags().StringVarP(&itemDescription, "add", "a", "", "Item to add to the list item")
	ListCmd.Flags().IntVarP(&itemID, "remove", "r", 0, "ID of the list item to remove")
}

func listFunc(cmd *cobra.Command, args []string) {
	listID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid list ID provided. It must be an integer.")
		return
	}

	db, err := sql.Open("sqlite3", utils.DBPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// get list from lists table
	var listName string
	err = db.QueryRow("SELECT name FROM lists WHERE id = ?", listID).Scan(&listName)
	if err != nil {
		fmt.Printf("Error querying list with ID '%d': %v\n", listID, err)
		return
	}

	// Handle adding an item to the list if -a flag is provided
	if itemDescription != "" {
		err := addItemToList(db, listID, itemDescription)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Added item '%s' to list '%s'\n", itemDescription, listName)
	}

	// Handle removing an item from the list if -r flag is provided
	if itemID != 0 {
		_, err := db.Exec("DELETE FROM list_items WHERE id = ? AND list_id = ?", itemID, listID)
		if err != nil {
			fmt.Printf("Error removing item with ID '%d' from list '%s': %v\n", itemID, listName, err)
			return
		}
		fmt.Printf("Removed item with ID '%d' from list '%s'\n", itemID, listName)
	}

	// List all items in the list passed in via -l flag
	listItems, err := findListItems(db, listID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("  Items in %s list\n", listName)
	for _, item := range listItems {
		fmt.Printf("    %d | %s\n", item.ID, item.Description)
	}
}

func addItemToList(db *sql.DB, listID int, description string) error {
	insertSQL := `INSERT INTO list_items (list_id, description) VALUES (?, ?);`
	_, err := db.Exec(insertSQL, listID, description)
	if err != nil {
		return fmt.Errorf("error inserting item into list: %v", err)
	}
	return nil
}

func findListItems(db *sql.DB, listID int) ([]ListItem, error) {
	rows, err := db.Query("SELECT id, description FROM list_items WHERE list_id = ? ORDER BY id ASC", listID)
	if err != nil {
		return nil, fmt.Errorf("error querying list items for list ID '%d': %v", listID, err)
	}
	defer rows.Close()

	listItems := []ListItem{}
	for rows.Next() {
		var (
			id          int
			description string
		)
		if err := rows.Scan(&id, &description); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		listItems = append(listItems, ListItem{ID: id, Description: description})
	}
	return listItems, nil
}

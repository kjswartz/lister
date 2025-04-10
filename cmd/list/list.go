/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"database/sql"
	"fmt"

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
		
	Available Commands:
		add 		Add items to a list.
		remove 	Remove items from a list.
		move 		Move an item within a list.

	For example:
		lister list add -l 123 -n "read a book".
		lister list remove -l 123 -i 456.
		lister list move -l 123 -i 456 (-u|-d).
`,
	Run: listFunc,
}

var (
	listID          int
	itemDescription string
	itemID          int
	// moveUp   bool
	// moveDown bool
)

func init() {
	ListCmd.Flags().IntVarP(&listID, "list", "l", 0, "ID of the list")
	ListCmd.Flags().IntVarP(&itemID, "id", "i", 0, "ID of the list item")
	ListCmd.Flags().StringVarP(&itemDescription, "add", "a", "", "Item to add to the list item")
	// ListCmd.Flags().BoolVarP(&moveUp, "up", "u", false, "Move item up in the list")
	// ListCmd.Flags().BoolVarP(&moveDown, "down", "d", false, "Move item down in the list")
}

func listFunc(cmd *cobra.Command, args []string) {
	if listID == 0 {
		fmt.Println("No list ID provided.")
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

	if itemDescription != "" {
		// Handle adding an item to a list
		insertSQL := `INSERT INTO list_items (list_id, description) VALUES (?, ?);`
		_, err = db.Exec(insertSQL, listID, itemDescription)
		if err != nil {
			fmt.Println("Error adding item:", err)
			return
		}
		fmt.Printf("Added item '%s' to list with ID %d\n", itemDescription, listID)
	}

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

func findListItems(db *sql.DB, listID int) ([]ListItem, error) {
	rows, err := db.Query("SELECT id, description FROM list_items WHERE list_id = ?", listID)
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

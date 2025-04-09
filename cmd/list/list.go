/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
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
	listName string
	name     string
	itemID   int
	moveUp   bool
	moveDown bool
)

func init() {
	ListCmd.Flags().StringVarP(&listName, "list", "l", "", "Name of the list")
	ListCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the list item")
	ListCmd.Flags().IntVarP(&itemID, "id", "i", 0, "ID of the list item")
	ListCmd.Flags().BoolVarP(&moveUp, "up", "u", false, "Move item up in the list")
	ListCmd.Flags().BoolVarP(&moveDown, "down", "d", false, "Move item down in the list")
}

func listFunc(cmd *cobra.Command, args []string) {
	fmt.Println("HERE")
}

// func addFunc(cmd *cobra.Command, args []string) {
// 	// Handle list creation if only list name is provided
// 	if listName != "" && name == "" && itemID == 0 && !moveUp && !moveDown {
// 		addList(listName)
// 		return
// 	}

// 	// Handle adding an item to a list
// 	if listName != "" && name != "" && itemID == 0 && !moveUp && !moveDown {
// 		addItemToList(listName, name)
// 		return
// 	}

// 	// Handle removing an item from a list
// 	if listName != "" && itemID > 0 && name == "" && !moveUp && !moveDown {
// 		removeItemFromList(listName, itemID)
// 		return
// 	}

// 	// Handle moving an item up or down in a list
// 	if listName != "" && itemID > 0 && (moveUp || moveDown) {
// 		moveItemInList(listName, itemID, moveUp)
// 		return
// 	}

// 	fmt.Println("Invalid combination of flags. Please check the usage.")
// }

// func addList(listName string) {
// 	db, err := sql.Open("sqlite3", utils.DBPath)
// 	if err != nil {
// 		fmt.Println("Error opening database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	insertSQL := `INSERT INTO lists (name) VALUES (?);`
// 	_, err = db.Exec(insertSQL, listName)
// 	if err != nil {
// 		fmt.Println("Error creating list:", err)
// 		return
// 	}
// 	fmt.Printf("Created list: %s\n", listName)
// }

// func addItemToList(listName, itemName string) {
// 	db, err := sql.Open("sqlite3", utils.DBPath)
// 	if err != nil {
// 		fmt.Println("Error opening database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	// First, find the list ID
// 	var listID int
// 	err = db.QueryRow("SELECT id FROM lists WHERE name = ?", listName).Scan(&listID)
// 	if err != nil {
// 		fmt.Printf("List '%s' not found: %v\n", listName, err)
// 		return
// 	}

// 	// Insert the new item
// 	insertSQL := `INSERT INTO list_items (list_id, name) VALUES (?, ?);`
// 	_, err = db.Exec(insertSQL, listID, itemName)
// 	if err != nil {
// 		fmt.Println("Error adding item:", err)
// 		return
// 	}
// 	fmt.Printf("Added item '%s' to list '%s'\n", itemName, listName)
// }

// func removeItemFromList(listName string, id int) {
// 	db, err := sql.Open("sqlite3", utils.DBPath)
// 	if err != nil {
// 		fmt.Println("Error opening database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	// First, find the list ID
// 	var listID int
// 	err = db.QueryRow("SELECT id FROM lists WHERE name = ?", listName).Scan(&listID)
// 	if err != nil {
// 		fmt.Printf("List '%s' not found: %v\n", listName, err)
// 		return
// 	}

// 	// Delete the item
// 	deleteSQL := `DELETE FROM list_items WHERE id = ? AND list_id = ?;`
// 	result, err := db.Exec(deleteSQL, id, listID)
// 	if err != nil {
// 		fmt.Println("Error removing item:", err)
// 		return
// 	}

// 	rowsAffected, _ := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		fmt.Printf("Item with ID %d not found in list '%s'\n", id, listName)
// 		return
// 	}

// 	fmt.Printf("Removed item with ID %d from list '%s'\n", id, listName)
// }

// func moveItemInList(listName string, id int, moveUp bool) {
// 	db, err := sql.Open("sqlite3", utils.DBPath)
// 	if err != nil {
// 		fmt.Println("Error opening database:", err)
// 		return
// 	}
// 	defer db.Close()

// 	// First, find the list ID
// 	var listID int
// 	err = db.QueryRow("SELECT id FROM lists WHERE name = ?", listName).Scan(&listID)
// 	if err != nil {
// 		fmt.Printf("List '%s' not found: %v\n", listName, err)
// 		return
// 	}

// 	// Get current position
// 	var position int
// 	err = db.QueryRow("SELECT position FROM list_items WHERE id = ? AND list_id = ?", id, listID).Scan(&position)
// 	if err != nil {
// 		fmt.Printf("Item with ID %d not found: %v\n", id, err)
// 		return
// 	}

// 	// Calculate new position
// 	newPosition := position
// 	if moveUp {
// 		newPosition--
// 	} else {
// 		newPosition++
// 	}

// 	// Check if new position is valid
// 	var maxPosition int
// 	err = db.QueryRow("SELECT MAX(position) FROM list_items WHERE list_id = ?", listID).Scan(&maxPosition)
// 	if err != nil {
// 		fmt.Println("Error checking position range:", err)
// 		return
// 	}

// 	if newPosition < 1 || newPosition > maxPosition {
// 		fmt.Println("Cannot move item further in that direction")
// 		return
// 	}

// 	// Begin transaction to swap positions
// 	tx, err := db.Begin()
// 	if err != nil {
// 		fmt.Println("Error starting transaction:", err)
// 		return
// 	}

// 	// Update the other item that will swap positions
// 	_, err = tx.Exec("UPDATE list_items SET position = ? WHERE list_id = ? AND position = ?",
// 		position, listID, newPosition)
// 	if err != nil {
// 		tx.Rollback()
// 		fmt.Println("Error updating position:", err)
// 		return
// 	}

// 	// Update our item with the new position
// 	_, err = tx.Exec("UPDATE list_items SET position = ? WHERE id = ?",
// 		newPosition, id)
// 	if err != nil {
// 		tx.Rollback()
// 		fmt.Println("Error updating position:", err)
// 		return
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		fmt.Println("Error committing transaction:", err)
// 		return
// 	}

// 	direction := "up"
// 	if !moveUp {
// 		direction = "down"
// 	}
// 	fmt.Printf("Moved item with ID %d %s in list '%s'\n", id, direction, listName)
// }

# Lister

Lister is a simple command-line application for managing lists and their items. It allows you to create lists, add items to lists, remove items, and view all items in a list.

## Features

- Create and manage multiple lists.
- Add items to a specific list.
- Remove items from a list.
- View all items in a list.
- Move items within a list (up or down).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kjswartz/lister.git
   ```
2. Navigate to the project directory:
   ```bash
   cd lister
   ```
3. Build the application:
   ```bash
   go build -o lister
   ```

## Usage

Run the application using the following command:
```bash
./lister [command] [flags]
```

### Commands

#### `list`
Manage list items.

- **Add an item to a list**:
  ```bash
  ./lister list -l <list_id> -a "<item_description>"
  ```
  Example:
  ```bash
  ./lister list -l 1 -a "Read a book"
  ```

- **Remove an item from a list**:
  ```bash
  ./lister list -l <list_id> -r <item_id>
  ```
  Example:
  ```bash
  ./lister list -l 1 -r 2
  ```

- **View all items in a list**:
  ```bash
  ./lister list -l <list_id>
  ```
  Example:
  ```bash
  ./lister list -l 1
  ```

#### `lister`
View all available lists:
```bash
./lister
```

### Database

The application uses SQLite as its database. The database file path is configured in the `utils.DBPath` variable.

### Example Workflow

1. Create a new list:
   ```bash
   ./lister list -l 1 -a "Groceries"
   ```

2. Add items to the list:
   ```bash
   ./lister list -l 1 -a "Buy milk"
   ./lister list -l 1 -a "Buy bread"
   ```

3. View all items in the list:
   ```bash
   ./lister list -l 1
   ```

4. Remove an item from the list:
   ```bash
   ./lister list -l 1 -r 2
   ```

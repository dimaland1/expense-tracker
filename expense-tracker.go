package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
}

type ExpenseTracker struct {
	Expenses []Expense `json:"expenses"`
	NextID   int       `json:"next_id"`
}

const dataFile = "expenses.json"

func main() {
	tracker := loadExpenseTracker()

	app := &cli.App{
		Name:  "expense-tracker",
		Usage: "Track your expenses",
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add a new expense",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Required: true},
					&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}, Required: true},
					&cli.StringFlag{Name: "category", Aliases: []string{"c"}},
				},
				Action: func(c *cli.Context) error {
					return addExpense(c, tracker)
				},
			},
			{
				Name:  "update",
				Usage: "Update an existing expense",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "id", Required: true},
					&cli.StringFlag{Name: "description", Aliases: []string{"d"}},
					&cli.Float64Flag{Name: "amount", Aliases: []string{"a"}},
					&cli.StringFlag{Name: "category", Aliases: []string{"c"}},
				},
				Action: func(c *cli.Context) error {
					return updateExpense(c, tracker)
				},
			},
			{
				Name:  "delete",
				Usage: "Delete an expense",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "id", Required: true},
				},
				Action: func(c *cli.Context) error {
					return deleteExpense(c, tracker)
				},
			},
			{
				Name:  "list",
				Usage: "List all expenses",
				Action: func(c *cli.Context) error {
					return listExpenses(tracker)
				},
			},
			{
				Name:  "summary",
				Usage: "View expense summary",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "month", Aliases: []string{"m"}},
				},
				Action: func(c *cli.Context) error {
					return showSummary(c, tracker)
				},
			},
			{
				Name:  "export",
				Usage: "Export expenses to CSV",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					return exportToCSV(c, tracker)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func loadExpenseTracker() *ExpenseTracker {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return &ExpenseTracker{NextID: 1}
	}

	var tracker ExpenseTracker
	err = json.Unmarshal(data, &tracker)
	if err != nil {
		fmt.Println("Error loading data:", err)
		return &ExpenseTracker{NextID: 1}
	}

	return &tracker
}

func saveExpenseTracker(tracker *ExpenseTracker) error {
	data, err := json.MarshalIndent(tracker, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dataFile, data, 0644)
}

func addExpense(c *cli.Context, tracker *ExpenseTracker) error {
	expense := Expense{
		ID:          tracker.NextID,
		Date:        time.Now(),
		Description: c.String("description"),
		Amount:      c.Float64("amount"),
		Category:    c.String("category"),
	}

	tracker.Expenses = append(tracker.Expenses, expense)
	tracker.NextID++

	err := saveExpenseTracker(tracker)
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
	return nil
}

func updateExpense(c *cli.Context, tracker *ExpenseTracker) error {
	id := c.Int("id")
	for i, expense := range tracker.Expenses {
		if expense.ID == id {
			if c.IsSet("description") {
				tracker.Expenses[i].Description = c.String("description")
			}
			if c.IsSet("amount") {
				tracker.Expenses[i].Amount = c.Float64("amount")
			}
			if c.IsSet("category") {
				tracker.Expenses[i].Category = c.String("category")
			}

			err := saveExpenseTracker(tracker)
			if err != nil {
				return err
			}

			fmt.Printf("Expense updated successfully (ID: %d)\n", id)
			return nil
		}
	}

	return fmt.Errorf("expense with ID %d not found", id)
}

func deleteExpense(c *cli.Context, tracker *ExpenseTracker) error {
	id := c.Int("id")
	for i, expense := range tracker.Expenses {
		if expense.ID == id {
			tracker.Expenses = append(tracker.Expenses[:i], tracker.Expenses[i+1:]...)

			err := saveExpenseTracker(tracker)
			if err != nil {
				return err
			}

			fmt.Printf("Expense deleted successfully (ID: %d)\n", id)
			return nil
		}
	}

	return fmt.Errorf("expense with ID %d not found", id)
}

func listExpenses(tracker *ExpenseTracker) error {
	fmt.Println("ID\tDate\t\tDescription\tAmount\tCategory")
	for _, expense := range tracker.Expenses {
		fmt.Printf("%d\t%s\t%s\t\t$%.2f\t%s\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount, expense.Category)
	}
	return nil
}

func showSummary(c *cli.Context, tracker *ExpenseTracker) error {
	var total float64
	month := c.Int("month")

	for _, expense := range tracker.Expenses {
		if month == 0 || (month > 0 && int(expense.Date.Month()) == month) {
			total += expense.Amount
		}
	}

	if month > 0 {
		fmt.Printf("Total expenses for month %d: $%.2f\n", month, total)
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}

	return nil
}

func exportToCSV(c *cli.Context, tracker *ExpenseTracker) error {
	file := c.String("file")
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("ID,Date,Description,Amount,Category\n")
	if err != nil {
		return err
	}

	for _, expense := range tracker.Expenses {
		line := fmt.Sprintf("%d,%s,%s,%.2f,%s\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount, expense.Category)
		_, err := f.WriteString(line)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Expenses exported to %s\n", file)
	return nil
}
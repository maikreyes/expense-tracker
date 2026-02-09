package cli

import (
	"expense-tracker/cmd/cli/handlers/expense"
	"expense-tracker/internal/services/expenses"
	"flag"
	"fmt"
	"os"
	"time"
)

func Run(cmd string) {
	fs := flag.NewFlagSet("expense-tracker "+cmd, flag.ContinueOnError)
	id := fs.Int("id", 0, "The ID of the expense")
	description := fs.String("description", "", "The description of the expense")
	amount := fs.Float64("amount", 0.0, "The amount of the expense")
	month := fs.Int("month", 0, "The month of the expense")

	service := expenses.NewService("data.json")
	handler := expense.NewHandler(service)

	// os.Args: [program, cmd, ...flags]
	if err := fs.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		return
	}

	switch cmd {
	case "add":
		if err := handler.AddExpense(*description, *amount); err != nil {
			fmt.Println("Error:", err)
		}
	case "list":
		if err := handler.ListExpenses(); err != nil {
			fmt.Println("Error:", err)
		}
	case "delete":
		if err := handler.DeleteExpense(*id); err != nil {
			fmt.Println("Error:", err)
		}
	case "summary":
		if *month < 0 || *month > 12 {
			fmt.Println("Error: month must be between 0 and 12")
			return
		}

		if err := handler.GetSummary(time.Month(*month)); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		println("Unknown command")
	}

}

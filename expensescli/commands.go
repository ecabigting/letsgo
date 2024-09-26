package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add  string
	Del  int
	Edit string
	Show bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new expense with format: <description>:<amount>")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an existing expense with format: <#id>:<description>:<amount>")
	flag.IntVar(&cf.Del, "del", -1, "Delete an existing espense with #id")
	flag.BoolVar(&cf.Show, "show", false, "Show All expense")
	flag.Parse()
	return &cf
}

func (cf *CommandFlags) Execute(expenses *Expenses) {
	switch {
	case cf.Show:
		expenses.show()
	case cf.Add != "":
		addingExpense := strings.SplitN(cf.Add, ":", 2)
		if len(addingExpense) != 2 {
			fmt.Println("Invalid add arguments, use <description>:<amount>")
			os.Exit(1)
		}
		amount, err := strconv.ParseFloat(addingExpense[1], 64)
		if err != nil {
			fmt.Println("Invalid amount: accepts decimal 0.00")
			os.Exit(1)
		}
		expenses.add(addingExpense[0], amount)
		expenses.show()
	case cf.Edit != "":
		editExpense := strings.SplitN(cf.Edit, ":", 3)
		if len(editExpense) != 3 {
			fmt.Println("Invalid edit arguments, <expense #id>:<description>:<amount>")
			os.Exit(1)
		}
		expenseId, err := strconv.Atoi(editExpense[0])
		if err != nil {
			fmt.Println("Invalid expense #id")
		}
		amount, err := strconv.ParseFloat(editExpense[2], 64)
		if err != nil {
			fmt.Println("Invalid edit amount: accepts decimal 0.00")
			os.Exit(1)
		}
		expenses.edit(expenseId, editExpense[1], amount)
		expenses.show()
	case cf.Del != -1:
		expenses.delete(cf.Del)
		expenses.show()
	default:
		fmt.Println("Invalid Commands")
	}
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Expense struct {
	Description string
	Amount      float64
	AddedAt     time.Time
}

type Expenses []Expense

func (expenses *Expenses) add(desc string, amount float64) {
	expense := Expense{
		Description: desc,
		AddedAt:     time.Now(),
		Amount:      float64(amount),
	}
	*expenses = append(*expenses, expense)
}

func (expenses *Expenses) validateExpenseId(id int) error {
	if id < 0 || id > len(*expenses) {
		err := errors.New("Invalid Expense ID")
		fmt.Println(err)
		return err
	}

	return nil
}

func (expenses *Expenses) delete(id int) error {
	e := *expenses
	if err := e.validateExpenseId(id); err != nil {
		return err
	}
	*expenses = append(e[:id], e[id+1:]...)
	return nil
}

func (expenses *Expenses) edit(id int, desc string, amount float64) error {
	e := *expenses
	if err := e.validateExpenseId(id); err != nil {
		return err
	}
	e[id].Amount = amount
	e[id].Description = desc

	return nil
}

func (expenses *Expenses) show() {
	table := table.New(os.Stdout)
	// table.SetRowLines(false)
	table.SetHeaders("#", "Description", "Amount")
	total := 0.00

	for id, e := range *expenses {
		total = total + e.Amount
		table.AddRow(strconv.Itoa(id), e.Description, fmt.Sprintf("$ %.2f", e.Amount))
	}
	table.AddRow("Total", "", fmt.Sprintf("$ %.2f", total))
	table.Render()
}

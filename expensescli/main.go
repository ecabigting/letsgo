package main

func main() {
	expenses := Expenses{}
	store := NewStorage[Expenses]("expenses.json")
	store.Load(&expenses)
	// Populate initial list
	if len(expenses) <= 0 {
		expenses.add("Lunch", 14.3)
		expenses.add("Train Ride", 2.44)
		expenses.add("Rent", 1.14)
		expenses.add("New Mouse", 5.14)
		expenses.add("New Data Cable", 3.14)
		expenses.add("Groceries", 1121.14)
	}
	commands := NewCommandFlags()
	commands.Execute(&expenses)
	store.Save(expenses)
}

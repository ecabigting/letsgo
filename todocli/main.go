package main

func main() {
	todos := Todos{}
	// Load the file from the filesystem
	// with filename todos.json
	storage := NewStorage[Todos]("todos.json")
	// Load the file data and
	// reference it into our todos
	// declared at line 5
	storage.Load(&todos)
	if len(todos) <= 0 {
		todos.InitTodoFile()
	}
	commands := NewCommandFlags()
	commands.Execute(&todos)
	// todos.print()
	storage.Save(todos)
}

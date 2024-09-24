package main

func main() {
	todos := Todos{}
	todos.add("Create go todo cli")
	todos.add("Create go web api")
	todos.add("Create go http server")
	todos.toggle(1)
	todos.print()
}

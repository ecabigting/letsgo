package main

func main() {
	todos := Todos{}
	todos.add("Create go todo cli")
	todos.add("Create go web api")
	todos.add("Create go http server")
	todos.toggle(0)
	todos.toggle(1)
	todos.toggle(2)
	todos.print()
}

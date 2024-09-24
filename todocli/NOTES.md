## running the app
The command `go run ./` runs all the `.go` inside the root folder

On the contrary, if you run `go run main.go` it will throw an error looking for the other file reference specifically the line where `todos := Todos{}` as it is looking for the reference.

To fix this issue, we need to organize the file structure, we need to move the `todos.go` file into a folder, and update the `package main` into `package todos` and then in the `main.go` where we have the line `todos := Todos{}` we need to reference by adding into `import "todos/todos"`

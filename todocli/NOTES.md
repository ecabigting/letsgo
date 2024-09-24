## running the app
The command `go run ./` runs all the `.go` inside the root folder

On the contrary, if you run `go run main.go` it will throw an error looking for the other file reference specifically the line where `todos := Todos{}` as it is looking for the reference.

To fix this issue, we need to organize the file structure, we need to move the `todos.go` file into a folder, and update the `package main` into `package todos` and then in the `main.go` where we have the line `todos := Todos{}` we need to reference by adding into `import "todos/todos"`

## Method Receiver/Receiver Functions
you can attached a function to a type using the `receiver` parenthesis in a function signature.

Let say you have a `type` Todo like this:
```go
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
```

Then you create a slice like this `type Todos []Todo`

You can add function to your `Todos` type by adding another parenthesis after the `func` keyword before the function name in the function signature, like this:
```go
func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}
```
From the function above the `(todos *Todos)` at the function signature denotes that this function is now a property/member of type `Todos`
> Good to know: the `*` in the `*Todos` in the function receiver means you are modifying the `todos` that called the function. Without the `*` you are modifying a copy of the todos sent to the function.


## understanding pointers

`*` is used to dereference a pointer (access the value at the pointer's address).

As an example:

```go
var x int = 42
var p *int = &x // p is a pointer to x
fmt.Println("Value pointed to by p:", *p) // Output: Value pointed to by p: 42
```

`&` is used to get the address of a variable (create a pointer).

As an example:

```go
var x int = 42
fmt.Println("Address of x:", &x)     // Output: Address of x: <some memory address>
```

### Key Points to remember

1. Pointer Assignment:
In the `toggle` function, when you do(see [todo.go](/todo.go)):
```go
completionTime := time.Now()
t[index].CompletedAt = &completionTime
```
You are creating a variable `completionTime` that holds the current time. You then take the address of this variable using `&completionTime` and assign it to `t[index].CompletedAt`, which is a pointer to a time.Time.

2. Dereferencing in Print:
In your `print` function, you check if `t.CompletedAt` is not nil:
```go
if t.CompletedAt != nil {
    completedAt = t.CompletedAt.Format(time.RFC1123)
}
```
Here, `t.CompletedAt` is a pointer to a `time.Time`. When you call `t.CompletedAt.Format(time.RFC1123)`, Go automatically dereferences the pointer for you to access the underlying `time.Time` value. This is why you see the correct formatted date and time instead of a memory address.

## Running the app
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


## Understanding pointers

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

> GOOD TO KNOW: any declaration of `struct`, `function` or `variable` names in upper case characters makes them public while lower case makes them private. This means these are accessible from other packages in the same scope.

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

## Working with Generics in Go

Take the snippet below as an example:
```go
type Storage[T any] struct {
	FileName string
}
```
- Generics: The `Storage` type is defined as a generic type with a type parameter `T`. The `any` constraint means that `T` can be any type ie. JSON, `Todos`.
- Field: The `Storage` struct has a single field, `FileName`, which is a string that presumably represents the name of the file where data will be stored.

The constructor function:
```go
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}
```
- Constructor: The `NewStorage` function is a constructor for creating a new instance of `Storage`. It takes a `fileName` as a parameter and returns a pointer to a `Storage[T]` instance.
- Type Parameter: The type parameter `T` is again specified, allowing the caller to define what type of data will be stored in this instance of `Storage`.

**Using the Generics on a Save function:**
```go
func (s *Storage[T]) Save(data T) error {
	// Using MarshalIndent saves the data into json format with 4 spaces
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Write the File
	// 0644 - owner can read write to the file, user group members can read, everyone else can read
	return os.WriteFile(s.FileName, fileData, 0644)
}
```
- Method Receiver: The `Save` method is defined on a pointer receiver of type `*Storage[T]`, meaning it can modify the `Storage` instance it is called on.
- Parameter: The method takes a parameter `data` of type `T`, which is the data to be saved.
- JSON Marshalling: The method uses `json.MarshalIndent` to convert the `data` into a JSON format with indentation for readability. If there is an error during this process, it returns the error.
> File Writing: The method then writes the JSON data to the file specified by `FileName` using `os.WriteFile`. The permission mode 0644 means:
 - The owner can read and write to the file.
 - Group members can read the file.
 - Others can also read the file.

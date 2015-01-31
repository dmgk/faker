package faker

import "fmt"

type App struct{}

// Example:
//  App{}.Name() // Alphazap
func (a App) Name() string {
	return Fetch("app.name")
}

// Example:
//  App{}.Version() // 2.6.0
func (a App) Version() string {
	return Numerify(Fetch("app.version"))
}

// Example:
//  App{}.Author() // Dorian Shields
func (a App) Author() string {
	return Fetch("app.author")
}

// Example:
//  fmt.Println(App{}) // Tempsoft 4.51
func (a App) String() string {
	return fmt.Sprintf("%v %v", a.Name(), a.Version())
}

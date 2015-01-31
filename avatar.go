package faker

import "fmt"

type Avatar struct{}

// Url returns robohash.org avatar URL with "format" image format and "w" x "h" px size.
// Example:
//  Avatar{}.Url("jpg", 100, 200) // http://robohash.org/NX34rZw7s0VFzgWY.jpg?size=100x200
func (a Avatar) Url(format string, w, h int) string {
	return fmt.Sprintf("http://robohash.org/%s.%s?size=%dx%d", Lorem{}.Characters(16), format, w, h)
}

// String returns URL of 300x300 px avatar.
// Example:
//  fmt.Println(Avatar{}) // http://robohash.org/XRWjFigoImqdeDuA.png?size=300x300
func (a Avatar) String() string {
	return a.Url("png", 300, 300)
}

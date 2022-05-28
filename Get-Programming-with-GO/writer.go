package main

import (
	"fmt"
	"io"
	"os"
)
/* In January 2015, a marvelous article on error handling was published on the Go blog (blog.golang.org/errors-are-values). The article describes a simple way to write to a file without repeating the same error-handling code after every line. To apply this technique, you need to declare a new type, which we call safeWriter in listing 28.5. If an error occurs while safeWriter is writing to a file, it stores the error instead of returning it. Subsequent attempts to write to the same file will be skipped if writeln sees that an error occurred previously. */
type safeWriter struct {
	w 	io.Writer
	err error // A place to store the first error
}
func (sw * safeWriter) writeln(s string) {
	if sw.err != nil {
		return // Skips the write if an error occurred previously
	}
	_, sw.err = fmt.Println(sw.w, s) // Writes a line and store any error
}

func main() {

}
// Using safeWriter, the following listing writes several lines without repetitive error handling, yet still returns any errors that occur.
func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
	return err
}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don’t just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don’t communicate by sharing memory, share memory by communicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")
	return sw.err
}
/* This is a far cleaner way to write a text file, but that isn’t the point. The same technique
can be applied to creating zip files or to completely different tasks, and the big idea is
even greater than a single technique: */
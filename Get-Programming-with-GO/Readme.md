Go is suitable for programmers with a wide range of skill levels—a necessity for any large project. Being a relatively small language, with minimal syntax and few conceptual hurdles, Go could be the next great language for beginners.
Unfortunately, many resources for learning Go presume a working knowledge of the C programming language. Get Programming with Go exists to fill the gap for scripters, hobbyists, and newcomers looking for a direct path to Go. To make it easier to get started, every code listing and exercise in this book can run inside the Go Playground (play.golang.org), so there’s nothing to install! If you’ve ever used a scripting language like JavaScript, Lua, PHP, Perl, Python, or Ruby, you’re ready to learn Go. If you’ve used Scratch or Excel formulas, or written HTML, you’re not alone in choosing Go as your first “real” programming language (see the video “A Beginner’s Mind” featuring Audrey Lim at youtu.be/fZh8uCInEfw). Mastering Go will take patience and effort, but we hope Get Programming with Go is a helpful
resource in your quest.

Get Programming with Go gradually explains the concepts needed to use Go effectively and provides a plethora of exercises to hone your skills. This is a beginner’s guide to Go, intended to be read from cover to cover, with each lesson building on the last. It isn’t a complete specification (golang.org/ref/spec) of every language feature, but it covers most of the language and touches on advanced topics like object-oriented design and concurrency. 
Whether you go on to write massively concurrent web services or small scripts and simple tools, this book will help you establish a solid foundation.

 Unit 1 brings together variables, loops, and branches to build tiny apps, from greet-
ings to rocket launches.
 Unit 2 explores types for both text and numbers. Decode secret messages with
ROT13, investigate the destruction of the Arianne 5 rocket, and use big numbers
to calculate how long light takes to reach Andromeda.
 Unit 3 uses functions and methods to build a fictional weather station on Mars with
sensor readouts and temperature conversions.
 Unit 4 demonstrates how to use arrays and maps while terraforming the solar sys-
tem, tallying up temperatures, and simulating Conway’s Game of Life.
 Unit 5 introduces concepts from object-oriented languages in a distinctly non-
object-oriented language. Use structures and methods to navigate the surface of
Mars, satisfy interfaces to improve output, and embed structures in one another to
create even bigger structures!
 Unit 6 digs into the nitty-gritty. Here, you use pointers to enable mutation, over-
come the knights who say nil, and learn how to handle errors without panicking.
 Unit 7 introduces Go’s concurrency primitives, enabling communication between
thousands of running tasks while constructing assembly lines in a gopher factory.
 The appendix provides our solutions for the exercises, but coming up with your
own solutions is what makes programming fun!

About the code
All code is in a fixed-width font to separate it from ordinary text. Code annotations
accompany many of the listings, highlighting important concepts.
You can download the source code for all listings from the Manning website
(www.manning.com/books/get-programming-with-go). The download also includes
solutions for all the exercises in this book. If you prefer to browse the source code
online, you can find it in the GitHub repository at github.com/nathany/get-
programming-with-go.
Although you could copy and paste code from GitHub, we encourage you to type in the
examples yourself. You’ll get more out of the book by typing the examples, fixing typos,
and experimenting with the code.

Book forum
The purchase of Get Programming with Go includes free access to a private web forum run by Manning Publications where you can make comments about the book, ask technical questions, share your solutions to exercises, and receive help from the authors and from other users. To access the forum and subscribe to it, point your web browser to forums.manning.com/forums/get-programming-with-go. You can learn more about Manning’s forums and the rules of conduct at forums.manning.com/ forums/about. Manning’s commitment to our readers is to provide a venue where a meaningful dialogue between individual readers and between readers and the authors can take place. It’s not a commitment to any specific amount of participation on the part of the authors, whose contribution to the forum remains voluntary (and unpaid). We suggest you try asking the authors some challenging questions lest their interest stray! The forum and the archives of previous discussions will be accessible from the publisher’s website as long as the book is in print.


Note: You probably wondering why there's a red line when you open folder Get-Programming-with-Go because Go only can call one main functions inside Go the right way to create functions is to create another file but only contain functions so you can declare in main function but don't worry you can open the folder outside the Get-Programming-with-Go, some file intentionally to error so you know what when wrong.
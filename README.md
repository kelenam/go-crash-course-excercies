# Golang Crash Course
These exercies were following along with Traversy Medias golang course.

Notes following the course can be found here: https://kelena-mori.netlify.com/traversy-golang-crash-course

## Initial Takeaways
Go feels really familiar and C-like. It's incredibly how easy it is setting up a web server. I don't know how I feel about the opinated folder structuring, especially built around version control platforms and usernames. I could imagine where this could be useful, but I'd need further exposure to it in a team environment to appreciate it.

## Key Insights
- Didn't know the rune type existed
- The syntax for declaring variables `:=` inside functions is interesting, one of the things that made it less familiar. Once you know it though, it becomes pretty comfortable. It's interesting though that it seems `:=` is more commonly used than using `const` everywhere. Would be interesting to dive deeper into scoping rules
- format pointers, or whatever they are called, are familiar from CS50 and C. An area I need to gain more familiarity with.
- importing feels very familiar, you can see how the strictness related to folder structure can start coming in handy with imports
- Super cool/easy to create binaries/exe files that get added to the `bin` folder. So awesome. 
- Like the option for slices. JS really spoils you with the flexibility of arrays/objects.
- I actually kind of like using parens on if/else/if conditions like JS does, I think it makes it more readable, but that's just me.
- Interesting to see loops pretty much the same, but also if we're not ever referencing the incrementor, or `i` var in the loop we'll get an error, and to use `_` instead. Kind of shows how JS hand-holds a lot of things. Don't have to worry about memory management.
- The FizzBuzz challenge came pretty naturally, I guess I never though to use traversy's approach to use 15 instead of and operators for 5 and 3.
- I definitely need to get a better handle on using maps (objects), and drilling into docs about the `make` keyword
- I like ranges, need to get more practice with them. 
- From Obj-C, and C learning days, pointers here were addressed decently enough in the need to know only kind of sense (really topically), but I liked Traversy's take on them.
- Interesting to learn everything in Go is passed by value. 
- Structs feel familiar from Rust.
- Closures work, from what I can tell from this course, very similarly to JS. I don't know under the hood what they'll be doing but probably aren't making use of `[[scope]]` and hidden passed properties on the returned function.
- Incredible how easy it is setting up a web server with Go. 


## The Compilation Process  

Computers need machine code, they don't understand English or even Go. We need to convert our high-level (Go) code into machine language, which is really just a set of instructions that some specific hardware can understand. In your case, your CPU.

The Go compiler's job is to take Go code and produce machine code, an .exe file on Windows or a standard executable on Mac/Linux.

![Go Compilation Process](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/rfR5MNc.png)

## Two Kinds of Bugs

Generally speaking, there are two kinds of errors in programming:

### Compilation errors

Occur when code is compiled. It's generally better to have compilation errors because they'll never accidentally make it into production. You can't ship a program with a compiler error because the resulting executable won't even be created.

### Runtime errors  

Occur when a program is running. These are generally worse because they can cause your program to crash or behave unexpectedly.

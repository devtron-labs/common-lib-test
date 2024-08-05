package commonlib

func Greet(name string, formal bool) string {
    if formal {
        return "Hello, " + name + "."
    }
    return "Hi, " + name + "!"
}

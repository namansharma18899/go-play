# Arrays
## Caveats
- What’s more, you can’t use a type conversion to directly convert arrays of different sizes to identical
    types. Because you can’t convert arrays of different sizes into each other, you can’t write a function that
    works with arrays of any size and you can’t assign arrays of different sizes to the same variable.
- The main reason arrays exist in Go is to provide the backing store for slices, which are one of the most 
    useful features of Go. 

# Slices
## Caveats
- Go is a call-by-value language. Every time you pass a parameter to a function, Go makes a copy of the value
    that’s passed in. Passing a slice to the append function actually passes a copy of the slice to the function.
    The function adds the values to the copy of the slice and returns the copy. You then assign the returned slice
    back to the variable in the calling function.

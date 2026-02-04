so basically what i have understood till now is 

### How GO actually work and its folder structure 

1. **bin**: 
   - the name itself it has all the executable files in binary form which will be understood by kernal
2. **pkg**:
   - it keeps the compiled packages of the go files 
3. **src**:
   - this is the directory where developers write all the codes 

### How to run GO files

```bash
go run <filename>(ex- main.go)
```

### GO files

- all Go files in a single directory are considered part of the same package

### work and importance of `go.mod`

- it is similar to package.json from Nodejs
- it keeps all the imports dependencies like 
  - package name versions
  - go version 

### All about `main` function 

- main function is executable itself we dont have to call it explicitly
- and main function will only work if it is used in the **main**(main package defines that it is the part of go directory) package
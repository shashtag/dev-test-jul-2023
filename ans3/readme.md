# Answer 3

## Running of code /installation3

Navigate into the directory ans3 of the cloned repository.

Run

```bash
go run main.go
```

## Understanding the question

The given function needs to be converted from recursive to a non recursive (iterative) version

## Approach and complexity analysis

### Base case

- If cur is not provided set cur to 0
  - Time Complexity - O ( 1 )
  - Space Complexity - O ( 1 )
- If n is < 2 throw error (panic) - also handle the error
  - Time Complexity - O ( 1 )
  - Space Complexity - O ( 1 )

### Recursive case

- For loop from n to 3
  - Time Complexity - O ( n )
  - Space Complexity - O ( 1 )

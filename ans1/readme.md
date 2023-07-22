# Answer 1

## Running of code /installation

Navigate into the directory ans1 of the cloned repository.

Run

```bash
go run main.go
```

## Understanding the question

There are 2 API to be developed.

### Input function

- Takes in a string with multiple words
- Words are seperated by a deliminator
- Need to store the stings.

### Output function

- Takes in a string with multiple words
- Words are seperated by a deliminator
- Need to print out percentage of stings having this input to the function as prefix

### Point of clarity

For inputs

```{r, tidy=FALSE, eval=FALSE, highlight=FALSE }
ingest('leoluca:uk:dev')
ingest('leoluca:hk:design')
ingest('leoluca:hk:pm')
ingest('leoluca:hk:dev')
ingest('skymaker')
```

What will be the output for

```{r, tidy=FALSE, eval=FALSE, highlight=FALSE }
appearance('leo')
```

- Will it be 0.8 (as leo is the prefix of of 4 out of 5 strings)
- Will it be 0 (as leo is not a complete word which ends with a delimiter)

I have written the answer with the understanding that it should be 0 or else there will not be any use of the delimiter

## Approach and complexity analysis

n - number of letters  
w - number of words

### Data Structures

- structure{trie root node and ingests} (trie) - enters all the stings and stores them in form of a tree. Keeps the count of number of ingests
- map (dict) - stores the occurrence of a word.

### Input func

> Time complexity - O ( w + n ), Space complexity - O ( w + n )

- Increses ingests by 1
  - Time complexity - O ( 1 )
- Records all the words in the map
  - Time complexity - O ( w )
  - Space complexity - O ( w )
- Adds all the letter to the trie - O ( n )
  - Time complexity - O ( n )
  - Space complexity - O ( n )

### Output func

> Time complexity - O ( 1 ) or O ( w ) or O ( w + n ) - depending on the input  
> Worst case Time complexity - O ( w + n ), Space complexity - O ( 1 )

- Checks in for the **edge case** input to the function is an empty string
  - Time complexity - O ( 1 )
- Check if all the words are in the map
  - Time complexity - O ( w )
- If yes then checks the trie - O ( n )
  - Time complexity - O ( n )

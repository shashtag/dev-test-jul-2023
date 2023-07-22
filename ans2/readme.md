# Answer 2

## Running of code /installation

Navigate into the directory ans2 of the cloned repository.

Run

```bash
go run main.go
```

## Understanding the question

Have to develop a cache of fixed capacity. The cache has two functions and a formula to calculate which item is to removed when cache is full.

### Input function

- Takes in key, value and weight and saves it into the cache

### Output function

- Takes in key and returns the value for the key

## Approach and complexity analysis

> n - capacity of cache

### Data Structures

- Structure{capacity,map of key to cache data} (cache) - hashtable with a capacity.
- Heap - min heap to get the key with lowest value.

### Input func

> Time complexity - O ( 1 ) or O ( n ) - depending upon if capacity is reached  
> Worst case Time complexity O ( n ), Space complexity - O ( 1 )

- If key exists, update the value and weight
  - Time complexity - O ( 1 )
- If cache is full, pop the least score item
  - Time complexity - O ( n ) - O ( log n ) for heapify and O ( 1 ) for calculating score
- Add the new item into heap list and cache map
  - Time complexity - O ( 1 )
  - Space complexity - O ( 1 )

### Output func

> Time complexity - O ( 1 ), Space complexity - O ( 1 )

- If key exists, return the value else return -1
  - Time complexity - O ( 1 )

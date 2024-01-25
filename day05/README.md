## Ex00
Implemented the `areToysBalanced` function, which checks if a binary tree's left 
and right subtrees have an equal number of toys (represented by boolean values). 
The function returns true if balanced and false otherwise.

## Ex01
Created the `unrollGarland` function to traverse a binary tree layer by layer, generating 
a slice of booleans based on a zig-zag pattern. The result represents the presence of a 
garland on each layer.

## Ex02
Introduced the Present struct and implemented a heap-based data structure called 
`PresentHeap` using the "container/heap" package. The `getNCoolestPresents` function 
utilizes this heap to return a sorted slice of the coolest presents based on value 
and size.

## Ex03
Implemented the classic dynamic programming algorithm, known as the "Knapsack Problem," 
in the `grabPresents` function. This function takes a slice of presents with values and 
sizes, along with the capacity of a hard drive. It returns a slice of presents that 
maximizes the cumulative value within the given capacity.
## Ex00
Reviewed and fixed the existing `minCoins` function, which had a bug in handling duplicate denominations, 
non-sorted input slices and an empty list. Implemented test cases in test/ex00/minCoins_test.go files to 
demonstrate the incorrect behavior of the original code and created a new function,`minCoins2`, to handle 
error cases correctly.

## Ex01
Utilized Go's built-in tools to profile the execution of the minCoins2 function and identify the top 10 
functions consuming the most CPU time. Additionally, created benchmark tests to compare the performance 
of the minCoins2 function.

## Ex02
Enhanced code comments to describe the differences between different versions of the minCoins function. 
Utilized a documentation tool to generate HTML documentation based on the comments. Submitted compressed 
HTML documentation (docs.zip) for accessibility.
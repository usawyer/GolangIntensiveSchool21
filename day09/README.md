## Ex00
Implemented `sleepSort` function that takes an unsorted slice of integers and returns an integer channel. 
The function uses goroutines to sleep for a duration equal to each number and then sends them to the 
output channel.

## Ex01
Developed `crawlWeb` function for web crawling. The function accepts an input channel for URLs. Added 
context to facilitate graceful shutdown, allowing the program to stop after processing all given URLs.

## Ex02
Implemented `multiplex` function that accepts a variable number of channels (chan interface{}) as 
arguments and returns a single output channel of the same type. The function performs the `fan-in 
pattern` by redirecting messages from input channels to the output channel. A test demonstrates 
that values sent randomly to any input channel are received on the output channel.
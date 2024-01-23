## Ex00
A `find`-like utility that locates directories, regular files, and symbolic links 
based on user-specified command-line options. Options include printing specific 
types (-f, -d, -sl), filtering by extension (-ext), and resolving and handling broken 
symbolic links. Handle permission errors gracefully and allow users to specify one or 
multiple types in a single command.

## Ex01
A`wc`-like utility to gather basic statistics about utf-8 encoded text files, 
considering both English and Russian languages. Implement three mutually exclusive 
flags (-l, -m, -w) to count lines, characters, and words, respectively. Utilize 
goroutines for concurrent processing of independent files. If no flags are specified, 
default to word count.

## Ex02
A utility similar to `xargs` that treats all parameters as a command and builds and 
executes the command by appending lines from stdin as arguments. Allow users to chain 
this tool with the previous exercises, enabling scenarios like counting line numbers 
for specific log files recursively.

## Ex03
A log rotation tool in Go that archives old log files based on timestamps. Create 
tar.gz files for individual logs or multiple logs and store them in a specified archive 
directory.

## Ex00
Create an interface DBReader and two implementations, one for reading JSON
and one for reading XML. Both implementations should return the same object
type. Use the DBReader interface in a CLI application that reads databases 
from files, distinguishing formats by file extension.

## Ex01
Develop a CLI application that compares two databases (in JSON and XML formats) 
for modifications. Identify cases such as added or removed cakes, changed cooking 
times, added or removed ingredients, and modified unit details. Ensure the 
application reuses the code from Exercise 00.

## Ex02
Extend the application to handle filesystem backups. Create a CLI tool that 
compares two plain text files containing file paths. Output the added and removed
files. Consider memory constraints and propose solutions, excluding the "CHANGED" case.

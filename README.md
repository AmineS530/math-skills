# math-skills

## Overview
This program calculates the following statistics from data in a file:
- **Average**
- **Median**
- **Variance**
- **Standard Deviation**

## Input
The input file should contain one numerical value per line, as the following example:

```
189
113
121
114
145
110
...
```

## Usage
#### Run the program from the command line, while passing the input file as an argument or use the Makefile associated.

#### Example:
```zsh
go run main.go <file_name.txt>
```
#### or
```zsh
./[program_name] <file_name.txt>
```

## How to test
#### Run the premade Makefile using:
```zsh
make
```

## Output
#### The program will output the calculated statistics in this format:

```
Average: xx
Median: xx
Variance: xx
Standard Deviation: xx
```
#### or this when using the Makefile
```
Results from go app
Average: w
Median: x
Variance: y
Standard Deviation: z

Solution
Average: w
Median: x
Variance: y
Standard Deviation: z
```

##### <p align="center">Made by asadik</p>
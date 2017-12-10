# What is this?

There are times when I want to insert a series of tasks into
[Taskwarror](https://taskwarrior.org) that look very similar.

For example:

    - task add Read chapter 1 project:"Book.My Epic Book"
    - task add Read chapter 2 project:"Book.My Epic Book"
    - task add Read chapter 3 project:"Book.My Epic Book"
    - task add Read chapter 4 project:"Book.My Epic Book" 
    ...

but I don't want to insert all of these tasks one by one.

This is the problem that task template solves.

# How to install

## Prerequisites

Make sure that you have [Golang](https://golang.org) and
[Taskwarror](https://taskwarrior.org) installed.

## Installation steps

1. Clone the repository to your $GOPATH

    `git clone https://github.com/sirrah23/TaskTemplate.git`

2. Run `go install Tasktemplate/tasktemplate` at the root of your $GOPATH

3. Retrieve the binary from the bin directory

# How to use it

After you have installed TaskTemplate you can use it by running,

    tasktemplate -template="Read chapter #" -start=1 -end=4
    -project="Book.My Epic Book"

which will insert the following tasks into TaskWarrior:

    Read Chapter 1 project:"Book.My Epic Book"
    Read Chapter 2 project:"Book.My Epic Book"
    Read Chapter 3 project:"Book.My Epic Book"
    Read Chapter 4 project:"Book.My Epic Book"

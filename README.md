# What is this?

There are times where I want to insert a series of tasks into
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

TODO

# How to use

After you have installed TaskTemplate you can use it by running,

    tasktemplate -template="Read chapter #" -start=1 -end=4
    -project="Book.My Epic Book"

which will insert the following tasks into TaskWarrior:

    Read Chapter 1 project:"Book.My Epic Book"
    Read Chapter 2 project:"Book.My Epic Book"
    Read Chapter 3 project:"Book.My Epic Book"
    Read Chapter 4 project:"Book.My Epic Book"

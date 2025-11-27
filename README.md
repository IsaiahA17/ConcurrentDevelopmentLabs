# Concurrent Development Labs
### Contact Author: C00286361@setu.ie
This repository contains the code for solutions to various concurrency problems written in GO for the Concurrent Development module. The concurrency problems tackled in this repository are:
- The Barrier Problem, in which goroutines must be synchronised in order to print certain messages in a separate order, as without a barrier, both messages may be printed at the same time instead of separately
- The Reusable Barrier Problem, which is a continuation of the previous problem, where a loop is now used, and a reusable barrier must be implemented as the loop may cause a single barrier to be ineffective at separating the messages printed.
- The Dining Philosophers Problem, which is a problem where, in this case, 5 goroutines representing philosophers must be able to eat and think, however, they require "forks" on their left and right to eat, causing a circular wait as at least more than one philosopher must be eating.
- Producer Consumer, which is a problem where multiple threads may try to write and read from a buffer, causing issues if multiple threads were to try and write to the buffer or read from the buffer at the same time.
- Concurrent Essentials is a folder containing examples of mutexes, atomic integers, semaphores, a problem involving a semaphore, and a semaphore implementation using channels in go.
- Game Of Life is an implementation of Conway's Game Of Life, that was looked over at a lab and the drawing code was used and modified for my WaTor project.

# Introduction
Barrier2.go is a reusable barrier implementation through the use of an atomic int and an unbuffered channel. Since having only one barrier would not work in a loop as it may be possible for a goroutine to move on to print Part A while other goroutines may still be printing part B.

# How It Works
In order for the barrier to be reusable, another barrier has to be added to the beginning to prevent the goroutine from proceeding to printing part A while others are still printing Part B. The goroutines will increment an atomic int and get blocked by trying to receive from a channel. The atomic integer keeps track of the number of goroutines so that the last one can send a bool variable to the others, signalling them to continue. Once the others continue, the atomic integer will decrement to keep track of how many have passed through, and the same pattern takes place as the first barrier to print part A and part B separately, even in loops.

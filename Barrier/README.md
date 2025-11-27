# Introduction
barrier.go is an example of a barrier through the use of a buffered channel and an atomic integer. All goroutines are supposed to print part A concurrently and then print Part B.

# How It Works
Without the barrier implementation, the program may print both part A and part B at the same time. Through the use of an empty buffered channel to block goroutines from heading to the section to print part B in the doStuff function until all goroutines print part A. The last goroutine will place an empty struct into the buffered channel to signal all the goroutines to move on and print part B, which will show each part printed seperately. The last goroutine will wait at the second last channel by trying to take from it and the others will signal it to move by placing an empty struct in the channel

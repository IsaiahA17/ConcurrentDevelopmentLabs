# Introduction
This is an attempt at a C++ implementation for a solution to the producer consumer problem in which several threads may try to write to and take from a buffer. Problems with this arise however as multiple threads writing to the same index or memory location at the same time may cause issues and similar issues may occur with multiple threads taking from the same memory location at the same time.

# How It Works
The SafeBuffer file should have a mutex when both writing to and getting information from the queue to prevent the problems mentioned above. However it's also necessary to implement a semaphore to block all threads in the event the threads try to read from an empty buffer or writing to a supposedly full buffer, causing more problems
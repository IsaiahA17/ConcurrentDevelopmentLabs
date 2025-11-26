# Introduction
dinPhil.go is a solution to the dining philosophers problem, which involves 5 goroutines (representing the philosophers) and a map of boolean channels that can be accessed by an integer index (representing the forks required to eat). Each philosopher must have access to two forks, on their left and right (represented by the forks[index] and forks[index+1] in code) in order to eat, otherwise, they're thinking. Only one philosopher may hold a fork at a time, which creates a problem since it's required that at least more than one philosopher should be able to eat at a time. 

# How It Works
This is solved by an if statement in which, if the index is 0, meaning if the first philosopher is trying to eat, it grabs the fork to the right and then left, rather than trying to grab the one on the left like all the other philosophers, which avoids a circular wait. This is a known solution to the problem.

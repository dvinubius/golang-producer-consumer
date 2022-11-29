# Producer-Consumer via Channel

Solution to a classic CS problem in go, making use of go routines, a channel and colored terminal output 🤩


## Bounded-Buffer Problem

Producer and consumer are coupled via a buffer of finite size. 

## Our solution

Buffer implemented as go channel. Producer and consumer both run as go routines.

Producer produces at regular intervals. Consumer consumes at randomized intervals.

Producer's buffer (remaining capacity) is displayed at each production step.

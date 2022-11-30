# Producer-Consumer via Channel

Solution to a classic CS problem in go, making use of go routines, a channel and colored terminal output 🤩

## Bounded-Buffer Problem

Producer and consumer are coupled via a buffer of finite size. 

## Our solution

![Screenshot 2022-11-30 at 18 43 41](https://user-images.githubusercontent.com/32189942/204857244-935b9020-65f2-473a-9459-742d5ab613d6.png)

Buffer implemented as go channel. Producer and consumer both run as go routines.

Producer produces at regular intervals. Consumer consumes at randomized intervals.

Producer's buffer (remaining capacity) is displayed at each production step.

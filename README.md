# Fibers

## What are they

A Fiber is a very lightweight threaded task that can not only be created and executed, but
also can be yielded. Once a fiber is executing it can finish or it can launch one or more 
fibers and wait for their completion. This makes it very easy to reach 100% CPU utilization 
in a game, also you do not immediately have to pull apart your deep execution style systems. 
And you still have the ability to care about cache coherency if you want. 

## Threads

[N] Threads are locked to a core

Each Thread is going to query the fiber queues, we will have 3 of them; LOW, NORMAL, HIGH.

When a Thread finds no fibers it will sleep until it is signaled that a fiber has been added.

If a fiber is calling `run_all_sync(fibers, 100)` the fiber, let's call it X, is swapped out 
with a new fiber. How do we figure out when to continue fiber X and how do we store this 
information? Well if we link each of the 100 fibers to a `conditional` that also contains a 
pointer to fiber X we could quickly know which fiber to resume when the conditional is true. 
Also the conditional will be true on only one of the threads, the one that picks the last fiber.

## IO (File and Network)

For doing file/network IO we have a dedicated thread that processes file/network requests. 
So the IO thread needs to be able to queue fibers.

## Example

[Game Engine | Fiber]
Construct/Setup FrameParams
Lauch Fibers:
    Input
    Update AI
    Update Components
    Update Physics
    Update Animation
    Update UI/HUD
    Process VFX 
    Process Audio 
    Process Rendering

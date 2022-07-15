# xfiber

## What is it

A Fiber is a very lightweight threaded task that can be created and executed.
This makes it very easy to reach 100% CPU utilization in a game.

## Main API

```c++
    namespace fibers
    {
        void *create(size_t stack_size, void (*proc)(void *), void *arg);
        void  destroy(void *fiber);
        void  switch_to(void *fiber);
        bool  convert_to_thread(void);
        void *convert_from_thread(void *arg);
    }
```

## Fiber API

```c++
    namespace fibers
    {
        run_async(void* fiber);
        run_sync(void* fiber);
        run_all_async(void** fibers, u32 count);
        run_all_sync(void** fibers, u32 count);
    }
```


## Threads

[N] Threads are locked to a core

Each Thread is going to query the fiber queues, we will have 3 of them; LOW, NORMAL, HIGH.

When a Thread finds no fibers it will sleep until it is signaled that a fiber has been added.

If a fiber is calling `run_all_sync(fibers, 100)` the fiber, let's call it X, is swapped out with a new fiber. 
How do we figure out when to continue fiber X and how do we store this information ?
Well if we link each of the 100 fibers to a `conditional` that also contains a pointer to fiber X we could quickly
know which fiber to resume when the conditional is true.


## IO

For doing file/network IO we have a dedicated thread that process file/network requests. So the IO system needs to be able to queue fibers.

## IO Example

```c++
    // This call will post the IO request and switch to another fiber, when the IO request
    // has finished (or failed), the processing will come back here.
    IOData mapdata;
    PostFileIORequest("main.map", IO:OPENREADCLOSE, 0, 0, &mapdata);
    
    // Process the mapdata
    ...
```

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

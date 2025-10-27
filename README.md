# Project 02 — Go_Bootcamp

**Summary:** In this project, you will get familiar with the concurrency mechanisms in Go. You will learn to use channels for synchronizing goroutines and for data exchange between them. You will also study the sync package, which provides an extended set of synchronization tools, such as WaitGroup, Mutex, Once, and RWMutex.

💡 [Click here](https://new.oprosso.net/p/4cb31ec3f47a4596bc758ea1861fb624) to leave feedback on this project. It’s anonymous and will help our team improve the quality of your education. We recommend filling out the form right after finishing the project.

## Contents

  - [Chapter I](#chapter-i)
    - [Instructions](#instructions)
  - [Chapter II](#chapter-ii)
    - [General Information](#general-information)
  - [Chapter III](#chapter-iii)
  - [Project “Concurrent Execution”](#project-concurrent-execution)
    - [Task 1. Stopwatch for Asynchronous Tasks](#task-1-stopwatch-for-asynchronous-tasks)
    - [Task 2. Square Number Generator](#task-2-square-number-generator)
    - [Task 3. Ticker](#task-3-ticker)
    - [Bonus Task 4. LRU Cache Using Generics](#bonus-task-4-lru-cache-using-generics)


## Chapter I

### Instructions

1. Throughout the course, you will often feel uncertain and face a lack of information — that’s normal. Don’t forget that the information in the repository and on Google is always at your fingertips. So are your peers and Rocket.Chat. Communicate. Search. Use common sense. Don’t be afraid to make mistakes.
2. Be critical of your sources. Verify. Think. Analyze. Compare.
3. Read the assignments carefully. Then read them again.
4. It’s also a good idea to read the examples attentively. There might be something in them that’s not explicitly stated in the task itself.
5. You might encounter contradictions where something new in a task or example contradicts what you already know. If that happens, try to figure it out. If it doesn’t work — write down your question and resolve it as you go. Don’t leave open questions unanswered.
6. If a task seems unclear or impossible — it only seems that way. Try breaking it down. Most likely, individual parts will start to make sense.
7. You’ll face a variety of tasks along the way. Bonus tasks are for the most meticulous and curious learners. These tasks are more difficult and optional, but if you complete them, you’ll gain extra experience and knowledge.
8. Don’t try to fool the system or others. Ultimately, you’ll be fooling yourself.
9. Got a question? Ask the person to your right. If that doesn’t help — the one to your left.
10. When using help — make sure you fully understand the why, how, and what for. Otherwise, the help is meaningless.
11. Always push only to the develop branch! The master branch will be ignored. Work inside the src directory.
12. Your directory must not contain any files other than those specified in the assignment.

## Chapter II

### General Information

In Go, a program can consist of multiple simultaneously running subtasks that communicate with each other via channels.

The concurrency model in Go is based on the following abstractions:

- **Goroutine** — a lightweight thread that uses much less memory than an OS thread (on average, 4 KB). The Go runtime can run multiple goroutines on a single OS thread by quickly switching between them due to their small size.
- **Channel** — the second key element of concurrency in Go. Channels form the basis for data exchange and synchronization between threads. Goroutines can send data into a channel, and other goroutines can receive that data from the channel.

The Go runtime scheduler manages all created goroutines. It creates queues, each bound to an OS thread. During execution, the next goroutine is taken from a queue and, along with its context, is dispatched to an OS thread.

The **preemptive scheduler** attempts to evenly distribute CPU time among goroutines. This creates the illusion of parallel execution even when the number of goroutines far exceeds the number of available system threads.

**Concurrency issues** arise when multiple threads access shared data or memory areas. The most common issue is a **race condition**. All arithmetic and logical operations in code are **non-atomic**. If multiple threads run on different processor cores and write data to the same memory area, a race occurs: one thread might copy an outdated version of the data into its register, another overwrites it, and so on. To handle race conditions, we recommend using primitives from the sync package, such as sync.Atomic, sync.Mutex, and sync.RWMutex.

**Tip:** Before starting, we recommend studying how the Go scheduler works.  
[A three-part article series is available.](https://habr.com/ru/articles/478168/)  
For learning channels, check out the following [article](https://habr.com/ru/articles/490336/).

## Chapter III

## Project “Concurrent Execution”

### Task 1. Stopwatch for Asynchronous Tasks

1. The program reads two launch arguments: `N` and `M`.
1. The parameters `N` and `M` are of type `int`.
1. Parameters `N` and `M` are passed as arguments when launching the program.
1. The program launches `N` goroutines, each of which sleeps (`time.Sleep`) for a random duration of up to `M` milliseconds.
1. The program waits for all goroutines to finish.
1. The program prints a list to the console consisting of pairs `<goroutine number, sleep time>`, sorted in descending order of sleep time.
1. The goroutine number is the iteration index of the loop in which the goroutine was launched.
1. The sleep time is the number of milliseconds the goroutine slept.

**The use of channels is not allowed.**

**Hint:** Use the sync package. Wait for all goroutines to complete before starting the output. Use the flag package for parsing arguments.

### Task 2. Square Number Generator

1. The program reads two arguments from the command line: `K` and `N`.
2. The parameters `K` and `N` are passed via command-line arguments.
3. The program launches two functions: a generator and a squaring function.
4. The parameters `K` and `N` are of type `int`.
5. The generator function starts a goroutine and returns channel 1. Inside the goroutine, numbers from `K` to `N` (inclusive) are generated and sent into channel 1.
6. The squaring function starts a goroutine and returns channel 2. Inside the goroutine, numbers are read from channel 1, squared, and the result is sent to channel 2.
7. The main program (`main`) reads numbers from channel 2 and prints them to the console.
8. The squaring function must accept channel 1 as a read-only channel, which is returned by the generator function.
9. Both the squaring and generator functions must run concurrently.
10. Squaring must occur sequentially. After reading a number from channel 1, it must immediately be squared and sent to the next channel, and only then should the next number be processed.

**Hint:**  
Channels must be created inside the functions, and returned with read/write restrictions applied.  
Channels must be closed once the function finishes its work.  
The generator and squaring functions must operate concurrently.

### Task 3. Ticker

1. The program reads the parameter `K` from the command-line arguments.
2. The parameter `K` is passed through the arguments when launching the program.
3. `K` defines the ticker interval in seconds and must be of type `uint`.
4. The program prints to stdout the message Tick `<i>` since `<time>`, where `<i>` is the tick number and `<time>` is the time in seconds since the ticker started.
5. The program runs until the user sends a SIGTERM or SIGINT signal.
6. Upon receiving one of these signals, the program stops the ticker and prints the message Termination.
7. The ticker must operate asynchronously. It is forbidden to use functions from the `time` package such as `time.After` or `time.Ticker`.  
    You are allowed to use constants from the `time` package and the `Sleep` function.

### Bonus Task 4. LRU Cache Using Generics

1. You need to implement a package that contains a Cache structure using the **LRU** (Least Recently Used) algorithm.
2. A constructor function is required, which accepts the element type and cache capacity, and returns a pointer to the Cache structure.
3. The `Cache` should provide the following methods:
    - `Set` — add an item with a given key;
    - `Get` — retrieve an item by key;
    - `Clear` — delete all items in the cache.
4. The time complexity for `Set` and `Get` operations must be **O(1)**.
5. **LRU logic:**
    - When adding an item:
        - If the item is already in the cache, its position moves to the front.
        - If the item is not in the cache and capacity is not exceeded, it is added to the front.
        - If the item is not in the cache and capacity is exceeded, it is added to the front, and the last (least recently used) item is removed.
    - When retrieving an item:
        - If the item exists, it is moved to the front and returned along with an additional `true` flag.
        - If the item does not exist, return the type’s zero value and the `false` flag.
6. The cache must only work with the type specified when the constructor function is called (use **Golang generics**).
7. The cache must be **thread-safe**, meaning concurrent access from different goroutines must not lead to race conditions.
8. You must implement tests for the following scenarios:
    - Rarely used items are removed from the cache.
    - If capacity is exceeded, rarely used items should be evicted.
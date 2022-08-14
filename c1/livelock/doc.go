package livelock

// -------------Livelock-------------------
// Livelocks are programs that are actively performing concurrent operations,
// but these operations do nothing to move the state of the program forward.

//----------------- type Starvation -------------------
// Livelocks are a subset of a larger set of problems called starvation
// a livelock : all the concurrent processes are starved equally, and no work is accomplished

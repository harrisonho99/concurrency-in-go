package atomic

// --------------Critical section----------------
// A section of your program that needs
// exclusive access to a shared resource.
// is called a critical section

// ----------------Atomicity-------------------
// When something is considered Atomic,
// or to have the property of Atomicity,
// means that within the context that it is operating,
// it is indivisible, or uninterruptible.

//---------------Atomic ops context--------------------
// Operations that are atomic within the context of your process may not be atomic in the context of the operating system;
// operations that are atomic within the context of the operating system may not be atomic within the context of your machine;
// and operations that are atomic within the context of your machine may not be atomic within the context of your application.

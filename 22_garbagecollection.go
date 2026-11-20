// Garbage Collection (GC) in Go is an automatic memory management mechanism. Its purpose is to identify objects in heap memory that are no longer being used by the program and free that memory automatically. This helps prevent memory leaks and eliminates the need for developers to manually free memory, unlike languages such as C.

// How does it work?
Go uses a concurrent mark-and-sweep garbage collector. It starts from root references such as global variables and stack variables, marks all reachable objects as live, and then sweeps away the objects that are no longer reachable. Since it runs concurrently with the application, pause times are very short.

// Why do we need GC?
Without Garbage Collection, developers would have to manually free memory after using it. Forgetting to do so can cause memory leaks. Go GC automatically cleans up unused objects, making programs safer and easier to develop.

// Does Go use Garbage Collection for stack memory?
A: No. Stack memory is automatically cleaned up when a function returns. GC mainly manages heap memory.

// When does GC run?
A: The Go runtime decides when to run GC automatically based on memory usage. Developers usually don't need to trigger it manually.

// Can Go still have memory leaks?
A: Yes. If the program keeps references to objects (for example, in global variables, slices, maps, or caches), those objects remain reachable, so the GC cannot free them.


// What is a Memory Leak?
A memory leak happens when memory is no longer needed, but the program still keeps a reference to it, so the Garbage Collector (GC) cannot free it.

In simple words:
Memory leak = Unused memory that cannot be released because something is still pointing to it.

Real-Life Example

Imagine you rented a hotel room.

You leave the room.
But you never return the key.

The hotel thinks you're still using the room because you still have the key.
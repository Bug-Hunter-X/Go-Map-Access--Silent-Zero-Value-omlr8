# Go Map Access: Silent Zero Value

This example demonstrates a potential pitfall in Go when working with maps: accessing a non-existent key doesn't cause an error, but instead returns the zero value for the map's value type.

This can lead to subtle bugs because you might not notice that a value you expected to be present is actually absent.

**bug.go** contains code that exhibits this behavior.  **bugSolution.go** provides a corrected version.
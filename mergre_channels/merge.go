package mergre_channels

/*
Write a function func Merge(c ...chan string) <-chan string, which takes any amount of channels and returns a new channel.
All messages from the input channels must be forwarded to the new channel. Once all input channels are closed,
also the returned channel must be closed.

The order of the forwarded messages doesn't matter, but you should consume from all incoming channels concurrenly.
*/

func Merge(c ...chan string) <-chan string {
	return c[0]
}

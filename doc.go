/*
	Package autonomous provides interfaces and structures for thinking about goroutines as independent processes,
	or agents, each performing some function independent of other agents.

	We introduce the concept of an Agent, which has a "life," and can therefore be started and stopped.

	We then introduce the concept of a Manager, which is an Agent that can start and stop and maintain references
	to other agents.
*/
package autonomous

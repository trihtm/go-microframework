package queue

type Factory interface {
	/**
	Resolve a queue connection instance
	 */
	Connection(name string) Queue
}
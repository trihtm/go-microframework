package queue

type Monitor interface {
	/**
	Register a callback to be executed on every iteration through the queue loop.
	 */
	Looping(callback interface{})

	/**
	Register a callback to be executed when a job fails after the maximum amount of retries.
	 */
	Failing(callback interface{})

	/**
	Register a callback to be executed when a daemon queue is stopping.
	 */
	Stopping(callback interface{})
}
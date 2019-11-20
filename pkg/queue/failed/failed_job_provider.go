package failed

type FailedJobProvider interface {
	/**
	Log a failed job into storage.
	 */
	Log(connection string, queue string, payload string, exception interface{})

	/**
	Get a list of all of the failed jobs.
	 */
	All()

	/**
	Get a single failed job
	 */
	Find(id interface{})

	/**
	Delete a single failed job from storage
	 */
	Forget(id interface{})

	/**
	Flush all of the failed jobs from storage
	 */
	Flush()
}
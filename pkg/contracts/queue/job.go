package queue

type Job interface {
	/**
	Get the job identifier.
	 */
	GetJobId() string

	/**
	Get the decoded body of the job.
	 */
	Payload() interface{}

	/**
	Fire the job.
	 */
	Fire()

	/**
	Release the job back into the queue.
	Accepts a delay specified in seconds.
	 */
	Release(delay int)

	/**
	Determine if the job was released back into the queue.
	 */
	IsReleased() bool

	/**
	Delete the job from the queue.
	 */
	Delete()

	/**
	Determine if the job has been deleted.
	 */
	IsDeleted() bool

	/**
	Determine if the job has been deleted or released.
	 */
	IsDeletedOrReleased() bool

	/**
	Get the number of times the job has been attempted.
	 */
	Attempts() int

	/**
	Determine if the job has been marked as a failure.
	 */
	HasFailed() bool

	/**
	Mark the job as "failed".
	 */
	MarkAsFailed()

	/**
	Delete the job, call the "failed" method, and raise the failed job event.
	 */
	Fail()

	/**
	Get the number of times to attempt a job.
	 */
	MaxTries() int

	/**
	Get the number of seconds the job can run.
	 */
	Timeout() int

	/**
	Get the timestamp indicating when the job should timeout.
	 */
	TimeoutAt() int

	/**
	Get the name of the queued job class.
	 */
	GetName() string

	/**
	Get the resolved name of the queued job class.
	Resolves the name of "wrapped" jobs such as class-based handlers.
	 */
	ResolveName() string

	/**
	Get the name of the connection the job belongs to.
	 */
	GetConnectionName() string

	/**
	Get the name of the queue the job belongs to.
	 */
	GetQueue() string

	/**
	Get the raw body string for the job.
	 */
	GetRawBody() string
}
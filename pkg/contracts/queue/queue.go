package queue

import (
	"time"
)

type Queue interface {
	/**
	Get the size of the queue.
	 */
	Size(queue string) string

	/**
	Push a new job onto the queue.
	 */
	Push(job Job, data interface{}, queue string)

	/**
	Push a new job onto the queue.
	 */
	PushOn(queue string, job Job, data interface{})

	/**
	Push a raw payload onto the queue
	 */
	PushRaw(payload string, queue string, options interface{})

	/**
	Push a new job onto the queue after a delay
	 */
	Later(delay time.Duration, job Job, queue string)

	/**
	Push a new job onto the queue after a delay
	 */
	LaterOn(queue string, delay time.Duration, job Job, data interface{})

	/**
	Push an array of jobs on to the queue
	 */
	Bulk(jobs []Job, data interface{}, queue string)

	/**
	Pop the next job off of the queue
	 */
	Pop(queue string) Job

	/**
	Get the connection name for the queue
	 */
	GetConnectionName()

	/**
	Set the connection name for the queue
	 */
	SetConnectionName()
}
package eventstore

type EventStore interface {
	Load()

	LoadFromPlayHead()

	Append()
}

package storage

type Storage interface {
	WriteEvent()
	GetEvent()

	WriteSpan()
	GetSpan()
}

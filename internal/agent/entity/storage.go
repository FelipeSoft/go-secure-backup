package entity

type Storage interface {
	PutFile(content *Content)
}

package storage

type Storage interface {
	PutFile(byte *[]byte)
}

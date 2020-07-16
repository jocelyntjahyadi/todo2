package storage

import "log"

type StorageType int

const (
	StorageTypeMemory StorageType = iota
	StorageTypeMock
)

func GetStorage(t StorageType) Storage{
	var s Storage

	switch t{
	case StorageTypeMemory:
		s = newMemory()
	case StorageTypeMock:
		s = mock{}
	default:
		log.Fatal("storage type tidak ditemukan")
	}
	return s
}
package main

import "io"

type Store interface {
    Open(string)(io.ReadWriteCloser, error)
}

type StorageType int

const (
    DiskStorage StorageType  = 1
    TempStorage = 2
    MemStorage = 3
)

func NewStore(t StorageType) Store {
    switch t {
    case DiskStorage:
        return NewDiskStorage()
    case MemStorage:
        return NewMemStorage()
    default:
        return NewTempStorage()
    }
}
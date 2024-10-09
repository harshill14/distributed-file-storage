package main

import (
    pb "distributed-file-storage/proto"
)

func (w *worker) GetChunk(id string) (*pb.Chunk, error) {
    w.mutex.RLock()
    defer w.mutex.RUnlock()
    data, exists := w.storage[id]
    if !exists {
        return nil, ErrChunkNotFound
    }
    return &pb.Chunk{Id: id, Data: data}, nil
}

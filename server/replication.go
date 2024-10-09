package main

import (
    "context"
    "log"
    "sync"
    "time"

    pb "distributed-file-storage/proto"
    "google.golang.org/grpc"
)

func (s *server) replicateChunk(chunk *pb.Chunk) {
    workers := []string{"worker1:50052", "worker2:50053"}
    var wg sync.WaitGroup
    for _, workerAddr := range workers {
        wg.Add(1)
        go func(addr string) {
            defer wg.Done()
            conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
            if err != nil {
                log.Println("Failed to connect to worker:", err)
                return
            }
            defer conn.Close()
            client := pb.NewStorageServiceClient(conn)
            ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
            defer cancel()
            _, err = client.UploadChunk(ctx, chunk)
            if err != nil {
                log.Println("Failed to replicate chunk:", err)
                return
            }
            s.mutex.Lock()
            s.chunkLocations[chunk.Id] = append(s.chunkLocations[chunk.Id], addr)
            s.mutex.Unlock()
        }(workerAddr)
    }
    wg.Wait()
}

package main

import (
    "context"
    "log"
    "net"
    "os"
    "sync"

    pb "distributed-file-storage/proto"
    "google.golang.org/grpc"
)

type worker struct {
    pb.UnimplementedStorageServiceServer
    storage map[string][]byte
    mutex   sync.RWMutex
}

func (w *worker) UploadChunk(ctx context.Context, chunk *pb.Chunk) (*pb.UploadStatus, error) {
    // Store the chunk locally
    w.mutex.Lock()
    w.storage[chunk.Id] = chunk.Data
    w.mutex.Unlock()

    return &pb.UploadStatus{Success: true, Message: "Chunk stored successfully"}, nil
}

func StartWorker() {
    // Get worker name from environment variable
    workerName := os.Getenv("WORKER_NAME")
    port := os.Getenv("PORT")
    if port == "" {
        port = "50052"
    }

    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterStorageServiceServer(grpcServer, &worker{
        storage: make(map[string][]byte),
    })

    log.Printf("%s is running on port %s...\n", workerName, port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

package main

import (
    "context"
    "log"
    "net"
    "sync"

    pb "distributed-file-storage/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedStorageServiceServer
    chunkLocations map[string][]string
    mutex          sync.RWMutex
}

func (s *server) UploadChunk(ctx context.Context, chunk *pb.Chunk) (*pb.UploadStatus, error) {
    // Error handling with context
    if ctx.Err() == context.Canceled {
        return nil, ctx.Err()
    }

    // Store chunk metadata
    s.mutex.Lock()
    s.chunkLocations[chunk.Id] = []string{}
    s.mutex.Unlock()

    // Replicate chunk to workers
    go s.replicateChunk(chunk)

    return &pb.UploadStatus{Success: true, Message: "Chunk uploaded successfully"}, nil
}

func StartServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterStorageServiceServer(grpcServer, &server{
        chunkLocations: make(map[string][]string),
    })

    log.Println("Server is running on port 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

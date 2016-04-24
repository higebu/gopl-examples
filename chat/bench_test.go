package main

import (
	"log"
	"net"
	"testing"
)

func startServer(addr string) string {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen!!!!")
	go broadcaster()
	go listen(lis)
	return lis.Addr().String()
}

func benchmarkChat(b *testing.B, clientCount int) {
	target := startServer("localhost:0")
	b.SetParallelism(clientCount)
	b.RunParallel(func(pb *testing.PB) {
		conn, err := net.Dial("tcp", target)
		if err != nil {
			b.Fatal(err)
		}
		for pb.Next() {
			conn.Write([]byte("hello"))
		}
	})
}

func BenchmarkChatC1(b *testing.B) {
	benchmarkChat(b, 1)
}

//func BenchmarkChatC8(b *testing.B) {
//	benchmarkChat(b, 8)
//}
//
//func BenchmarkChatC64(b *testing.B) {
//	benchmarkChat(b, 64)
//}
//
//func BenchmarkChatC512(b *testing.B) {
//	benchmarkChat(b, 512)
//}

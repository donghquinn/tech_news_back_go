package utilities

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutDown(server *http.Server){
	go func() {
		// 서비스 접속
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 5초의 타임아웃으로 인해 인터럽트 신호가 서버를 정상종료 할 때까지 기다립니다.
	quit := make(chan os.Signal)
	// kill (파라미터 없음) 기본값으로 syscanll.SIGTERM를 보냅니다
	// kill -2 는 syscall.SIGINT를 보냅니다
	// kill -9 는 syscall.SIGKILL를 보내지만 캐치할수 없으므로, 추가할 필요가 없습니다.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// 5초의 타임아웃으로 ctx.Done()을 캐치합니다.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
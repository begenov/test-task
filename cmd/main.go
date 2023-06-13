package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/begenov/test-task/internal/config"
	delivery "github.com/begenov/test-task/internal/delivery/http"
	"github.com/begenov/test-task/internal/service"
	"github.com/gofiber/fiber/v2"
)

var websites = []string{
	"google.com", "youtube.com", "facebook.com", "baidu.com", "wikipedia.org", "qq.com", "taobao.com", "yahoo.com", "tmall.com",
	"amazon.com", "google.co.in", "twitter.com", "sohu.com", "jd.com", "live.com", "instagram.com", "sina.com.cn", "weibo.com",
	"google.co.jp", "reddit.com", "vk.com", "360.cn", "login.tmall.com", "blogspot.com", "yandex.ru", "google.com.hk",
	"netflix.com", "linkedin.com", "pornhub.com", "google.com.br", "twitch.tv", "pages.tmall.com", "csdn.net", "yahoo.co.jp",
	"mail.ru", "aliexpress.com",
}

func main() {

	cfg := config.NewConfig()

	service := service.NewService(websites)

	go service.Availability.CheckAvailability()

	handler := delivery.NewHandler(service)

	server := fiber.New(fiber.Config{
		ReadTimeout:    cfg.Server.ReadTimeout,
		WriteTimeout:   cfg.Server.WriteTimeout,
		ReadBufferSize: cfg.Server.MaxHeaderMegabytes << 20,
	})

	handler.Init(server)

	go func() {
		if err := server.Listen(":" + cfg.Server.Port); err != nil {
			log.Printf("Ошибка при запуске сервера: %v\n", err)
		}
	}()
	log.Printf("Сервер запущен и работает на порту %v\n", cfg.Server.Port)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.ShutdownWithContext(ctx); err != nil {
		log.Printf("Ошибка при остановке сервера: %v\n", err)
	}

	log.Println("Сервер остановлен")
}

package app

import (
	"camStream/internal/rtsp/h265"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gopkg.in/yaml.v3"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
)

type camStruct struct {
	ch chan image.Image
}

func Run(ctx context.Context) {
	//cfg := parseConfig()
	//bot, err := tgbotapi.NewBotAPI(cfg.TGToken)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//// Включаем дебаг-режим, если необходимо
	//bot.Debug = true
	//
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	//
	//u := tgbotapi.NewUpdate(0)
	//u.Timeout = 60

	camSt := camStruct{ch: make(chan image.Image, 10)}

	go h265.SetRTSPH265(camSt.ch)
	startServer2(camSt)

}

func startServer2(cam camStruct) {
	r := chi.NewRouter()

	// Использование стандартных middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Обслуживание статических изображений из папки ./images по пути /static/
	FileServer(r, "/static", http.Dir("./images"))

	// Маршруты для разных способов отправки изображений
	r.Get("/dynamic", cam.getImageFromCam)

	// Запуск сервера на порту 8080
	log.Println("Сервер запущен на :7070")
	if err := http.ListenAndServe(
		":8080",
		r,
	); err != nil {
		log.Fatal(err)
	}
}

func (c *camStruct) getImageFromCam(w http.ResponseWriter, r *http.Request) {
	fmt.Println("before chan")
	img, ok := <-c.ch
	if !ok {
		fmt.Println("chan closed")
		return
	}
	fmt.Println("after chan")
	w.Header().Set("Content-Type", "image/img")
	// Отправка изображения в формате PNG
	if err := png.Encode(w, img); err != nil {
		http.Error(w, "Не удалось отправить изображение.", http.StatusInternalServerError)
		return
	}

	return
}
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Fatal("FileServer не поддерживает параметры маршрута")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	r.Get(path+"/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

func parseConfig() *Config {
	f, err := os.Open("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

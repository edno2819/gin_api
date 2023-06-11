package config

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupOutputGin(config *Config) {
	log_path := fmt.Sprintf("%s/log_%v.log", config.LogPath, time.Now().Unix())
	s := fmt.Sprintf(log_path)
	f, _ := os.Create(s)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

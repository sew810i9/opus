package main

import (
	"fmt"
	"github.com/sew810i9/opus"
	"github.com/sew810i9/opus/configs"
	"github.com/sew810i9/opus/internal/delivery/http/handler"
	"github.com/sew810i9/opus/internal/domain/issue"
	"github.com/sew810i9/opus/internal/domain/user"
	"github.com/sew810i9/opus/internal/store/mysql"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s := new(opus.Server)
	c := configs.Init()

	connection, err := mysql.NewConnection(c)
	if err != nil {
		panic(err)
	}

	userRepo := user.NewRepository(connection)
	issueRepo := issue.NewRepository(connection)
	userService := user.NewService(userRepo)
	issueService := issue.NewService(issueRepo)

	h := handler.NewHandler(userService, issueService)

	fmt.Println("==> Starting app on port " + c.Port)

	if err := s.Run(c.Port, h.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

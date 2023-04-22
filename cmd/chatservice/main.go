package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
	"github.com/visko93/chatservice/cmd/configs"
	"github.com/visko93/chatservice/internal/infra/grpc/server"
	"github.com/visko93/chatservice/internal/infra/repository"
	"github.com/visko93/chatservice/internal/infra/web"
	"github.com/visko93/chatservice/internal/infra/web/webserver"
	"github.com/visko93/chatservice/internal/usecase/chatcompletion"
	"github.com/visko93/chatservice/internal/usecase/chatcompletionstream"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	conn, err := sql.Open(configs.DBDriver, fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := repository.NewChatRepositoryPostgres(conn)
	client := openai.NewClient(configs.OpenAIApiKey)

	chatConfig := chatcompletion.ChatCompletionConfigInputDTO{
		Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
	}

	chatConfigStream := chatcompletionstream.ChatCompletionConfigInputDTO{
		Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
	}

	usecase := chatcompletion.NewChatCompletionUseCase(repo, client)

	streamChannel := make(chan chatcompletionstream.ChatCompletionOutputDTO)
	usecaseStream := chatcompletionstream.NewChatCompletionUseCase(repo, client, streamChannel)

	grpcServer := server.NewGRPCServer(*usecaseStream, chatConfigStream, configs.GRPCServerPort, configs.AuthToken, streamChannel)
	fmt.Println("gRPC server running on port " + configs.GRPCServerPort)
	go grpcServer.Start()

	webserver := webserver.NewWebServer(":" + configs.WebServerPort)
	webserverChatHandler := web.NewWebChatGPTHandler(*usecase, chatConfig, configs.AuthToken)
	webserver.AddHandler("/chat", webserverChatHandler.Handle)

	fmt.Println("Server running on port " + configs.WebServerPort)
	webserver.Start()

}

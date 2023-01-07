package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	request := gpt3.CompletionRequest{
		Prompt:    []string{"Tell me a bedtime story"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	}
	err := client.CompletionStream(ctx, request, onData)

	if err != nil {
		log.Fatalln(err)
	}
}

func onData(resp *gpt3.CompletionResponse) {
	fmt.Println(resp.Choices[0].Text)
}

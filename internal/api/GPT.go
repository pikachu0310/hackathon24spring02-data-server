package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
)

type FinishReason int

const (
	Stop FinishReason = iota
	Length
	ErrorHappen
)

var (
	apiKey string
)

type Message = openai.ChatCompletionMessage
type Messages = []openai.ChatCompletionMessage

func InitGPT() {
	apiKey = getAPIKey()
}

func getAPIKey() string {
	key, exist := os.LookupEnv("OPENAI_API_KEY")
	if !exist {
		log.Fatal("OPENAI_API_KEY is not set")
	}

	return key
}

func RequestGPTAndGetResponseText(messages *Messages) (responseText string, finishReason FinishReason, err error) {
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	model := openai.GPT4oMini

	req := openai.ChatCompletionRequest{
		Model:    model,
		Messages: *messages,
	}

	response, err := c.CreateChatCompletion(ctx, req)
	if err != nil {
		return
	}

	if err != nil {
		finishReason = ErrorHappen
	}
	if errors.Is(err, io.EOF) {
		_ = errors.New("stream closed")
		fmt.Println("steam closed")
		finishReason = ErrorHappen
	}

	if response.Choices[0].FinishReason == "Stop" {
		finishReason = Stop
	} else if response.Choices[0].FinishReason == "Length" {
		finishReason = Length
	}

	responseText = response.Choices[len(response.Choices)-1].Message.Content

	return
}

func CreateNewMessages() (messages *Messages) {
	return &Messages{}
}

func AddMessageAsUser(messages *Messages, message string) {
	newMessage := Message{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}

	*messages = append(*messages, newMessage)
}

func AddMessageAsAssistant(messages *Messages, message string) {
	newMessage := Message{
		Role:    openai.ChatMessageRoleAssistant,
		Content: message,
	}

	*messages = append(*messages, newMessage)
}

func AddSystemMessageIfNotExist(messages *Messages, message string) {
	newMessage := Message{
		Role:    openai.ChatMessageRoleSystem,
		Content: message,
	}

	*messages = append(*messages, newMessage)
}

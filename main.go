package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	c := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: "Recording.mp3",
		Language: "fa",
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	os.WriteFile("text.txt", []byte(resp.Text), 0622)
	fmt.Println(resp.Text)
	res, err := c.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Fix spelling and grammatical issues of the provided persian content",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: resp.Text,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(res.Choices[0].Message.Content)
	os.WriteFile("text_fixed.txt", []byte(res.Choices[0].Message.Content), 0622)

}

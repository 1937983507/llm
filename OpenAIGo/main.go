package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	// 开启日志记录
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 是否流式输出
	isStream := true

	// 构建 LLM 客户端
	client := openai.NewClientWithConfig(
		openai.DefaultAnthropicConfig("xx", "http://localhost:8080/v1"),
	)

	// 用户问题
	prompt := "你好，请你介绍一下自己"

	if isStream {
		// 流式输出

		// 发起请求
		stream, err := client.CreateChatCompletionStream(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: prompt,
					},
				},
				Stream: true,
			},
		)
		if err != nil {
			log.Fatalf("ChatCompletion error: %v\n", err)
			return
		}
		defer stream.Close()

		// 流式输出，不断返回结果
		log.Println("Stream response: ")
		for {
			var response openai.ChatCompletionStreamResponse
			response, err = stream.Recv()
			if errors.Is(err, io.EOF) {
				log.Println("Stream finished")
				return
			}

			if err != nil {
				log.Fatalf("Stream error: %v\n", err)
				return
			}

			// 输出结果
			jsonData, _ := json.Marshal(response)
			log.Println(string(jsonData) + "\n")
		}

	} else {
		// 非流式输出

		// 发起请求
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: prompt,
					},
				},
				Stream: true,
			},
		)
		if err != nil {
			log.Fatalf("ChatCompletion error: %v\n", err)
			return
		}

		// 返回结果
		log.Println(resp.Choices[0].Message.Content)
	}

}

package main

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func main() {

	// 开启日志记录
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 构建上下文
	ctx := context.Background()

	// 创建 LLM
	llm, err := openai.New(
		openai.WithBaseURL("http://localhost:8080/v1"), // 向本地部署的模型发起请求
		openai.WithToken("xxx"),                        // api token可任意填，因为当前模型是部署在本地的
	)
	if err != nil {
		log.Fatal(err)
	}

	// 用户问题
	prompt := "你好，请问你能做些什么呢?"

	// 构建请求消息
	messages := []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextPart(prompt)},
		},
		// 可以继续添加历史消息，再一起发起请求
	}

	// 发起请求
	res, err := llm.GenerateContent(ctx, messages)
	if err != nil {
		log.Fatal(err)
	}

	// 输出结果
	for _, one := range res.Choices {
		log.Println(one.Content)
	}

	// // 若是只发1条消息，则可以直接调用这个函数，其内部会自动构建 message
	// answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(answer)

}

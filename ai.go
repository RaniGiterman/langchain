package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
)

const (
	URL           = "https://www.benda.co.il/product/36208-498-32/"
	SYSTEM_PROMPT = `
You are a precise data extraction assistant. You will receive an  page that describes a consumer product with details and must extract information to fill a JSON object with exactly three fields.

## Input Format
You will receive:
1. An empty JSON template: {"title": "", "description": "", "img": ""}
2. A base URL for the website
3. A complete HTML page

## Field Descriptions

*title* (string):
- Extract the product title
- Return only the text content, no HTML tags
- If multiple candidates exist, choose the most descriptive and prominent one

*description* (string):
- Extract product description from the page content

*img* (string):
- Extract the product image. if more than one, then concat them all with comma
- If image URL is relative (starts with / or does not include http/https), prepend the base URL to create absolute URL
- Example: if base URL is "https://example.com" and img is "/images/product.png", return "https://example.com/images/product.png"

## Output Requirements
- Return ONLY the filled JSON object
- Ensure valid JSON syntax
- Use empty string "" for any field that cannot be determined

Now process the provided URL and return the filled JSON.
`
)

func run() error {
	llm, err := openai.New(
		openai.WithModel("gpt-4o-mini"),
	)
	if err != nil {
		return err
	}
	agentTools := []tools.Tool{
		// tools.Calculator{},
		GetPage{},
	}

	agent := agents.NewOneShotAgent(llm,
		agentTools,
		agents.WithMaxIterations(3),
	)
	executor := agents.NewExecutor(agent)

	// question := "What is the title of the page: https://www.hvr.co.il/signin.aspx"
	question := fmt.Sprintf("%s %s", SYSTEM_PROMPT, URL)
	answer, err := chains.Run(context.Background(), executor, question)
	fmt.Println(answer)
	return err
}

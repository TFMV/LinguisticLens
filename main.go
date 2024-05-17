package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

func main() {
	r := gin.Default()
	r.POST("/analyze", analyzeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func analyzeHandler(c *gin.Context) {
	var req struct {
		Text string `json:"text" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received text for analysis: %s", req.Text)

	result, err := analyzeText(req.Text)
	if err != nil {
		log.Printf("Error analyzing text: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Analysis result: %s", result)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func analyzeText(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("missing the OpenAI API key, set it in the OPENAI_API_KEY environment variable")
	}

	os.Setenv("OPENAI_API_KEY", apiKey)

	llm, err := openai.New(openai.WithModel("gpt-4"))
	if err != nil {
		return "", fmt.Errorf("failed to create OpenAI client: %w", err)
	}

	prompt := prompts.NewPromptTemplate(
		"Analyze the specified content, considering emotional, factual, and implicit aspects. Identify and explore the presence of dark triad traits (narcissism, Machiavellianism, and psychopathy) in each party involved, and examine how these traits manifest in their behavior and interactions. Additionally, analyze the hidden meaning and tone of the text, and explain how the underlying messages and tonal nuances influence the narrative and character dynamics. {{.text}}",
		[]string{"text"},
	)
	llmChain := chains.NewLLMChain(llm, prompt)

	ctx := context.Background()
	outputValues, err := chains.Call(ctx, llmChain, map[string]any{
		"text": text,
	})
	if err != nil {
		return "", fmt.Errorf("failed to call LLM chain: %w", err)
	}

	out, ok := outputValues[llmChain.OutputKey].(string)
	if !ok {
		return "", fmt.Errorf("invalid chain return")
	}

	return out, nil
}

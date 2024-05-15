# LinguisticLens

![LinguisticLens](ll.webp)

This repository contains a simple API for analyzing text using OpenAI's language model. The API takes a text input and provides a comprehensive analysis considering emotional, factual, and implicit aspects. It identifies and explores the presence of dark triad traits (narcissism, Machiavellianism, and psychopathy), examines the hidden meaning and tone, and explains how these influence the narrative and character dynamics.

## Underlying Technology

### Language Model

The core of the application leverages OpenAI's language model to perform sophisticated text analysis. The model is accessed using the `github.com/tmc/langchaingo` package, which provides a streamlined interface for integrating language models into Go applications.

### Prompt Template

A customized prompt template guides the language model's analysis:

Analyze the specified content, considering emotional, factual, and implicit aspects. Identify and explore the presence of dark triad traits (narcissism, Machiavellianism, and psychopathy) in each party involved, and examine how these traits manifest in their behavior and interactions. Additionally, analyze the hidden meaning and tone of the text, and explain how the underlying messages and tonal nuances influence the narrative and character dynamics.

### API Design

The API is built using the `github.com/gin-gonic/gin` framework, known for its high performance and minimalistic design. The main components of the application include:

- **main.go**: The entry point of the application, setting up the HTTP server and defining the `/analyze` endpoint.
- **analyzeHandler**: The handler function for the `/analyze` endpoint. It processes the input text, calls the `analyzeText` function, and returns the result.
- **analyzeText**: This function communicates with the OpenAI API to perform the analysis based on the provided prompt template.

### Flow of Execution

1. **Input Handling**: The API receives a POST request with the text to be analyzed.
2. **Text Analysis**: The `analyzeText` function formulates a request to the OpenAI API using the prompt template and the input text.
3. **Response Generation**: The analysis result is returned as a JSON response.

### Environment Variables

The application requires an OpenAI API key, which is accessed via the `OPENAI_API_KEY` environment variable.

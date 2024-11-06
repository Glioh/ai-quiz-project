package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"quiz.com/quiz/internal/entity"
)

type AIService struct {
	apiKey string
}

type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type GeneratedQuiz struct {
	Questions []entity.QuizQuestion `json:"questions"`
}

func NewAIService() *AIService {

	apiKey := "YOUR KEY"
	if apiKey == "" {
		fmt.Println("Warning: OPENAI_API_KEY is not set")
	}

	return &AIService{
		apiKey: apiKey,
	}
}

func (s *AIService) GenerateQuiz(prompt string, numQuestions int) ([]entity.QuizQuestion, error) {
	// Add validation
	if numQuestions < 1 || numQuestions > 20 {
		return nil, errors.New("number of questions must be between 1 and 20")
	}

	systemPrompt := fmt.Sprintf(`Generate a quiz based on the given topic. Return the response in valid JSON format with the following structure:
    {
        "questions": [
            {
                "id": "uuid-string",
                "name": "question text",
                "time": 25,
                "choices": [
                    {
                        "id": "uuid-string",
                        "name": "choice text",
                        "correct": boolean
                    }
                ]
            }
        ]
    }
    Rules:
    - Generate exactly %d questions
    - [IMPORTANT] Each question AND answer should have a maximum of 25 characters (including spaces)
    - Each question should have exactly 4 choices
    - Only one choice should be correct
    - Each question should have a reasonable time between 10-30 seconds to solve
    - Generate valid UUIDs for all IDs
    - Ensure all fields match the exact names shown above
    - Return ONLY the JSON response, no additional text`, numQuestions) // Added this line

	messages := []Message{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: prompt},
	}

	reqBody := OpenAIRequest{
		Model:       "gpt-3.5-turbo",
		Messages:    messages,
		Temperature: 0.7,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, fmt.Errorf("API error (status %d)", resp.StatusCode)
		}
		return nil, fmt.Errorf("API error: %v", errorResponse)
	}

	var openAIResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return nil, fmt.Errorf("error decoding OpenAI response: %w", err)
	}

	if len(openAIResp.Choices) == 0 {
		return nil, errors.New("no response from AI")
	}

	// Add debug logging
	fmt.Printf("Raw AI response content: %s\n", openAIResp.Choices[0].Message.Content)

	// Parse the AI response into quiz questions
	var generatedQuiz GeneratedQuiz
	if err := json.Unmarshal([]byte(openAIResp.Choices[0].Message.Content), &generatedQuiz); err != nil {
		return nil, fmt.Errorf("error parsing AI response JSON: %w\nResponse content: %s", err, openAIResp.Choices[0].Message.Content)
	}

	if len(generatedQuiz.Questions) == 0 {
		return nil, errors.New("no questions generated")
	}

	// Verify the number of questions matches the request
	if len(generatedQuiz.Questions) != numQuestions {
		return nil, fmt.Errorf("AI generated %d questions but %d were requested", len(generatedQuiz.Questions), numQuestions)
	}

	return generatedQuiz.Questions, nil
}

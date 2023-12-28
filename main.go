package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type Questions struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	QuestionId   string     `json:"question_id"`
	QuestionText string     `json:"question_text"`
	Responses    []Response `json:"responses"`
}

type Response struct {
	ResponseId    string `json:"response_id"`
	ResponseValue int    `json:"response_value"`
	ResponseText  string `json:"response_text"`
}

type NormalizedQuestion struct {
	NormalizedQuestionText string               `json:"norm_question_text"`
	OriginalQuestionText   string               `json:"original_question_text"`
	QuestionType           string               `json:"question_type"`
	Year                   int                  `json:"year"`
	Election               string               `json:"election"`
	NormalizedResponses    []NormalizedResponse `json:"norm_responses"`
}

type NormalizedResponse struct {
	NormalizedResponseText string `json:"norm_response_text"`
	OriginalResponseText   string `json:"original_response_text"`
}

func processFile(fileName string) []Question {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		os.Exit(1)
	}

	var questions Questions
	err = json.Unmarshal(file, &questions)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data:", err)
		os.Exit(1)
	}

	return questions.Questions
}

func normalizeQuestion(question Question) NormalizedQuestion {
	tmpl, err := template.ParseFiles("prompt.txt")
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	// Execute the template, writing to os.Stdout or another writer.
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, question)
	if err != nil {
		log.Fatalf("Error executing template: %s", err)
	}

	prompt := buf.String()

	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	completion, err := llm.Call(context.Background(), prompt, llms.WithMaxTokens(2048))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)

	var llmResp NormalizedQuestion
	e := json.Unmarshal([]byte(completion), &llmResp)
	if e != nil {
		log.Fatal(e)
	}

	return llmResp
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: prenorm <file_name>")
		return
	}

	questions := processFile(os.Args[1])

	var normQuestions []NormalizedQuestion

	for _, q := range questions {
		normQuestions = append(normQuestions, normalizeQuestion(q))
	}

}

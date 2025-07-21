package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"net/http"
)

const (
	openJTalkDirPath = "open_jtalk_dic_utf_8-1.11"
	onnxruntimePath  = "libvoicevox_onnxruntime.so.1.17.3"
	voiceModelPath   = "vvms/0.vvm"
	styleID          = 3
)

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Filename string `json:"filename"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get Environment
	bucketName := os.Getenv("BUCKET_NAME")

	
	// Parse Request
	var req Request
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return errorResponse(http.StatusBadRequest, "不正なJSONなのだ")
	}

	// Generate Voice
	voicevox := VoicevoxWrapper{}
	data, err := voicevox.generate(
		req.Text,
		openJTalkDirPath,
		onnxruntimePath,
		voiceModelPath,
		styleID,
	)
	if err != nil {
		fmt.Printf("音声の生成に失敗したのだ: %s", err.Error())
		return errorResponse(http.StatusInternalServerError, "音声の生成に失敗したのだ")
	}

	// Load AWS Setting
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Printf("SDKの読み込みに失敗したのだ: %s", err.Error())
		return errorResponse(http.StatusInternalServerError, "SDKの読み込みに失敗したのだ")
	}

	// Upload Audio File
	filename := fmt.Sprintf("%s.wav", uuid.New().String())
	client := s3.NewFromConfig(cfg)
	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("audio/wav"),
	}
	_, err = client.PutObject(ctx, input)
	if err != nil {
		fmt.Printf("S3へのアップロードに失敗したのだ: %s", err.Error())
		return errorResponse(http.StatusInternalServerError, "S3へのアップロードに失敗したのだ")
	}

	// Return Response
	body, _ := json.Marshal(Response{Filename: filename})
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func errorResponse(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(MessageResponse{Message: message})
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}

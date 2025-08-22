package main

import (
	"github.com/amazon-sagemaker-metrics-service/mcp-server/config"
	"github.com/amazon-sagemaker-metrics-service/mcp-server/models"
	tools_batchputmetrics "github.com/amazon-sagemaker-metrics-service/mcp-server/tools/batchputmetrics"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_batchputmetrics.CreateBatchputmetricsTool(cfg),
	}
}

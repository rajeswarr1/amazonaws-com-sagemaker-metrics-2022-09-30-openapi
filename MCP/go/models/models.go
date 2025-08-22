package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// BatchPutMetricsRequest represents the BatchPutMetricsRequest schema from the OpenAPI specification
type BatchPutMetricsRequest struct {
	Metricdata interface{} `json:"MetricData"`
	Trialcomponentname interface{} `json:"TrialComponentName"`
}

// RawMetricData represents the RawMetricData schema from the OpenAPI specification
type RawMetricData struct {
	Step interface{} `json:"Step,omitempty"`
	Timestamp interface{} `json:"Timestamp"`
	Value interface{} `json:"Value"`
	Metricname interface{} `json:"MetricName"`
}

// BatchPutMetricsError represents the BatchPutMetricsError schema from the OpenAPI specification
type BatchPutMetricsError struct {
	Code interface{} `json:"Code,omitempty"`
	Metricindex interface{} `json:"MetricIndex,omitempty"`
}

// BatchPutMetricsResponse represents the BatchPutMetricsResponse schema from the OpenAPI specification
type BatchPutMetricsResponse struct {
	Errors interface{} `json:"Errors,omitempty"`
}

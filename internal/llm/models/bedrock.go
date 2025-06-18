package models

const (
	ProviderBedrock ModelProvider = "bedrock"

	// Models
	BedrockClaude37Sonnet ModelID = "bedrock.claude-3.7-sonnet"
	BedrockClaude4Sonnet  ModelID = "bedrock.claude-4.0-sonnet"
	BedrockClaude4Opus    ModelID = "bedrock.claude-4.0-opus"
)

var BedrockModels = map[ModelID]Model{
	BedrockClaude37Sonnet: {
		ID:                  BedrockClaude37Sonnet,
		Name:                "Bedrock: Claude 3.7 Sonnet",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-3-7-sonnet-20250219-v1:0",
		CostPer1MIn:         3.0,
		CostPer1MInCached:   3.75,
		CostPer1MOutCached:  0.30,
		CostPer1MOut:        15.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    50_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	BedrockClaude4Sonnet: {
		ID:                  BedrockClaude4Sonnet,
		Name:                "Bedrock: Claude 4 Sonnet",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-sonnet-4-20250514-v1:0",
		CostPer1MIn:         3.0,
		CostPer1MInCached:   3.75,
		CostPer1MOutCached:  0.30,
		CostPer1MOut:        15.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    50_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
	BedrockClaude4Opus: {
		ID:                  BedrockClaude4Opus,
		Name:                "Bedrock: Claude 4 Opus",
		Provider:            ProviderBedrock,
		APIModel:            "anthropic.claude-opus-4-20250514-v1:0",
		CostPer1MIn:         15.0,
		CostPer1MInCached:   18.0,
		CostPer1MOutCached:  1.50,
		CostPer1MOut:        75.0,
		ContextWindow:       200_000,
		DefaultMaxTokens:    50_000,
		CanReason:           true,
		SupportsAttachments: true,
	},
}

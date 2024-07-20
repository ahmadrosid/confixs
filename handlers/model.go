package handlers

type GetFileContentRequest struct {
	FilePath string `json:"filePath"`
}

type ResourceUsageResponse struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
	Disk   float64 `json:"disk"`
}

package handler

type VersionResponse struct {
    Service   string `json:"service"`
    Env       string `json:"env"`
    Version   string `json:"version"`
    Commit    string `json:"commit"`
    BuildTime string `json:"buildTime"`
}

type HealthResponse struct {
    Status string `json:"status"`
}

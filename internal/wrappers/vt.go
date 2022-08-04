package wrappers

type File struct {
	Capabilities          []string   `json:"capabilities_tags"`
	CreationDate          int        `json:"creation_date"`
	FirstSubmissionDate   int        `json:"first_submission_date"`
	LastAnalysisDate      int        `json:"last_analysis_date"`
	AnalysisResults       []Analysis `json:"last_analysis_results"`
	LastModificationDate  int        `json:"last_modification_date"`
	LastSubmissionDate    int        `json:"last_submission_date"`
	MalwareMeaningFulName string     `json:"meaningful_name"`
	MalwareNames          []string   `json:"names"`
	MD5                   string     `json:"md5"`
	SHA1                  string     `json:"sha1"`
	SHA256                string     `json:"sha256"`
}

type Analysis struct {
	Category      string `json:"category"`
	EngineName    string `json:"engine_name"`
	EngineUpdate  string `json:"engine_update"`
	EngineVersion string `json:"engine_version"`
	Method        string `json:"method"`
	Result        string `json:"result"`
}

type VirusTotal interface {
	GetFileIntel(hash string) (File, error)
}

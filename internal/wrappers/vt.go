package wrappers

type VirusTotalResponse struct {
	Attributes            []string     `json:"capabilities_tags"`
	CreationDate          int          `json:"creation_date"`
	FirstSubmissionDate   int          `json:"first_submission_date"`
	LastAnalysisDate      int          `json:"last_analysis_date"`
	ImportListData        []ImportList `json:"import_list"`
	AnalysisResults       []Analysis   `json:"last_analysis_results"`
	LastModificationDate  int          `json:"last_modification_date"`
	LastSubmissionDate    int          `json:"last_submission_date"`
	MalwareMeaningfulName string       `json:"meaningful_name"`
	MalwareNames          []string     `json:"names"`
	MD5                   string       `json:"md5"`
	SHA1                  string       `json:"sha1"`
	SHA256                string       `json:"sha256"`
}

type Analysis struct {
	Category      string `json:"category"`
	EngineName    string `json:"engine_name"`
	EngineUpdate  string `json:"engine_update"`
	EngineVersion string `json:"engine_version"`
	Method        string `json:"method"`
	Result        string `json:"result"`
}

type ImportList struct {
	LibraryName       string                     `json:"library_name"`
	ImportedFunctions []ImportedFunctionsEntries `json:"imported_functions"`
}

type ImportedFunctionsEntries struct {
	Function []string
}

type VirusTotalWrapper interface {
	GetFileAnalysis(hash string) (VirusTotalResponse, error)
}

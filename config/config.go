package config

type Config struct {
	Crop struct {
		Height int    `json:"height"`
		OutDir string `json:"outDir"`
		Width  int    `json:"width"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
	} `json:"crop"`
	Flip struct {
		OutDir string `json:"outDir"`
	} `json:"flip"`
	Generate struct {
		OutDir            string     `json:"outDir"`
		OutFileNameFormat string     `json:"outFileNameFormat"`
		SrcPatterns       [][]string `json:"srcPatterns"`
	} `json:"generate"`
	Paste struct {
		Col               int    `json:"col"`
		Height            int    `json:"height"`
		OutDir            string `json:"outDir"`
		OutFileNameFormat string `json:"outFileNameFormat"`
		Row               int    `json:"row"`
		Width             int    `json:"width"`
	} `json:"paste"`
	Scale struct {
		OutDir string `json:"outDir"`
		Size   int    `json:"size"`
	} `json:"scale"`
}

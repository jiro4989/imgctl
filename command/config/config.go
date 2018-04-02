package config

type TOML struct {
	Images []Image `toml:"image"`
}

type Image struct {
	OutDir   string   `toml:"out_dir"`
	Scale    Scale    `toml:"scale"`
	Trim     Trim     `toml:"trim"`
	Paste    Paste    `toml:"paste"`
	Generate Generate `toml:"generate"`
}

type Scale struct {
	Size int `toml:"size"`
}

type Trim struct {
	X      int `toml:"x"`
	Y      int `toml:"y"`
	Width  int `toml:"width"`
	Height int `toml:"height"`
}

type Paste struct {
	Row    int `toml:"row"`
	Column int `toml:"column"`
}

type Generate struct {
	SaveFilenameFormat string     `toml:"save_filename_format"`
	Pattern            [][]string `toml:"pattern"`
}

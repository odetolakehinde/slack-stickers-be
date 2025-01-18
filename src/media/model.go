package media

type (
	// UploadPayload a struct defining payload to for a media
	UploadPayload struct {
		Name string
	}

	// TenorAPIResponse is tenor search api response
	TenorAPIResponse struct {
		Results []struct {
			ID           string `json:"id"`
			Title        string `json:"title"`
			MediaFormats struct {
				Gif struct {
					URL      string  `json:"url"`
					Duration float64 `json:"duration"`
					Preview  string  `json:"preview"`
					Dims     []int   `json:"dims"`
					Size     int     `json:"size"`
				} `json:"gif"` // we only care about the gif media format here. ignoring other formats (tinygif,mp4,tinymp4)
			} `json:"media_formats"`
			Created                  float64       `json:"created"`
			ContentDescription       string        `json:"content_description"`
			Itemurl                  string        `json:"itemurl"`
			URL                      string        `json:"url"`
			Tags                     []string      `json:"tags"`
			Flags                    []interface{} `json:"flags"`
			Hasaudio                 bool          `json:"hasaudio"`
			ContentDescriptionSource string        `json:"content_description_source"`
		} `json:"results"`
		Next string `json:"next"`
	}
)

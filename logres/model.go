package logger

type Context struct {
	ServiceName    string      `json:"app_name"`
	ServiceVersion string      `json:"app_version"`
	ServicePort    int         `json:"app_port"`
	ThreadID       string      `json:"app_thread_id"`
	Header         interface{} `json:"header"`
	ReqMethod      string      `json:"app_method"`
	ReqURI         string      `json:"app_uri"`
}

type LogTdrModel struct {
	RequestId    string `json:"request_id"`
	Path         string `json:"path"`
	Method       string `json:"method"`
	Port         int    `json:"port"`
	RespTime     int64  `json:"rt"`
	ResponseCode string `json:"rc"`

	Header   interface{} `json:"header"`
	Request  interface{} `json:"req"`
	Response interface{} `json:"resp"`
	Error    string      `json:"error"`
}

type LogresConfig struct {
	MaxSize    int    // Ukuran maksimum file dalam megabyte sebelum rotasi
	MaxBackups int    // Jumlah file backup yang akan disimpan
	MaxAge     int    // Maksimum umur file dalam hari sebelum dihapus (1 hari)
	Compress   bool   // Mengompresi file-file lama
	LocalTime  bool   // Menggunakan waktu lokal
	FolderPath string // Lokasi untuk penempatan file logger
	LogsWrite  bool   // menulis logger dalam file
}

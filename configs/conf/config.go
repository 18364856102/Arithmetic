package conf

type HTTPConfig struct {
	HttpAddr string `default:"127.0.0.1:8081"`
}

var Conf *HTTPConfig
package main

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/viper"
	"goskeleton/app/global/variable"
	_ "goskeleton/bootstrap"
	"goskeleton/routers"
	"net/http"
)

type CertificateConfig struct {
	Name string `mapstructure:"name"`
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}

func main() {
	router := routers.InitApiRouter()

	if variable.ConfigYml.GetBool("HttpServer.UseHttps") {
		viper.SetConfigName("config")   // 配置文件名 (不包括扩展名)
		viper.AddConfigPath("./config") // 查找配置文件所在的路径列表
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		var certificateConfigs []CertificateConfig
		if err := viper.UnmarshalKey("HttpServer.domains", &certificateConfigs); err != nil {
			return
		}

		Certificates := make(map[string]*tls.Certificate)
		for _, v := range certificateConfigs {
			if cert, err := tls.LoadX509KeyPair(v.Cert, v.Key); err != nil {
				panic(fmt.Sprint("Unable to load cert", v.Cert, v.Key, err))
			} else {
				Certificates[v.Name] = &cert
			}
		}

		cfg := &tls.Config{
			GetCertificate: func(ch *tls.ClientHelloInfo) (*tls.Certificate, error) {
				cert := Certificates[ch.ServerName]
				return cert, nil
			},
		}

		srv := &http.Server{
			Addr:      variable.ConfigYml.GetString("HttpServer.Port"),
			Handler:   router,
			TLSConfig: cfg,
		}

		_ = srv.ListenAndServeTLS("", "")

	} else {
		_ = router.Run(variable.ConfigYml.GetString("HttpServer.Port"))
	}

}

package cmd

import (
	"context"
	"crypto/tls"
	"doit/urlshortener/internal/application/usecase/shortener"
	"doit/urlshortener/internal/domain/service"
	"doit/urlshortener/internal/infrastructure/repository/url_shorten"
	"doit/urlshortener/internal/infrastructure/utils/clock"
	"doit/urlshortener/internal/presentation/http/route"
	"doit/urlshortener/pkg/dotenv"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	helmet "github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type MainConnection struct{}

func main() {
	var (
		connMain             = MainConnection{}
		serverListentTimeOut = dotenv.GetInt("SERVER_LISTEN_TIMEOUT", 5)
	)

	dotenv.Environment()

	router := connMain.SetupRouter()

	srv := &http.Server{
		Addr:              ":" + dotenv.APPPORT(),
		Handler:           router,
		ReadHeaderTimeout: time.Duration(serverListentTimeOut) * time.Second,
	}

	isSSL := dotenv.GetBool("USE_SSL", false)
	if isSSL {
		srv.TLSConfig = &tls.Config{
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				//#nosec
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}
		certFilePath := os.Getenv("APP_BASE_DIR") + os.Getenv("PATH_SSL_CERT")
		keyFilePath := os.Getenv("APP_BASE_DIR") + os.Getenv("PATH_SSL_KEY")

		go func() {
			fmt.Println("server is running with SSL on port " + dotenv.APPPORT())
			err := srv.ListenAndServeTLS(certFilePath, keyFilePath)
			if err != nil && err != http.ErrServerClosed {
				log.Printf("Server error %s", err.Error())
			}
		}()
	} else {
		go func() {
			log.Println("server is running without SSL on port " + dotenv.APPPORT())
			err := srv.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Printf("server error %s", err.Error())
			}
		}()
	}

	// Handle graceful shutdown
	connMain.gracefulShutdown(srv, 10*time.Second)
}

func (conn *MainConnection) SetupRouter() *gin.Engine {
	// Setup router
	router := gin.New()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 10 << 20 // 10 MiB
	router.RemoveExtraSlash = true

	// Setup Mode Application
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Call global middleware
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	var (
		repo        = url_shorten.NewUrlMappingRepository()
		service     = service.NewShortenerService()
		systemClock = clock.NewSystemClock()
		uc          = shortener.NewShortenerUsecase(repo, service, systemClock)
	)

	// Initiate routes
	route.InitShortenRoute(router, uc)

	return router
}

func (conn *MainConnection) gracefulShutdown(srv *http.Server, timeout time.Duration) {
	// Create a channel to listen for OS signals (like Ctrl+c)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Graceful shutdown of server failed: %v", err)
	}

	// Close all other connections gracefully
	log.Println("Server exited gracefully")
}

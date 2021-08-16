package main

import (
	"MaisrForAdvancedSystems/go-biller/engine"
	"MaisrForAdvancedSystems/go-biller/masservices"
	"MaisrForAdvancedSystems/go-biller/middlewares"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"strings"
	"time"

	"log"

	chrg "github.com/MaisrForAdvancedSystems/biller-charger"
	prov "github.com/MaisrForAdvancedSystems/biller-mas-provider"
	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	//"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var empty *billing.Empty = &billing.Empty{}

//go:embed public/*
var content embed.FS

func main() {
	port := 25566
	port2 := 25568
	flag.Parse()
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(middlewares.TokenAuthFunc)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middlewares.TokenAuthFunc)),
	}
	grpcServer := grpc.NewServer(opts...)
	grpcServerh2 := grpc.NewServer()
	charger := &chrg.BillingChargeService{IsTrace: false}
	masProvider := &prov.MasProvider{}
	engi, err := engine.NewEngine(masProvider, charger, masProvider, masProvider)
	if err != nil {
		log.Println(err)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("serive listen:%v", port)
	log.Printf("serivehttp2 listen:%v", port2)
	billing.RegisterEngineServer(grpcServer, engi)
	billing.RegisterEngineServer(grpcServerh2, engi)
	wrappedServer := grpcweb.WrapServer(grpcServer /*, grpcweb.WithWebsockets(true)*/)
	/*go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50551))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("listening on grpc")
		//var opts []grpc.ServerOption
		grpcServer.Serve(lis)
	}()*/
	addr := fmt.Sprintf(":%v", port)
	addrr := fmt.Sprintf(":%v", port2)
	//STATIC FILE SERVER
	fsys := fs.FS(content)
	contentStatic, _ := fs.Sub(fsys, "public")
	staticFileServer := http.FileServer(http.FS(contentStatic))
	//staticFileServer := http.FileServer(http.Dir("./public"))
	httpSrv := &http.Server{
		// These interfere with websocket streams, disable for now
		ReadTimeout:       50 * time.Second,
		WriteTimeout:      100 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              addr,
		Handler: hstsHandler(
			grpcTrafficSplitter(
				staticFileServer,
				wrappedServer,
			),
		),
	}
	log.Printf("starting service :%v", engine.VERSION)
	go httpSrv.ListenAndServe()
	lis, err := net.Listen("tcp", addrr)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
		return
	}
	go grpcServerh2.Serve(lis)
	masservices.Masservicesmain()
}

func setupResponse(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "*")
}

// hstsHandler wraps an http.HandlerFunc such that it sets the HSTS header.
func hstsHandler(fn http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.Host)
		setupResponse(w)
		if r.Method == "OPTIONS" {
			return
		}
		fn(w, r)
	})
}

func grpcTrafficSplitter(fallback http.Handler, grpcHandler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Info(r.URL.String())
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			log.Println("grpc web request")
			grpcHandler.ServeHTTP(w, r)
		} else {
			log.Println("http web request")
			fallback.ServeHTTP(w, r)
		}
		logrus.Info("Done:" + r.URL.String())
	})
}

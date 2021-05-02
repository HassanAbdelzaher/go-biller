package main

import (
	"MaisrForAdvancedSystems/go-biller/engine"
	"MaisrForAdvancedSystems/go-biller/middlewares"
	stest "MaisrForAdvancedSystems/go-biller/test"
	"context"
	"flag"
	"fmt"
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
	"google.golang.org/protobuf/types/known/timestamppb"
)

var empty *billing.Empty = &billing.Empty{}

func main() {
	port := 25566
	flag.Parse()
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(middlewares.TokenAuthFunc)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middlewares.TokenAuthFunc)),
	}
	grpcServer := grpc.NewServer(opts...)
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
	billing.RegisterEngineServer(grpcServer, engi)
	wrappedServer := grpcweb.WrapServer(grpcServer /*, grpcweb.WithWebsockets(true)*/)
	addr := fmt.Sprintf(":%v", port)
	fs := http.FileServer(http.Dir("./public"))
	httpSrv := &http.Server{
		// These interfere with websocket streams, disable for now
		ReadTimeout:       50 * time.Second,
		WriteTimeout:      100 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              addr,
		Handler: hstsHandler(
			grpcTrafficSplitter(
				fs,
				wrappedServer,
			),
		),
	}
	log.Printf("starting service :%v", engine.VERSION)
	httpSrv.ListenAndServe()
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

func runTest() {
	charger := &chrg.BillingChargeService{IsTrace: true}
	masProvider := &prov.MasProvider{}
	sample := &stest.JsonTest{}
	sample.Init("test_pattern.json")
	eng, err := engine.NewEngine(masProvider, charger, masProvider, sample)
	if err != nil {
		log.Println(err)
		return
	}
	custkey := "100000992"
	var cycleLength int64 = 1
	bilngDate := time.Date(2021, 2, 28, 0, 0, 0, 0, time.UTC)
	stmapBilngDate := timestamppb.New(bilngDate)
	setting := billing.ChargeSetting{
		CycleLength:      &cycleLength,
		BilingDate:       stmapBilngDate,
		IgnoreTimeEffect: new(bool),
	}
	water := billing.SERVICE_TYPE_WATER
	sewer := billing.SERVICE_TYPE_SEWER
	var consump float64 = 155
	var zero float64 = 0
	readings := []*billing.ServiceReading{{
		ServiceType: &water,
		Reading: &billing.Reading{
			Consump:   &consump,
			PrReading: &zero,
			CrReading: &consump,
			PrDate:    timestamppb.New(time.Now().AddDate(-1, 0, 0)),
			CrDate:    timestamppb.Now(),
		},
	},
		{
			ServiceType: &sewer,
			Reading: &billing.Reading{
				Consump:   &consump,
				PrReading: &zero,
				CrReading: &consump,
				PrDate:    timestamppb.New(time.Now().AddDate(-1, 0, 0)),
				CrDate:    timestamppb.Now(),
			},
		}}
	rs, err := eng.HandleRequest(context.Background(), custkey, &setting, readings)
	if err != nil {
		log.Println(err)
		return
	}
	if rs == nil {
		log.Println("invalied responce")
	}
	if rs.Bills == nil {
		log.Println("invalied responce:null Bill")
	}
	if rs.Bills[0].FTransactions == nil {
		log.Println("invalied responce transcation")
	}
	if len(rs.Bills[0].FTransactions) == 0 {
		log.Println("invalied responce empty transcation")
	}
	log.Println("succssed")
	for _, t := range rs.Bills[0].FTransactions {
		log.Println(t.GetCode(), t.GetAmount())
	}
}

func TestJsonService() {
	errF := log.Println
	logF := log.Println
	errff := log.Printf
	logff := log.Printf
	stest.TestService(logF, errF, errff, logff, "test_patter.json")
}


package masservices

import (
	"MaisrForAdvancedSystems/go-biller/middlewares"
	"MaisrForAdvancedSystems/go-biller/tools"
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
	"time"

	"log"

	"github.com/HassanAbdelzaher/lama"
	pbMessages "github.com/MaisrForAdvancedSystems/go-biller-proto/go/messages"
	cancelled "github.com/MaisrForAdvancedSystems/go-biller-proto/go/services"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	irespo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/interfaces"
	respo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/repositories"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	//"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// //go:embed public/*
// var content embed.FS

// BARCODE_LENGTH
const BARCODE_LENGTH = int32(0)

type serverCollection struct {
	cancelled.UnimplementedCollectionServer
}

func getUser(_username *string, conn *lama.Lama) (*dbmodels.USERS, error) {
	if _username == nil {
		return nil, errors.New("اسم الدخول غير صحيح")
	}
	var users irespo.ICommonRepository = &respo.CommonRepository{Lama: conn}
	user, err := users.GetUserByUserName(*_username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("اسم الدخول او كلمة المرور غير صحيحة")
	}
	return user, nil
}
func getStation(_stationNo *int64, conn *lama.Lama) (*dbmodels.STATIONS, error) {
	if _stationNo == nil {
		return nil, errors.New("رقم الفرع غير صحيح")
	}
	stationNo := int32(*_stationNo)
	var stations irespo.ICommonRepository = &respo.CommonRepository{Lama: conn}
	station, err := stations.GetStationByStationNo(stationNo)
	if err != nil {
		return nil, err
	}
	if station == nil {
		return nil, errors.New("رقم الفرع غير معرف")
	}
	return station, nil
}
func throwsIfStationNoInvalied(user *dbmodels.USERS, stationNo *int32, conn *lama.Lama) error {
	station := int32(-99)
	if stationNo == nil {
		station = *stationNo
	}
	IS_HEADQUARTERS := int64(0)
	if user.STATION_NO != nil {
		sta, err := getStation(user.STATION_NO, conn)
		if err != nil {
			return err
		}
		if sta != nil && sta.IS_HEADQUARTERS != nil {
			IS_HEADQUARTERS = *sta.IS_HEADQUARTERS
		}
	}

	if IS_HEADQUARTERS == 1 {
		return nil
	}
	if user != nil && user.STATION_NO != nil && int32(*user.STATION_NO) == station {
		return nil
	}
	return sendError(codes.InvalidArgument, "لا تمتلك صلاحية كافية لعمل الاجراء بالفرع", "لا تمتلك صلاحية كافية لعمل الاجراء بالفرع")
}
func create_timestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
func create_time(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	ti := t.AsTime()
	return &ti
}
func Masservicesmain() {
	port := 25567
	flag.Parse()
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(middlewares.TokenAuthFunc)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middlewares.TokenAuthFunc)),
	}
	grpcServer := grpc.NewServer(opts...)
	log.Printf("serive listen:%v", port)
	cancelled.RegisterCollectionServer(grpcServer, &serverCollection{})
	wrappedServer := grpcweb.WrapServer(grpcServer /*, grpcweb.WithWebsockets(true)*/)
	addr := fmt.Sprintf(":%v", port)
	////STATIC FILE SERVER
	//fsys := fs.FS(content)
	//contentStatic, _ := fs.Sub(fsys, "public")
	//staticFileServer := http.FileServer(http.FS(contentStatic))
	staticFileServer := http.FileServer(http.Dir("./public"))
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
	log.Printf("starting service ...")
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

func sendError(codeStatus codes.Code, ErrorMessage string, ErrorTitle string) error {
	st := status.New(codeStatus, ErrorMessage)
	badRequest := &errdetails.BadRequest{}
	violations := make([]*errdetails.BadRequest_FieldViolation, 0)
	violation := errdetails.BadRequest_FieldViolation{Field: ErrorTitle, Description: ErrorMessage}
	(*badRequest).FieldViolations = append(violations, &violation)
	det, _ := st.WithDetails(badRequest)
	return det.Err()
}

func cleanString(inp *string, defult *string, toUpper *bool, toLower *bool) {
	if inp == nil {
		return
	}
	toUpperv := false
	toLowerv := false
	if toUpper != nil {
		toUpperv = *toUpper
	}
	if toLower != nil {
		toLowerv = *toLower
	}
	if *inp == "''" {
		inp = defult
	}
	inp = tools.ToStringPointer(strings.TrimSpace(*inp))
	if strings.ToLower(*inp) == "undefined" || strings.ToLower(*inp) == "null" {
		inp = defult
	}
	inp = tools.ToStringPointer(strings.TrimLeft(*inp, "'"))
	inp = tools.ToStringPointer(strings.TrimRight(*inp, "'"))
	inp = tools.ToStringPointer(strings.Replace(*inp, "\"", " ", -1))
	inp = tools.ToStringPointer(strings.Replace(*inp, "'", " ", -1))
	inp = tools.ToStringPointer(strings.TrimSpace(*inp))
	if toUpperv {
		inp = tools.ToStringPointer(strings.ToUpper(*inp))
	}
	if toLowerv {
		inp = tools.ToStringPointer(strings.ToLower(*inp))
	}
}

type filterf func(interface{}) bool

func filterFirst(in interface{}, fn filterf) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		current := val.Index(i).Interface()

		if fn(current) {
			out = append(out, current)
			break
		}
	}

	return out
}

// Services
// CancelledBillList implements
func (s *serverCollection) CancelledBillList(ctx context.Context, in *pbMessages.CancelledBillListRequest) (rsp *pbMessages.CancelledBillListResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at CancelledBillList %v", string(debug.Stack())))
		}
	}()
	log.Println(".... CancelledBillList ....")
	Data, err := cancelledBillListP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// GetPayment implements
func (s *serverCollection) GetPayment(ctx context.Context, in *pbMessages.GetPaymentRequest) (rsp *pbMessages.GetPaymentResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetPayment %v", string(debug.Stack())))
		}
	}()
	log.Println(".... GetPayment ....")
	Data, err := getPaymentP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// GetCustomerPayments implements
func (s *serverCollection) GetCustomerPayments(ctx context.Context, in *pbMessages.GetCustomerPaymentsRequest) (rsp *pbMessages.GetCustomerPaymentsResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetCustomerPayments %v", string(debug.Stack())))
		}
	}()
	log.Println(".... GetCustomerPayments ....")
	Data, err := getCustomerPaymentsP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// CancelledBillRequest implements
func (s *serverCollection) CancelledBillRequest(ctx context.Context, in *pbMessages.CancelledBillRequestRequest) (rsp *pbMessages.CancelledBillRequestResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at CancelledBillRequest %v", string(debug.Stack())))
		}
	}()
	log.Println(".... CancelledBillRequest ....")
	Data, err := cancelledBillRequestP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// CancelledBillAction implements
func (s *serverCollection) CancelledBillAction(ctx context.Context, in *pbMessages.CancelledBillActionRequest) (rsp *pbMessages.CancelledBillActionResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at CancelledBillAction %v", string(debug.Stack())))
		}
	}()
	log.Println(".... CancelledBillAction ....")
	Data, err := cancelledBillActionP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// BillActions implements
func (s *serverCollection) BillActions(ctx context.Context, in *pbMessages.Empty) (rsp *pbMessages.BillActionsResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at BillActions %v", string(debug.Stack())))
		}
	}()
	log.Println(".... BillActions ....")
	Data, err := billActionsP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// BillStates implements
func (s *serverCollection) BillStates(ctx context.Context, in *pbMessages.Empty) (rsp *pbMessages.BillStatesResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at BillStates %v", string(debug.Stack())))
		}
	}()
	log.Println(".... BillStates ....")
	Data, err := billStatesP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// SaveBillCancelRequest implements
func (s *serverCollection) SaveBillCancelRequest(ctx context.Context, in *pbMessages.SaveBillCancelRequestRequest) (rsp *pbMessages.SaveBillCancelRequestResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at SaveBillCancelRequest %v", string(debug.Stack())))
		}
	}()
	log.Println(".... SaveBillCancelRequest ....")
	Data, err := saveBillCancelRequestP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// CancelBillsReport implements
func (s *serverCollection) CancelBillsReport(ctx context.Context, in *pbMessages.CancelBillsReportRequest) (rsp *pbMessages.CancelBillsReportResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at SaveBillCancelRequest %v", string(debug.Stack())))
		}
	}()
	log.Println(".... CancelBillsReport ....")
	Data, err := cancelBillsReportP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

// GetStations implements
func (s *serverCollection) GetStations(ctx context.Context, in *pbMessages.Empty) (rsp *pbMessages.GetStationsResponse, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at SaveBillCancelRequest %v", string(debug.Stack())))
		}
	}()
	log.Println(".... GetStations ....")
	Data, err := GetStationsP(&ctx, in)
	if err != nil {
		return Data, err
	}
	return Data, nil
}

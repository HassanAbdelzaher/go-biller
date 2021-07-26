package masservices

import (
	"context"

	pbMessages "github.com/MaisrForAdvancedSystems/go-biller-proto/go/messages"
)

func cancelledBillListP(errorstring *string, ctx *context.Context, in *pbMessages.CancelledBillListRequest) (rsp *pbMessages.CancelledBillListResponse, err error) {
	perfectreq := false
	defer panicce(errorstring, &perfectreq)
	return cancelledBillListPP(ctx, in, &perfectreq)
}
func getCustomerPaymentsP(errorstring *string, ctx *context.Context, in *pbMessages.GetCustomerPaymentsRequest) (rsp *pbMessages.GetCustomerPaymentsResponse, err error) {
	perfectreq := false
	defer panicce(errorstring, &perfectreq)
	return getCustomerPaymentsPP(ctx, in, &perfectreq)
}
func cancelledBillRequestP(errorstring *string, ctx *context.Context, in *pbMessages.CancelledBillRequestRequest) (rsp *pbMessages.CancelledBillRequestResponse, err error) {
	perfectreq := false
	defer panicce(errorstring, &perfectreq)
	return cancelledBillRequestPP(ctx, in, &perfectreq)
}
func cancelledBillActionP(errorstring *string, ctx *context.Context, in *pbMessages.CancelledBillActionRequest) (rsp *pbMessages.CancelledBillActionResponse, err error) {
	perfectreq := false
	defer panicce(errorstring, &perfectreq)
	return cancelledBillActionPP(ctx, in, &perfectreq)
}
func billActionsP(errorstring *string, ctx *context.Context, in *pbMessages.Empty) (rsp *pbMessages.BillActionsResponse, err error) {
	perfectreq := false
	defer panicce(errorstring, &perfectreq)
	return billActionsPP(ctx, in, &perfectreq)
}
func billStatesP(errorstring *string, ctx *context.Context, in *pbMessages.Empty) (rsp *pbMessages.BillStatesResponse, err error) {
	perfectreq := false
	defer panicce(errorstring, &perfectreq)
	return billStatesPP(ctx, in, &perfectreq)
}
func saveBillCancelRequestP(errorstring *string, ctx *context.Context, in *pbMessages.SaveBillCancelRequestRequest) (rsp *pbMessages.SaveBillCancelRequestResponse, err error) {
	perfectreq := false
	defer panicce(errorstring, &perfectreq)
	return saveBillCancelRequestPP(ctx, in, &perfectreq)
}

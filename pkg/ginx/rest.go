package ginx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	httpCodeClientClosed = 499
	headerRequestID      = "X-Request-Id"
)

var (
	CodeClientClosed        interface{} = 499
	CodeInternalServerError interface{} = 500
	CodeBadRequest          interface{} = 400
)

var (
	internalErrMeta     = ErrorMeta{HTTPStatus: http.StatusInternalServerError, Code: CodeInternalServerError, Message: "Internal server error"}
	clientClosedErrMeta = ErrorMeta{HTTPStatus: httpCodeClientClosed, Code: CodeClientClosed}
	badRequestErrMeta   = ErrorMeta{HTTPStatus: http.StatusBadRequest, Code: CodeBadRequest}
	defaultErrsMaps     = make(ErrMap)
)

type ErrorMeta struct {
	HTTPStatus int
	Code       interface{}
	Message    string
}

func (m ErrorMeta) ToResponseMeta() ResponseMeta {
	return ResponseMeta{
		Code:    m.Code,
		Message: m.Message,
	}
}

type ErrMap map[error]ErrorMeta

func RegisterDefaultErrorsMap(errsMap ErrMap) {
	defaultErrsMaps = errsMap
}

func BuildSuccessResponse(ctx *gin.Context, status int, body interface{}) {
	BuildStandardResponse(ctx, status, body, ResponseMeta{Code: int64(status), RequestID: ctx.GetHeader(headerRequestID)})
}

func BuildStandardResponse(ctx *gin.Context, status int, body interface{}, meta interface{}) {
	ctx.JSON(status, response{Data: body, Meta: meta})
}

func BuildErrorResponse(ctx *gin.Context, err error, body interface{}) {
	BuildErrorResponseWithErrorsMaps(ctx, err, body)
}

func BuildErrorResponseWithErrorsMaps(ctx *gin.Context, err error, body interface{}, errMapsList ...ErrMap) {
	errMapsList = append(errMapsList, defaultErrsMaps)
	rootErr := errors.Cause(err)

	httpStatus, respMeta := calculateStatusAndResponseMeta(ctx, rootErr, errMapsList)
	respMeta.RequestID = ctx.GetHeader(headerRequestID)

	ctx.JSON(httpStatus, response{
		Data: body,
		Meta: respMeta,
	})
}

func calculateStatusAndResponseMeta(ctx *gin.Context, rootErr error, errMapsList []ErrMap) (int, ResponseMeta) {
	err := rootErr

	meta := findErrMeta(err, errMapsList)
	if meta == nil {
		if isCanceledRequest(ctx.Request) && errors.Is(err, context.Canceled) {
			return clientClosedErrMeta.HTTPStatus, clientClosedErrMeta.ToResponseMeta()
		}
		return badRequestErrMeta.HTTPStatus, badRequestErrMeta.ToResponseMeta()
	}

	return meta.HTTPStatus, meta.ToResponseMeta()
}

func isCanceledRequest(request *http.Request) bool {
	return request.Context().Err() == context.Canceled
}

func findErrMeta(err error, errsMapList []ErrMap) *ErrorMeta {
	for _, errsMap := range errsMapList {
		if e, ok := errsMap[err]; ok {
			return &e
		}
	}
	return nil
}

type response struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}

type ResponseMeta struct {
	Code      interface{} `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	RequestID string      `json:"request_id"`
}

package helper

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	error_type_unauthenticated       = "unauthenticated"
	error_type_duplicate_auth        = "duplicate_auth"
	error_type_permission            = "permission"
	error_type_general               = "general"
	error_type_invalid_json_request  = "invalid_json_request"
	error_type_flow                  = "flow"
	error_type_flow_data_not_found   = "flow_data_not_found"
	error_type_flow_create           = "flow_create"
	error_type_flow_account_blocked  = "flow_account_blocked"
	error_type_already_login         = "already_login"
	error_type_flow_update           = "flow_update"
	error_type_flow_delete           = "flow_delete"
	error_type_flow_read             = "flow_read"
	error_type_token                 = "token"
	error_type_email_already_used    = "email_already_used"
	error_type_user_not_active       = "user_not_active"
	error_type_username_already_used = "username_already_used"
	error_type_login_in_same_browser = "auth_error"

	ERR_GENERAL_400     = "kesalahan dalam sistem"
	ERR_REQUESTBODY_400 = "kesalahan pada request"

	DATANOTFOUND_400    = "data tidak ditemukan"
	CONTENTNOTFOUND_404 = "konten tidak ditemukan"
	WRONGPASSWORD_400   = "password salah"

	CREATESUCCESS_200   = "sukses menambahkan data"
	CREATEFAILED_400    = "gagal menambahkan data"
	CREATEDUPLICATE_400 = "data sudah ada"

	READSUCCESS_200 = "sukses menampilkan data"
	READFAILED_400  = "gagal menampilkan data"

	UPDATESUCCESS_200 = "sukses mengubah data"
	UPDATEFAILED_400  = "gagal mengubah data"

	DELETESUCCESS_200 = "sukses menghapus data"
	DELETEFAILED_400  = "gagal menghapus data"

	UPLOADSUCCESS_200 = "sukses mengunggah data"
	UPLOADFAILED_400  = "gagal mengunggah data"

	DOWNLOADSUCCESS_200 = "sukses mengunduh data"
	DOWNLOADFAILED_400  = "gagal mengunduh data"

	AUTHISLOGIN_400 = "user sedang login di device lain"

	AUTHISLOGOUT_400 = "user sudah logout"
)

func generateErrorResponse(c *gin.Context, flag string, message string, http_code int) {
	response := BaseErrorResponse{
		Code:    http_code,
		Message: message,
		Error:   message,
	}

	c.JSON(http_code,
		response)
	c.Abort()
}

func GenerateUnAuthorizedResponse(c *gin.Context, message string) {
	generateErrorResponse(c, error_type_general, message, http.StatusUnauthorized)
}

func GenerateTokenEmptyResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_token, "token is empty", http.StatusUnauthorized)
}

func GenerateDuplicateAuthErrorResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_duplicate_auth, AUTHISLOGIN_400, http.StatusUnauthorized)
}

func GenerateAuthLogoutErrorResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_unauthenticated, AUTHISLOGOUT_400, http.StatusUnauthorized)
}

func GeneratePermissionErrorResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_permission, "user tidak punya permission", http.StatusForbidden)
}

func GenerateInvalidJsonResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_invalid_json_request, ERR_REQUESTBODY_400+". "+err.Error(), http.StatusBadRequest)
}

func GenerateFlowErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_flow, err.Error(), http.StatusBadRequest)
}

func GenerateFlowErrorFromMessageResponse(c *gin.Context, message string) {
	err := errors.New(message)
	generateErrorResponse(c, error_type_flow, err.Error(), http.StatusBadRequest)
}

func GenerateDataNotFoundErrorResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_flow_data_not_found, DATANOTFOUND_400, http.StatusNotFound)
}

func GenerateAccountBlockedErrorResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_flow_account_blocked, "The Account has been blocked by system, please contact administrator", http.StatusUnauthorized)
}

func GenerateAlreadyLoginErrorResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_already_login, "The Credentials Already Login in Same Browser", http.StatusInternalServerError)
}

func GenerateInsertErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_flow_create, CREATEFAILED_400+", "+err.Error(), http.StatusBadRequest)
}

func GenerateUpdateErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_flow_update, UPDATEFAILED_400+", "+err.Error(), http.StatusBadRequest)
}

func GenerateTokenErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_token, err.Error(), http.StatusUnauthorized)
}

func GenerateDeleteErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_flow_delete, DELETEFAILED_400+", "+err.Error(), http.StatusBadRequest)
}

func GenerateReadErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_flow_read, READFAILED_400+", "+err.Error(), http.StatusBadRequest)
}

func GenerateEmailAlreadyUsedResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_email_already_used, "Email telah digunakan", http.StatusBadRequest)
}

func GenerateUsernameAlreadyUsedResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_username_already_used, "Username telah digunakan", http.StatusBadRequest)
}

func GenerateContentNotFoundErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_flow_data_not_found, CONTENTNOTFOUND_404+", "+err.Error(), http.StatusNotFound)
}

func GenerateCustomContentNotFoundErrorResponse(c *gin.Context, err error) {
	generateErrorResponse(c, error_type_flow_data_not_found, err.Error(), http.StatusNotFound)
}

func GenerateUserNotActiveErrorResponse(c *gin.Context) {
	generateErrorResponse(c, error_type_user_not_active, "User Sedang Dalam Kondisi Tidak Aktif, Silahkan Hubungi Admin untuk mengaktifkan User ", http.StatusUnauthorized)
}

func GeneralStatusInternalServerErrorResponse(c *gin.Context, mess string) {
	generateErrorResponse(c, error_type_login_in_same_browser, mess, http.StatusInternalServerError)
}

func GeneralUnauthorizedErrorResponse(c *gin.Context, mess string) {
	generateErrorResponse(c, error_type_permission, mess, http.StatusUnauthorized)
}

type BaseErrorResponse struct {
	Code    int
	Message string
	Error   string
}

type Body struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data"`
	Pagination *pagination `json:"pagination,omitempty"`
}

type pagination struct {
	TotalData        interface{} `json:"total_data"`
	TotalDataPerPage interface{} `json:"total_data_perpage"`
	Limit            interface{} `json:"limit"`
	TotalPage        interface{} `json:"total_page"`
	Page             interface{} `json:"page"`
}

func FormatResponse(message string, err error, data ...interface{}) (statusCode int, b *Body) {
	var (
		msg string
		d   interface{}

		pg = pagination{}

		code int
	)

	switch message {
	case DATANOTFOUND_400:
		code = http.StatusBadRequest
		msg = DATANOTFOUND_400
	case CREATESUCCESS_200:
		code = http.StatusOK
		msg = CREATESUCCESS_200
	case CREATEFAILED_400:
		code = http.StatusBadRequest
		msg = CREATEFAILED_400
	case CREATEDUPLICATE_400:
		code = http.StatusBadRequest
		msg = CREATEDUPLICATE_400
	case READSUCCESS_200:
		code = http.StatusOK
		msg = READSUCCESS_200
	case CONTENTNOTFOUND_404:
		code = http.StatusNotFound
		msg = CONTENTNOTFOUND_404
	case READFAILED_400:
		code = http.StatusBadRequest
		msg = READFAILED_400
	case UPDATESUCCESS_200:
		code = http.StatusOK
		msg = UPDATESUCCESS_200
	case UPDATEFAILED_400:
		code = http.StatusBadRequest
		msg = UPDATEFAILED_400
	case DELETESUCCESS_200:
		code = http.StatusOK
		msg = DELETESUCCESS_200
	case DELETEFAILED_400:
		code = http.StatusBadRequest
		msg = DELETEFAILED_400
	case UPLOADSUCCESS_200:
		code = http.StatusOK
		msg = UPLOADSUCCESS_200
	case UPLOADFAILED_400:
		code = http.StatusBadRequest
		msg = UPLOADFAILED_400
	case DOWNLOADSUCCESS_200:
		code = http.StatusOK
		msg = DOWNLOADSUCCESS_200
	case DOWNLOADFAILED_400:
		code = http.StatusBadRequest
		msg = DOWNLOADFAILED_400
	default:
		code = http.StatusInternalServerError
		msg = ERR_GENERAL_400
	}

	if err != nil {
		log.Println(err)
	}

	if len(data) >= 1 {
		d = data[0]
	}

	b = &Body{
		Code:    code,
		Data:    d,
		Message: msg,
	}

	if err != nil {
		b.Error = err.Error()
	}

	if len(data) > 1 {
		pg.TotalData = data[1]
		pg.Page = data[2]
		pg.Limit = data[3]
		pg.TotalDataPerPage = data[4]
		pg.TotalPage = data[5]

		b.Pagination = &pg
	}

	statusCode = code

	return
}

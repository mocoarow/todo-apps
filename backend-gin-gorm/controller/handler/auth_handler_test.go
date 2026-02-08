package handler_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/handler"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

var testCookieConfig = &controller.CookieConfig{
	Name:                "access_token",
	Path:                "/",
	Secure:              false,
	SameSite:            "Lax",
	RefreshThresholdMin: 30,
}

func initAuthRouter(t *testing.T, ctx context.Context, authUsecase handler.AuthUsecase) *gin.Engine {
	t.Helper()

	router, err := handler.InitRootRouterGroup(ctx, config, domain.AppName)
	require.NoError(t, err)
	api := router.Group("api")
	v1 := api.Group("v1")

	initAuthRouterFunc := handler.NewInitAuthRouterFunc(authUsecase, testCookieConfig, 60)
	initAuthRouterFunc(v1)

	return router
}

func Test_AuthHandler_Authenticate_shouldReturn200_whenValidCredentials(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	output, err := domain.NewAuthenticateOutput("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.sig")
	require.NoError(t, err)
	authUsecase.EXPECT().Authenticate(mock.Anything).Return(output, nil).Once()
	r := initAuthRouter(t, ctx, authUsecase)
	w := httptest.NewRecorder()
	body := `{"loginId":"user1","password":"password1"}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	jsonObj := parseJSON(t, respBytes)
	accessTokenExpr := parseExpr(t, "$.accessToken")
	accessToken := accessTokenExpr.Get(jsonObj)
	require.Len(t, accessToken, 1)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.sig", accessToken[0])
}

func Test_AuthHandler_Authenticate_shouldReturn400_whenRequestBodyIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	r := initAuthRouter(t, ctx, authUsecase)
	w := httptest.NewRecorder()
	body := `{"loginId":"","password":""}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusBadRequest, w.Code)
	validateErrorResponse(t, respBytes, "invalid_authenticate_request", "request body is invalid")
}

func Test_AuthHandler_Authenticate_shouldReturn401_whenUsecaseReturnsErrUnauthenticated(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	authUsecase.EXPECT().Authenticate(mock.Anything).Return(nil, fmt.Errorf("authenticate user: %w", domain.ErrUnauthenticated)).Once()
	r := initAuthRouter(t, ctx, authUsecase)
	w := httptest.NewRecorder()
	body := `{"loginId":"user1","password":"password1"}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	validateErrorResponse(t, respBytes, "unauthenticated", "Unauthorized")
}

func Test_AuthHandler_Authenticate_shouldReturn500_whenUsecaseReturnsUnexpectedError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	authUsecase.EXPECT().Authenticate(mock.Anything).Return(nil, fmt.Errorf("unexpected error")).Once()
	r := initAuthRouter(t, ctx, authUsecase)
	w := httptest.NewRecorder()
	body := `{"loginId":"user1","password":"password1"}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	validateErrorResponse(t, respBytes, "internal_server_error", "Internal Server Error")
}

func Test_AuthHandler_Authenticate_shouldSetCookie_whenXTokenDeliveryIsCookie(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	output, err := domain.NewAuthenticateOutput("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.sig")
	require.NoError(t, err)
	authUsecase.EXPECT().Authenticate(mock.Anything).Return(output, nil).Once()
	r := initAuthRouter(t, ctx, authUsecase)
	w := httptest.NewRecorder()
	body := `{"loginId":"user1","password":"password1"}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token-Delivery", "cookie")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	// verify cookie is set
	cookies := w.Result().Cookies()
	var accessTokenCookie *http.Cookie
	for _, c := range cookies {
		if c.Name == "access_token" {
			accessTokenCookie = c
			break
		}
	}
	require.NotNil(t, accessTokenCookie, "access_token cookie should be set")
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.sig", accessTokenCookie.Value)
	assert.True(t, accessTokenCookie.HttpOnly)
	assert.Equal(t, "/", accessTokenCookie.Path)
	assert.Equal(t, 3600, accessTokenCookie.MaxAge)

	// verify response body does NOT contain accessToken
	jsonObj := parseJSON(t, respBytes)
	accessTokenExpr := parseExpr(t, "$.accessToken")
	accessToken := accessTokenExpr.Get(jsonObj)
	assert.Empty(t, accessToken)
}

func Test_AuthHandler_Authenticate_shouldReturn400_whenXTokenDeliveryIsInvalid(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	// No Authenticate call expected because validation rejects the header first
	r := initAuthRouter(t, ctx, authUsecase)
	w := httptest.NewRecorder()
	body := `{"loginId":"user1","password":"password1"}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token-Delivery", "invalid-value")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusBadRequest, w.Code)
	validateErrorResponse(t, respBytes, "invalid_token_delivery", "X-Token-Delivery must be 'json' or 'cookie'")
}

func Test_AuthHandler_Authenticate_shouldReturnTokenInBody_whenXTokenDeliveryIsJson(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	output, err := domain.NewAuthenticateOutput("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.sig")
	require.NoError(t, err)
	authUsecase.EXPECT().Authenticate(mock.Anything).Return(output, nil).Once()
	r := initAuthRouter(t, ctx, authUsecase)
	w := httptest.NewRecorder()
	body := `{"loginId":"user1","password":"password1"}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token-Delivery", "json")
	r.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	// verify no cookie is set
	cookies := w.Result().Cookies()
	for _, c := range cookies {
		assert.NotEqual(t, "access_token", c.Name)
	}

	// verify accessToken in body
	jsonObj := parseJSON(t, respBytes)
	accessTokenExpr := parseExpr(t, "$.accessToken")
	accessToken := accessTokenExpr.Get(jsonObj)
	require.Len(t, accessToken, 1)
	assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.sig", accessToken[0])
}

func Test_AuthHandler_Authenticate_shouldReturn500_whenCookieConfigIsNilAndXTokenDeliveryIsCookie(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	authUsecase := NewMockAuthUsecase(t)
	output, err := domain.NewAuthenticateOutput("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.sig")
	require.NoError(t, err)
	authUsecase.EXPECT().Authenticate(mock.Anything).Return(output, nil).Once()

	router, err := handler.InitRootRouterGroup(ctx, config, domain.AppName)
	require.NoError(t, err)
	v1 := router.Group("api").Group("v1")
	initAuthRouterFunc := handler.NewInitAuthRouterFunc(authUsecase, nil, 60)
	initAuthRouterFunc(v1)

	w := httptest.NewRecorder()
	body := `{"loginId":"user1","password":"password1"}`

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/api/v1/auth/authenticate", strings.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Token-Delivery", "cookie")
	router.ServeHTTP(w, req)
	respBytes := readBytes(t, w.Body)

	// then
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	validateErrorResponse(t, respBytes, "cookie_not_configured", "cookie delivery is not configured")
}

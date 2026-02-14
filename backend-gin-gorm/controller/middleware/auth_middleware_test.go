package middleware_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/middleware"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

var testCookieConfig = &controller.CookieConfig{
	Name:                "access_token",
	Path:                "/",
	Secure:              false,
	SameSite:            "Lax",
	RefreshThresholdMin: 30,
}

func setupRouter(t *testing.T, authUsecase middleware.AuthUsecase) *gin.Engine {
	t.Helper()
	r := gin.New()
	r.Use(middleware.NewAuthMiddleware(authUsecase, testCookieConfig, 60))
	r.GET("/protected", func(c *gin.Context) {
		userID := c.GetInt(controller.ContextFieldUserID{})
		c.JSON(http.StatusOK, gin.H{"userId": userID})
	})
	return r
}

func Test_AuthMiddleware_shouldReturn200AndSetUserID_whenValidBearerToken(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userInfo, err := domain.NewUserInfo(42, "user42", time.Now().Add(60*time.Minute))
	require.NoError(t, err)
	output, err := domain.NewGetUserInfoOutput(userInfo)
	require.NoError(t, err)

	mockUsecase := NewMockAuthUsecase(t)
	mockUsecase.EXPECT().GetUserInfo(mock.Anything).Return(output, nil).Once()
	r := setupRouter(t, mockUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/protected", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer valid-token")
	r.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"userId":42`)
}

func Test_AuthMiddleware_shouldReturn401_whenAuthorizationHeaderIsMissing(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	mockUsecase := NewMockAuthUsecase(t)
	r := setupRouter(t, mockUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/protected", nil)
	require.NoError(t, err)
	r.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_AuthMiddleware_shouldReturn401_whenAuthorizationHeaderIsNotBearer(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	// given
	mockUsecase := NewMockAuthUsecase(t)
	r := setupRouter(t, mockUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/protected", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Basic dXNlcjpwYXNz")
	r.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_AuthMiddleware_shouldReturn401_whenGetUserInfoReturnsError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	// given
	mockUsecase := NewMockAuthUsecase(t)
	mockUsecase.EXPECT().GetUserInfo(mock.Anything).Return(nil, errors.New("invalid token")).Once()
	r := setupRouter(t, mockUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/protected", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer invalid-token")
	r.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_AuthMiddleware_shouldReturn200_whenTokenIsInCookie(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userInfo, err := domain.NewUserInfo(42, "user42", time.Now().Add(60*time.Minute))
	require.NoError(t, err)
	output, err := domain.NewGetUserInfoOutput(userInfo)
	require.NoError(t, err)
	refreshOutput := domain.NewRefreshTokenOutput("")

	mockUsecase := NewMockAuthUsecase(t)
	mockUsecase.EXPECT().GetUserInfo(mock.Anything).Return(output, nil).Once()
	mockUsecase.EXPECT().RefreshToken(mock.Anything).Return(refreshOutput, nil).Once()
	r := setupRouter(t, mockUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/protected", nil)
	require.NoError(t, err)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: "cookie-token"})
	r.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"userId":42`)
}

func Test_AuthMiddleware_shouldSetNewCookie_whenSlidingRefreshNeeded(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userInfo, err := domain.NewUserInfo(42, "user42", time.Now().Add(60*time.Minute))
	require.NoError(t, err)
	output, err := domain.NewGetUserInfoOutput(userInfo)
	require.NoError(t, err)
	refreshOutput := domain.NewRefreshTokenOutput("new-refreshed-token")

	mockUsecase := NewMockAuthUsecase(t)
	mockUsecase.EXPECT().GetUserInfo(mock.Anything).Return(output, nil).Once()
	mockUsecase.EXPECT().RefreshToken(mock.Anything).Return(refreshOutput, nil).Once()
	r := setupRouter(t, mockUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/protected", nil)
	require.NoError(t, err)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: "old-cookie-token"})
	r.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	// verify new cookie is set
	cookies := w.Result().Cookies()
	var accessTokenCookie *http.Cookie
	for _, c := range cookies {
		if c.Name == "access_token" {
			accessTokenCookie = c
			break
		}
	}
	require.NotNil(t, accessTokenCookie, "access_token cookie should be set with refreshed token")
	assert.Equal(t, "new-refreshed-token", accessTokenCookie.Value)
	assert.True(t, accessTokenCookie.HttpOnly)
	assert.Equal(t, 3600, accessTokenCookie.MaxAge)
}

func Test_AuthMiddleware_shouldPreferBearerOverCookie(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// given
	userInfo, err := domain.NewUserInfo(42, "user42", time.Now().Add(60*time.Minute))
	require.NoError(t, err)
	output, err := domain.NewGetUserInfoOutput(userInfo)
	require.NoError(t, err)

	mockUsecase := NewMockAuthUsecase(t)
	mockUsecase.EXPECT().GetUserInfo(mock.Anything).Return(output, nil).Once()
	// No RefreshToken call expected because Bearer takes priority (not from cookie)
	r := setupRouter(t, mockUsecase)
	w := httptest.NewRecorder()

	// when
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/protected", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer bearer-token")
	req.AddCookie(&http.Cookie{Name: "access_token", Value: "cookie-token"})
	r.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	// verify no cookie refresh happened (no Set-Cookie header)
	cookies := w.Result().Cookies()
	for _, c := range cookies {
		assert.NotEqual(t, "access_token", c.Name)
	}
}

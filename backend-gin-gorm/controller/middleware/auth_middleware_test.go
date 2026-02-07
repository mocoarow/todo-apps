package middleware_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/controller/middleware"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter(t *testing.T, authUsecase middleware.AuthUsecase) *gin.Engine {
	t.Helper()
	r := gin.New()
	r.Use(middleware.NewAuthMiddleware(authUsecase))
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
	userInfo, err := domain.NewUserInfo(42, "user42")
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
	mockUsecase.EXPECT().GetUserInfo(mock.Anything).Return(nil, fmt.Errorf("invalid token")).Once()
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

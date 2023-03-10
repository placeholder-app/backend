/*
 * Produced: Thu Jan 26 2023
 * Author: Alec M., James A.
 * GitHub: https://github.com/placeholder-app
 * Copyright: (C) 2023 Alec M., James A.
 * License: License GNU Affero General Public License v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"placeholder-app/backend/middlewares"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestCorsHeadMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("HEAD", "/", nil)

	middlewares.Cors()(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "HEAD, OPTIONS, GET", w.Header().Get("Access-Control-Allow-Methods"))
}

func TestCorsOptionsMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("OPTIONS", "/", nil)

	middlewares.Cors()(c)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "HEAD, OPTIONS, GET", w.Header().Get("Access-Control-Allow-Methods"))
}

func TestCorsGetMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/", nil)

	middlewares.Cors()(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "HEAD, OPTIONS, GET", w.Header().Get("Access-Control-Allow-Methods"))
}

func TestCorsPostMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/", nil)

	middlewares.Cors()(c)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestCorsDeleteMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("DELETE", "/", nil)

	middlewares.Cors()(c)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestCorsPutMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("PUT", "/", nil)

	middlewares.Cors()(c)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestCorsPatchMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("PATCH", "/", nil)

	middlewares.Cors()(c)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

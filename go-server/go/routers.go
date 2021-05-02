/*
 * Salamander api server
 *
 * Salamander is the backend api server for HermitCrab.
 *
 * API version: 1.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/",
		Index,
	},

	Route{
		"SignIn",
		strings.ToUpper("Post"),
		"/api/v1/auth/signin",
		SignIn,
	},

	Route{
		"SignUp",
		strings.ToUpper("Post"),
		"/api/v1/auth/signup",
		SignUp,
	},

	Route{
		"DeleteParticipant",
		strings.ToUpper("Delete"),
		"/api/v1/participant",
		DeleteParticipant,
	},

	Route{
		"UpdateParticipantStatus",
		strings.ToUpper("Patch"),
		"/api/v1/participant/status",
		UpdateParticipantStatus,
	},

	Route{
		"CreateRoom",
		strings.ToUpper("Post"),
		"/api/v1/rooms",
		CreateRoom,
	},

	Route{
		"GetMenuInfo",
		strings.ToUpper("Get"),
		"/api/v1/rooms",
		GetMenuInfo,
	},

	Route{
		"GetRoom",
		strings.ToUpper("Get"),
		"/api/v1/rooms/{id}",
		GetRoom,
	},

	Route{
		"UpdateRoom",
		strings.ToUpper("Patch"),
		"/api/v1/rooms/{id}",
		UpdateRoom,
	},

	Route{
		"UpdateRoomStartDate",
		strings.ToUpper("Patch"),
		"/api/v1/rooms/{id}/start",
		UpdateRoomStartDate,
	},

	Route{
		"GetServices",
		strings.ToUpper("Get"),
		"/api/v1/services",
		GetServices,
	},

	Route{
		"GetServicesPlan",
		strings.ToUpper("Get"),
		"/api/v1/services/{id}",
		GetServicesPlan,
	},

	Route{
		"CreateRoom",
		strings.ToUpper("Patch"),
		"/api/v1/user",
		CreateRoom,
	},

	Route{
		"GetUser",
		strings.ToUpper("Get"),
		"/api/v1/users/{id}",
		GetUser,
	},

	Route{
		"GetUsers",
		strings.ToUpper("Get"),
		"/api/v1/users",
		GetUsers,
	},
}

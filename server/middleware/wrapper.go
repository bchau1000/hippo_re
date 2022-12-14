package middleware

import "net/http"

/**
 * Utility to make chaining middleware easier/cleaner
 *
 * Chaining middleware by default looks ugly...
 *	ex. middleWareOne(middleWareTwo(controller()))
 *
 * This function takes a list of type Middleware and chronologically wraps
 * the controller in said list of Middlewares
 * ex.
 *	Run(controller(), middleWareOne(), middleWareTwo())
 *  is equivalent to: middleWareOne(middleWareTwo(controller()))
 *
 * Same amount of code in this case, but consider how middlewares are normally used...
 * ex.
 *	baseMiddleWares = []Middleware{AuthMiddleWare, LogMiddleWare, PerformanceMiddleWare}
 *  Run(controller1(), baseMiddleWares)
 *  Run(controller2(), baseMiddleWares)
**/

type IMiddlewareWrapper interface {
	Wrap() MiddlewareWrapper
}

type MiddlewareWrapper func(http.HandlerFunc) http.HandlerFunc

func Wrap(handler http.HandlerFunc, wrapper ...IMiddlewareWrapper) http.HandlerFunc {

	var length int = len(wrapper)

	if length < 1 {
		return handler
	}

	wrappedHandler := handler

	for i := length - 1; i >= 0; i-- {
		wrappedHandler = wrapper[i].Wrap()(wrappedHandler)
	}

	return wrappedHandler
}

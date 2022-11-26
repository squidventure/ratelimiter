# package ratelimiter
ratelimiter offers a simple RateLimiter interface and a BasicRateLimiter type to limit the number of incoming connections to an API. it includes a simple way to bypass the rate limiter for certain endpoints.

## usage
wrap your http.Server's Handler with the ratelimiter.Limit() middleware to add to your project. Set the MaxConnections as appropriate. To exclude a path prefix or suffix from the rate limiter use the appropriate RegisterBypassPath... function.

## changing the rate limiter style
package ratelimiter declares a global RateLimiter `TheRateLimiter` which is used in the ratelimiter.Arrive() function. This could be changed to any other type of RateLimiter interface. 
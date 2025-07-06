## Rate Limiter written in golang


### Token Bucket Algorithm

we will use the token bucket algorithm

* We will have a bucket with fixed number of tokens
* Each request removes one token from the request
* If bucket is empty, any further requests are dropped
* Tokens are replenished at a steady interval


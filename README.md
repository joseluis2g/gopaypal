# gopaypal

Go client for the PayPal REST API

# Usage

First we create a gopaypal client

```go
client := NewClient(clientID, secretKey, URL)
```

Then we can start by getting PayPal OAuth2 token

```go
tokenResponse, err := client.GetAccessToken()
```

You dont really need to get the token response. The client will save your token and update it when needed. After we have the token we can start calling the PayPal REST API. For more information check the tests

# Missing endpoints

You can still use gopaypal even if the endpoint you look for is missing. Create a client and use `AuthRequest` or `BasicRequest`

# Testing

Testing **gopaypal** involves the following mandatory command line arguments:

- `clientid` Your application client identifier
- `secret` Your application secret key
- `redirect` Your application redirect URL. Must match your redirect URL of your PayPal application

The following arguments are optional:

- `token` PayPal OAuth2 token. You can pass this argument to skip the OAuth test
- `iden` Identifier code given by PayPal after a successful login using the identity API. You should pass this argument in order to test several identity API functions
- `paymentid` Payment identifier to test the execution of an approved PayPal payment
- `payerid` Payer identifier to test the execution of an approved PayPal payment

# License

gopaypal is made available under the MIT license
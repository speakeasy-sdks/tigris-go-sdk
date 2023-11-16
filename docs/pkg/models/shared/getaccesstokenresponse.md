# GetAccessTokenResponse

The response of GetAccessToken which contains access_token and optionally refresh_token.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `AccessToken`                               | **string*                                   | :heavy_minus_sign:                          | An Access Token.                            |
| `ExpiresIn`                                 | **int*                                      | :heavy_minus_sign:                          | Access token expiration timeout in seconds. |
| `RefreshToken`                              | **string*                                   | :heavy_minus_sign:                          | The Refresh Token.                          |
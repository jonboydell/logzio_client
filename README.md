# Logz.io Terraform client library

Client library for Logz.io API, see below for supported endpoints.

The primary purpose of this library is to act as the API interface for the logz.io Terraform provider.
To use it, you'll need to [create an API token](https://app.logz.io/#/dashboard/settings/api-tokens) and provide it to the client library along with your logz.io regional [API server address](https://docs.logz.io/user-guide/accounts/account-region.html#regions-and-urls).

##### Usage

Note: the lastest version of the API (1.3) is not backwards compatible with previous versions, specifically the client entrypoint names have changed to prevent naming conflicts. Use `UsersClient` ([Users API](#users)) ,`SubaccountClient` ([Sub-accounts API](#sub-accounts)), `AlertsClient` ([Alerts API](#alerts)) and `EndpointsClient` ([Endpoints API](#endpoints)) rather than `Users`, `Alerts` and `Endpoints`.

##### Alerts V2
To create an alert where the type field = 'mytype' and the loglevel field = ERROR, see the logz.io docs for more info

https://support.logz.io/hc/en-us/articles/209487329-How-do-I-create-an-Alert-

```go
client, _ := alerts_v2.New(apiToken, apiServerAddress)
alertQuery := alerts_v2.AlertQuery{
		Query:                    "loglevel:ERROR",
		Aggregation:              alerts_v2.AggregationObj{AggregationType: alerts_v2.AggregationTypeCount},
		ShouldQueryOnAllAccounts: true,
	}

	trigger := alerts_v2.AlertTrigger{
		Operator:               alerts_v2.OperatorEquals,
		SeverityThresholdTiers: map[string]float32{alerts_v2.SeverityHigh: 10, alerts_v2.SeverityInfo: 5},
	}
	
	subComponent := alerts_v2.SubAlert{
		QueryDefinition: alertQuery,
		Trigger:         trigger,
		Output:          alerts_v2.SubAlertOutput{},
	}

	createAlertType := alerts_v2.CreateAlertType{
		Title:                  "test create alert",
		Description:            "this is my description",
		Tags:                   []string{"some", "words"},
		Output:                 alerts_v2.AlertOutput{},
		SubComponents:          []alerts_v2.SubAlert{subComponent},
		Correlations:           alerts_v2.SubAlertCorrelation{},
		Enabled:                strconv.FormatBool(true),
	}

alert := client.CreateAlert(createAlertType)
```

|function|func name|
|---|---|
| Create alert | `func (c *AlertsV2Client) CreateAlert(alert CreateAlertType) (*AlertType, error)` |
| Delete alert | `func (c *AlertsV2Client) DeleteAlert(alertId int64) error` |
| Disable alert | `func (c *AlertsV2Client) DisableAlert(alert AlertType) (*AlertType, error)` |
| Enable alert | `func (c *AlertsV2Client) EnableAlert(alert AlertType) (*AlertType, error)` |
| Get alert | `func (c *AlertsV2Client) GetAlert(alertId int64) (*AlertType, error)` |
| List alerts | `func (c *AlertsV2Client) ListAlerts() ([]AlertType, error)` |
| Update alert | `func (c *AlertsV2Client) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error)` |

##### Alerts

To create an alert where the type field = 'mytype' and the loglevel field = ERROR, see the logz.io docs for more info

https://support.logz.io/hc/en-us/articles/209487329-How-do-I-create-an-Alert-

```go
client, _ := alerts.New(apiToken, apiServerAddress)
alert := client.CreateAlert(alerts.CreateAlertType{
    Title:       "this is my alert",
    Description: "this is my description",
    QueryString: "loglevel:ERROR",
    Filter:      "{\"bool\":{\"must\":[{\"match\":{\"type\":\"mytype\"}}],\"must_not\":[]}}",
    Operation:   alerts.OperatorGreaterThan,
    SeverityThresholdTiers: []alerts.SeverityThresholdType{
        alerts.SeverityThresholdType{
            alerts.SeverityHigh,
            10,
        },
    },
    SearchTimeFrameMinutes:       0,
    NotificationEmails:           []interface{}{},
    IsEnabled:                    true,
    SuppressNotificationsMinutes: 0,
    ValueAggregationType:         alerts.AggregationTypeCount,
    ValueAggregationField:        nil,
    GroupByAggregationFields:     []interface{}{"my_field"},
    AlertNotificationEndpoints:   []interface{}{},
})
```

|function|func name|
|---|---|
|create alert|`func (c *AlertsClient) CreateAlert(alert CreateAlertType) (*AlertType, error)`|
|update alert|`func (c *AlertsClient) UpdateAlert(alertId int64, alert CreateAlertType) (*AlertType, error)`
|delete alert|`func (c *AlertsClient) DeleteAlert(alertId int64) error`|
|get alert (by id)|`func (c *AlertsClient) GetAlert(alertId int64) (*AlertType, error)`|
|list alerts|`func (c *AlertsClient) ListAlerts() ([]AlertType, error)`|


##### Users

To create a new user, on a specific account or sub-account. you'll need [your account Id](https://docs.logz.io/user-guide/accounts/finding-your-account-id.html).

```go
client, _ := users.New(apiToken, apiServerAddress)
user := client.User{
    Username:  "createa@test.user",
    Fullname:  "my username",
    AccountId: 123456,
    Roles:     []int32{users.UserTypeUser},
}
```

|function|func name|
|---|---|
|create user|`func (c *UsersClient) CreateUser(user User) (*User, error)`|
|update user|`func (c *UsersClient) UpdateUser(user User) (*User, error)`|
|delete user|`func (c *UsersClient) DeleteUser(id int32) error`|
|get user|`func (c *UsersClient) GetUser(id int32) (*User, error)`|
|list users|`func (c *UsersClient) ListUsers() ([]User, error)`|
|suspend user|`func (c *UsersClient) SuspendUser(userId int32) (bool, error)`|
|unsuspend user|`func (c *UsersClient) UnSuspendUser(userId int32) (bool, error)`|

##### Sub-accounts

To create a new sub-account, on a main account.
```go
client, _ := sub_accounts.New(apiToken, apiServerAddress)
subaccount := sub_accounts.SubAccountCreate{
    Email:                 "test.user@email.com",
    AccountName:           "my account name",
    MaxDailyGB:            6.5,
    RetentionDays:         4,
    Searchable:            true,
    Accessible:            false,
    SharingObjectAccounts: []int32{accountId1, accountId2}, //Id's of the accounts who will be able to access this account
    DocSizeSetting:        true,
}
```

|function|func name|
|---|---|
|create sub-account|`func (c *SubAccountClient) CreateSubAccount(subAccount SubAccountCreate) (*SubAccount, error) `|
|update sub-account|`func (c *SubAccountClient) UpdateSubAccount(id int64, subAccount SubAccount) error`|
|delete sub-account|`func (c *SubAccountClient) DeleteSubAccount(id int64) error`|
|get sub-account|`func (c *SubAccountClient) GetSubAccount(id int64) (*SubAccount, error)`|
|get detailed sub-account|`func (c *SubAccountClient) GetDetailedSubAccount(id int64) (*SubAccountDetailed, error)`|
|list sub-accounts|`func (c *SubAccountClient) ListSubAccounts() ([]SubAccount, error)`|
|list detailed sub-accounts|`func (c *SubAccountClient) DetailedSubAccounts() ([]SubAccountDetailed, error) `|

##### Endpoints

For each type of endpoint there is a different structure, below you can find an example for creating a Slack endpoint.
For more info, see: https://docs.logz.io/api/#tag/Manage-notification-endpoints or check our endpoints tests for more examples.

```go
client, _ := endpoints.New(apiToken, apiServerAddress)
endpoint, err := underTest.CreateEndpoint(endpoints.Endpoint{
			Title:        "some_endpoint",
			Description:  "my description",
			Url:          "https://this.is.com/some/webhook",
			EndpointType: endpoints.EndpointTypeSlack,
		})
```

##### Log Shipping Tokens

Compatible with Logz.io's [Manage Log Shipping Tokens API](https://docs.logz.io/api/#tag/Manage-log-shipping-tokens).
To create a new log shipping token:

```go
client, _ := log_shipping_tokens.New(apiToken, apiServerAddress)
token, err := client.CreateLogShippingToken(log_shipping_tokens.CreateLogShippingToken{
                Name: "client_integration_test",
            })
```

|function|func name|
|---|---|
|create log shipping token| `func (c *LogShippingTokensClient) CreateLogShippingToken(token CreateLogShippingToken) (*LogShippingToken, error)` |
|delete log shipping token| `func (c *LogShippingTokensClient) DeleteLogShippingToken(tokenId int32) error` |
|get log shipping token| `func (c *LogShippingTokensClient) GetLogShippingToken(tokenId int32) (*LogShippingToken, error)` |
|get available number of tokens| `func (c *LogShippingTokensClient) GetLogShippingLimitsToken() (*LogShippingTokensLimits, error)` |
|retrieve tokens| `func (c *LogShippingTokensClient) RetrieveLogShippingTokens(retrieveRequest RetrieveLogShippingTokensRequest) (*RetrieveLogShippingTokensResponse, error)` |
|update log shipping token| `func (c *LogShippingTokensClient) UpdateLogShippingToken(tokenId int32, token UpdateLogShippingToken) (*LogShippingToken, error)` |

##### Drop Filters
Compatible with Logz.io's [drop filters API](https://docs.logz.io/api/#tag/Drop-filters).
Drop filters provide a solution for filtering out logs before they are indexed in your account to help lower costs and reduce account volume.
To create a new drop filter:

```go
client, _ := drop_filters.New(apiToken, apiServerAddress)
dropFilter, err := client.CreateDropFilter(drop_filters.CreateDropFilter{
                    LogType: "some_type",
                    FieldConditions: []drop_filters.FieldConditionObject{{
                        FieldName: "some_field_name",
                        Value:     "some_value",
                    }},
                })
```

|function|func name|
|---|---|
|activate drop filter| `func (c *DropFiltersClient) ActivateDropFilter(dropFilterID string) (*DropFilter, error)` |
|create drop filter| `func (c *DropFiltersClient) CreateDropFilter(createDropFilter CreateDropFilter) (*DropFilter, error)` |
|deactivate drop filter| `func (c *DropFiltersClient) DeactivateDropFilter(dropFilterId string) (*DropFilter, error)` |
|delete drop filter| `func (c *DropFiltersClient) DeleteDropFilter(dropFilterId string) error` |
|retrieve drop filters| `func (c *DropFiltersClient) RetrieveDropFilters() ([]DropFilter, error)` |

#### Contributing

1. Clone this repo locally
2. As this package uses Go modules, make sure you are outside of `$GOPATH` or you have the `GO111MODULE=on` environment variable set. Then run `go get` to pull down the dependencies.

##### Run tests
`go test -v -race ./...`


### Changelog
- v1.7.0
    - Add [drop filters API](https://docs.logz.io/api/#tag/Drop-filters).
- v1.6.0
    - Add [log shipping tokens API](https://docs.logz.io/api/#tag/Manage-log-shipping-tokens) compatibility.
- v1.5.3
    - Fix for `sub account`: return token & account id on Create. 
- v1.5.2
    - Fix `custom endpoint` -empty headers bug.
    - Allow empty array for sharing accounts in `sub account`.
- v1.5.1
    - Fix alerts_v2 sort bug.
- v1.5
    - Add alerts v2 compatibility.
- v1.3.2
   - fix client custom endpoint headers bug
   - improve tests 
- v1.3
    - unnecessary resource updates bug fix.
    - support tags in alerts
- v1.2
    - Add subaccount support



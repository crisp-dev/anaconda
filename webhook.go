package anaconda

import (
	"net/url"
)

func (a TwitterApi) GetEnvironmentWebhooks(v url.Values) (u EnvirnmentResp, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)

	a.queryQueue <- query{a.baseUrl + "/account_activity/all/webhooks.json", v, &u, _GET, responseCh}
	return u, (<-responseCh).err
}

//GetActivityWebhooks represents the twitter account_activity webhook
//Returns all URLs and their statuses for the given app. Currently,
//only one webhook URL can be registered to an application.
//https://dev.twitter.com/webhooks/reference/get/account_activity/webhooks
func (a TwitterApi) GetActivityWebhooks(v url.Values, env string) (u []WebHookResp, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)

	a.queryQueue <- query{a.baseUrl + "/account_activity/all/" + env + "/webhooks.json", v, &u, _GET, responseCh}
	return u, (<-responseCh).err
}

//EnvirnmentResp represents the Get environment responses
type EnvirnmentResp struct {
	Environments   []struct {
		EnvironmentName  string `json:"environment_name"`
		Webhooks         []WebHookResp `json:"webhooks"`
	} `json:"environments"`
}

//WebHookResp represents the webhook object
type WebHookResp struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Valid     bool   `json:"valid"`
	CreatedAt string `json:"created_at"`
}

//SetActivityWebhooks represents to set twitter account_activity webhook
//Registers a new webhook URL for the given application context.
//The URL will be validated via CRC request before saving. In case the validation fails,
//a comprehensive error is returned. message to the requester.
//Only one webhook URL can be registered to an application.
//https://api.twitter.com/1.1/account_activity/webhooks.json
func (a TwitterApi) SetActivityWebhooks(v url.Values, env string) (u WebHookResp, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/all/" + env + "/webhooks.json", v, &u, _POST, responseCh}
	return u, (<-responseCh).err
}

//DeleteActivityWebhooks Removes the webhook from the provided application’s configuration.
//https://dev.twitter.com/webhooks/reference/del/account_activity/webhooks
func (a TwitterApi) DeleteActivityWebhooks(v url.Values, webhookID string, env string) (u interface{}, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/all/" + env + "/webhooks/" + webhookID + ".json", v, &u, _DELETE, responseCh}
	return u, (<-responseCh).err
}

//PutActivityWebhooks update webhook which reenables the webhook by setting its status to valid.
//https://dev.twitter.com/webhooks/reference/put/account_activity/webhooks
func (a TwitterApi) PutActivityWebhooks(v url.Values, webhookID string, env string) (u interface{}, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/all/" + env + "/webhooks/" + webhookID + ".json", v, &u, _PUT, responseCh}
	return u, (<-responseCh).err
}

//SetWHSubscription Subscribes the provided app to events for the provided user context.
//When subscribed, all DM events for the provided user will be sent to the app’s webhook via POST request.
//https://dev.twitter.com/webhooks/reference/post/account_activity/webhooks/subscriptions
func (a TwitterApi) SetWHSubscription(v url.Values, env string) (u interface{}, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/all/" + env + "/subscriptions.json", v, &u, _POST, responseCh}
	return u, (<-responseCh).err
}

//GetWHSubscription Provides a way to determine if a webhook configuration is
//subscribed to the provided user’s Direct Messages.
//https://dev.twitter.com/webhooks/reference/get/account_activity/webhooks/subscriptions
func (a TwitterApi) GetWHSubscription(v url.Values, env string) (u interface{}, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/all/" + env + "/subscriptions.json", v, &u, _GET, responseCh}
	return u, (<-responseCh).err
}

//DeleteWHSubscription Deactivates subscription for the provided user context and app. After deactivation,
//all DM events for the requesting user will no longer be sent to the webhook URL..
//https://dev.twitter.com/webhooks/reference/del/account_activity/webhooks
func (a TwitterApi) DeleteWHSubscription(v url.Values, env string) (u interface{}, err error) {
	v = cleanValues(v)
	responseCh := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/account_activity/all/" + env + "/subscriptions.json", v, &u, _DELETE, responseCh}
	return u, (<-responseCh).err
}

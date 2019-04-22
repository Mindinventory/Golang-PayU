package main

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type PayUMoneyOrderResponse struct {
	Amount           string `json:"amount"`
	Txnid            string `json:"txnid"`
	Key              string `json:"key"`
	ProductInfo      string `json:"productinfo"`
	FirstName        string `json:"firstname"`
	Email            string `json:"email"`
	Hash             string `json:"hash"`
	Surl             string `json:"surl"`
	Phone            string `json:"phone"`
	ServiceProvider  string `json:"service_provider"`
	Furl             string `json:"furl"`
	Udf1             string `json:"udf1"`
	EnforcePaymethod string `json:"enforce_paymethod"`
}

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	var MERCHANT_KEY = os.Getenv("merchant_key")
	var SALT = os.Getenv("merchant_salt")
	var txnid = "6647747839544"
	var name = "UserName"
	var email = "xyz@gmail.com"
	var amount = "10.00"
	var phone = "0000000000"
	var surl = os.Getenv("PAYMENT_MIDDLEWARE_URL")
	var furl = os.Getenv("PAYMENT_MIDDLEWARE_URL")
	var enforcePaymethod = "creditcard"
	var productInfo = "PAYU"
	var orderId = "2010342311"

	var hashString = MERCHANT_KEY + "|" + txnid + "|" + amount + "|" + productInfo + "|" + name + "|" + email + "|" + orderId + "||||||||||" + SALT
	sha_512 := sha512.New()
	sha_512.Write([]byte(hashString))
	var final_key = fmt.Sprintf("%x", sha_512.Sum(nil))

	orderResponse := payUOrderResponseFun(amount, txnid, MERCHANT_KEY, productInfo, name, email, final_key, surl, phone, "", furl, enforcePaymethod, orderId)
	var setResponse = map[string]interface{}{}
	response := Message(0, "Order place successfully.")
	setResponse["request_data"] = orderResponse
	setResponse["redirect_url"] = os.Getenv("redirect_url")
	setResponse["callback_url"] = os.Getenv("PAYMENT_MIDDLEWARE_URL") //Get response to this URL
	response["data"] = setResponse

	fmt.Println("response", response)
}

func payUOrderResponseFun(amountString string, txnid string, merchantKey string, productInfo string, name string, email string, final_key string, surl string, phone string, serviceProvider string, furl string, enforcePaymethod string, orderId string) PayUMoneyOrderResponse {
	res := PayUMoneyOrderResponse{
		Amount:           amountString,
		Txnid:            txnid,
		Key:              merchantKey,
		ProductInfo:      productInfo,
		FirstName:        name,
		Email:            email,
		Hash:             final_key,
		Surl:             surl,
		Phone:            phone,
		ServiceProvider:  serviceProvider,
		Furl:             furl,
		EnforcePaymethod: enforcePaymethod,
		Udf1:             orderId,
	}
	return res
}

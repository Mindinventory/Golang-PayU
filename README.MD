# Quick PayU Integration using Golang

<a href="https://travis-ci.org/Mindinventory/Golang-Paytm" style="pointer-events: none;" target="_blank"><img src="https://travis-ci.org/Mindinventory/Golang-PayU.svg?branch=master"></a>
<a href="https://godoc.org/fyne.io/fyne" style="pointer-events: none;" target="_blank"><img src="https://img.shields.io/badge/go-documentation-blue.svg"></a>
<a href="https://goreportcard.com/report/github.com/Mindinventory/Golang-PayU" style="pointer-events: none;" target="_blank"><img src="https://goreportcard.com/badge/github.com/Mindinventory/Golang-PayU"></a>
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/mindinventory/Golang-PayU/blob/master/LICENSE)

<img src="https://raw.githubusercontent.com/Mindinventory/Golang-PayU/master/payU.png">

## Requirements : 
 
* i.  Provide the value for PAYU_MERCHANT_KEY in .env file. (The value for MERCHANT_KEY will be provided after the onboarding process is completed).
      
## Installation steps:

* i. PayU integration can be done in two stages (staging and production modes).
           staging and production modes have different PayU transaction urls and Merchant key.

ii. In staging mode we are using sandbox details for testing the payU working.
       
       Staging mode transaction url 
       https://test.payu.in/_payment
        
iii. If you are using production mode means it accepts original details of payU. 
  
        Production mode transaction url    
        https://secure.payu.in/_payment    
    
## LICENSE!

PayU integration using Golang is [MIT-licensed](https://github.com/mindinventory/Golang-PayU/blob/master/LICENSE)

## Let us know!
We’d be really happy if you sent us links to your projects where you use our component. Just send an email to sales@mindinventory.com And do let us know if you have any questions or suggestion regarding our work.
      

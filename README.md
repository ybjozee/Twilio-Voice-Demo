# Twilio-Voice-Demo
An application built using Go to demonstrate making programmable outbound calls via Twilio's voice API


Technical Requirements
* Go 1.19 or higher


Installation
------------
```bash
$ git clone https://github.com/ybjozee/Twilio-Voice-Demo.git
$ cd $_
```

Usage
-----

Make a local version of the `.env` file

```bash
$ cp .env .env.local
```

Update the relevant Twilio keys in `.env.local`

``` ini
TWILIO_PHONE_NUMBER="INSERT_TWILIO_PHONE_NUMBER"
TWILIO_ACCOUNT_SID="INSERT_TWILIO_ACCOUNT_SID"
TWILIO_AUTH_TOKEN="INSERT_TWILIO_AUTH_TOKEN"
```

Run your application
```bash
$ go run main.go
```

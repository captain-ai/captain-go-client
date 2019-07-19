# `captain-go-client`

Go client for `captain.ai` REST API in JSON format. 

### Testing Live Requests

If you want to test live requests, create a file with environment variables.

```
CAPTAIN_ACCOUNT_ID=
CAPTAIN_INTEGRATION_KEY=
CAPTAIN_DEVELOPER_KEY=
```

Inject the envrionment variables for testing. 

```sh
$ env $(cat .env | xargs) go test ./...
```

If you don't want to test live requests pass the `-short` testing option. 

```sh
$ go test -short ./...
```

### Example Errors

```json
{
  "errors": {
    "signature_url": [
      "must be filled"
    ],
    "items_link": [
      "must be filled"
    ],
    "recipient": {
      "full_name": [
        "must be filled"
      ],
      "phone_number": [
        "is in invalid format"
      ],
      "opted_out_of_SMS": [
        "must be filled",
        "must be boolean or must be equal to true or must be equal to false or must be equal to null"
      ],
      "opted_out_of_email": [
        "must be filled",
        "must be boolean or must be equal to true or must be equal to false or must be equal to null"
      ],
      "partners_internal_recipient_id": [
        "must be filled"
      ],
      "landline_number": [
        "must be filled"
      ]
    },
    "delivery_job": {
      "dropoff_location": {
        "longitude": [
          "must be filled"
        ],
        "latitude": [
          "must be filled"
        ]
      }
    },
    "custom_fields": [
      "must be filled"
    ]
  }
}
```
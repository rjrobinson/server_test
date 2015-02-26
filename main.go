package main

import (
	"net/http"
	"os"
)

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, _ := getJSONResponse()
	w.Write(response)

}

func serveRestError(w http.ResponseWriter, r *http.Request) {
	response, _ := getJSONResponseError()

	w.Write(response)

}

func main() {
	http.HandleFunc("/", serveRest)
	http.HandleFunc("/error", serveRestError)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

// http.HandleFunc("/", hello)
//    fmt.Println("listening...")
//    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
//    if err != nil {
//      panic(err)
//    }
// }

func getJSONResponse() ([]byte, error) {

	respBytes := []byte(`{
    "error": null,
    "paging_data": null,
    "resources": {
        "access_numbers": [
            {
                "city": "",
                "country_code": "",
                "country_iso": "",
                "date_created": "2013-05-23 16:51:23 -0400",
                "description": "",
                "did": "+15555551234",
                "sid": "123",
                "state": ""
            }
        ],
        "account": {
            "balance": 10.833335,
            "credits": 65,
            "date_created": "2014-04-18 15:43:26 -0400",
            "email_address": "JohnCarter@mars.com",
            "first_name": "John",
            "has_password": false,
            "is_active": true,
            "is_fraud": false,
            "is_paid": true,
            "last_name": "Carter",
            "max_transfer_value": 0.0,
            "plugins": {
                "auto_replenishers": [

                ]
            },
            "security_questions": [

            ],
            "sid": "e1b0956f-f03c-4eae-8f11-8a52542ea9b5",
            "total_phones": 1
        },
        "call": {
            "access_code": "9130",
            "account_sid": "e1b0956f-f03c-4eae-8f11-8a52542ea9b5",
            "calls": [
                {
                    "access_code": "1220",
                    "account_sid": "e1b0956f-f03c-4eae-8f11-8a52542ea9b5",
                    "calls": [

                    ],
                    "date_created": "2014-04-28 15:09:11 -0400",
                    "date_expires": "2014-04-28 15:11:11 -0400",
                    "date_scheduled": "2014-04-28 15:09:11 -0400",
                    "destination_address": "+17324961388",
                    "event_type": "outbound-call",
                    "owner_type": "account",
                    "parent_call_sid": "e91cd3af-2dfc-49e3-946f-6d25b46ea243",
                    "rated_event": {
                        "cost": 0,
                        "credits": 0,
                        "date_created": "2014-04-28 15:09:11 -0400",
                        "prefix": "+1",
                        "sid": "7d48ca33-d28b-424c-a8a4-750821073afc",
                        "units": 0,
                        "units_billed": 0
                    },
                    "sid": "07ca32b5-379e-4dfa-b57a-e89b53f28419",
                    "source_address": "+19082783791"
                },
                {
                    "access_code": "1220",
                    "account_sid": "e1b0956f-f03c-4eae-8f11-8a52542ea9b5",
                    "calls": [

                    ],
                    "date_created": "2014-04-28 15:09:11 -0400",
                    "date_expires": "2014-04-28 15:11:11 -0400",
                    "date_scheduled": "2014-04-28 15:09:11 -0400",
                    "destination_address": "+17324961388",
                    "event_type": "outbound-call",
                    "owner_type": "account",
                    "parent_call_sid": "e91cd3af-2dfc-49e3-946f-6d25b46ea243",
                    "rated_event": {
                        "cost": 0,
                        "credits": 0,
                        "date_created": "2014-04-28 15:09:11 -0400",
                        "prefix": "+1",
                        "sid": "7d48ca33-d28b-424c-a8a4-750821073afc",
                        "units": 0,
                        "units_billed": 0
                    },
                    "sid": "07ca32b5-379e-4dfa-b57a-e89b53f28419",
                    "source_address": "+19082783791"
                }
            ],
            "date_created": "2014-04-28 15:09:11 -0400",
            "date_expires": "2014-04-28 15:11:11 -0400",
            "date_scheduled": "2014-04-28 15:09:11 -0400",
            "destination_address": "",
            "event_type": "inbound-call",
            "owner_type": "account",
            "parent_call_sid": null,
            "rated_event": {
                "cost": 0,
                "credits": 0,
                "date_created": "2014-04-28 15:09:11 -0400",
                "prefix": "",
                "sid": "5fbb6af0-344d-4537-925c-ac335ed892cb",
                "units": 0,
                "units_billed": 0
            },
            "sid": "e91cd3af-2dfc-49e3-946f-6d25b46ea243",
            "source_address": "+15038884341"
        }
    },
    "status": 200
}`)

	return respBytes, nil
}

func getJSONResponseError() ([]byte, error) {

	respBytes2 := []byte(`
    "::::::errsdfggor": null,
    "paginug_data": null,
    "resourasdasdces": {
        "access_numbers": [
            {
                "city": "",
                "country_code": "",
                "country_iso": "",
                "date_created": "2013-05-23 16:51:23 -0400",
                "description": "",
                "did": "+15038884341",
                "sid": "123",
                "state": ""
            }
        ],
        "account": {
            "balance": 10.833335,
            "credits": 65,
            "datasde_created": "2014-04-18 15:43:26 -0400",
            "email_address": "example@email.com",
            "first_name": "John",
            "has_password": false,
            "is_active": true,
            "is_fraud": false,
            "is_paid": true,
            "last_name": "Carter",
            "max_transfer_value": 0.0,
            "plugins": {
                "auto_replenishers": [

                ]
            },
            "security_questions": [

            ],
            "sid": "e1asdb0956f-f03c-4eae-8f11-8a52542ea9b5",
            "total_phones": 1
        },
        "call": {
            "access_code": "9130",
            "account_sid": "e1b0956f-f03c-4eae-8f11-8a52542ea9b5",
            "calls": [
                {
                    "access_code": "1220",
                    "account_sid": "e1b0956f-f03c-4eae-8f11-8a52542ea9b5",
                    "calls": [

                    ],
                    "date_asdcreated": "2014-04-28 15:09:11 -0400",
                    "date_expires": "2014-04-28 15:11:11 -0400",
                    "date_scheduled": "2014-04-28 15:09:11 -0400",
                    "destination_address": "+17324961388",
                    "event_type": "outbound-call",
                    "owner_type": "account",
                    "parentasd_call_sid": "e91cd3af-2dfc-49e3-946f-6d25b46ea243",
                    "rated_event": {
                        "cost": 0,
                        "crasdaedits": 0,
                        "date_created": "2014-04-28 15:09:11 -0400",
                        "prefix": "+1",
                        "sid": "7d48ca33-d28b-424c-a8a4-750821073afc",
                        "units": 0,
                        "units_billed": 0
                    },
                    "sid": "07ca32b5-379e-4dfa-b57a-e89b53f28419",
                    "source_address": "+19082783791"
                },
                {
                    "access_code": "1220",
                    "account_sid": "e1b0956f-f03c-4eae-8f11-8a52542ea9b5",
                    "calls": [

                    ],
                    "date_created": "2014-04-28 15:09:11 -0400",
                    "date_expires": "2014-04-28 15:11:11 -0400",
                    "date_scheduled": "2014-04-28 15:09:11 -0400",
                    "destination_address": "+17324961388",
                    "event_type": "outbound-call",
                    "owner_type": "account",
                    "parenasdasdt_call_sid": "e91cd3af-2dfc-49e3-946f-6d25b46ea243",
                    "rated_event": {
                        "cost": 0,
                        "credits": 0,
                        "date_created": "2014-04-28 15:09:11 -0400",
                        "prefix": "+1",
                        "sid": "7d48ca33-d28b-424c-a8a4-750821073afc",
                        "units": 0,
                        "units_billed": 0
                    },
                    "sid": "07ca32b5-379e-4dfa-b57a-e89b53f28419",
                    "source_address": "+19082783791"
                }
            ],
            "date_created": "2014-04-28 15:09:11 -0400",
            "date_expisdasdres": "2014-04-28 15:11:11 -0400",
            "date_scheduled": "2014-04-28 15:09:11 -0400",
            "destination_address": "",
            "event_type": "inbound-call",
            "owner_type": "account",
            "parent_call_sid": null,
            "rated_event": {
                "cost": 0,
                "credits": 0,
                "date_created": "2014-04-28 15:09:11 -0400",
                "prefix": "",
                "sid": "5fbb6af0-344d-4537-925c-ac335ed892cb",
                "units": 0,
                "units_billed": 0
            },
            "sid": "e91cd3af-2dfc-49e3-946f-6d25b46ea243",
            "source_address": "+15038884341"
        }
    },
    "status": 500
}`)

	return respBytes2, nil
}

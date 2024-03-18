# Golang ChatGPT client

Just be sure to use your own credentials

Output:

Request:

```json
{"model": "gpt-3.5-turbo", "messages":[{"role":"user","content":"test\r\n"}], "temperature":0.7}

```

Response:

```json
{  
	"error": {  
		"message": "Incorrect API key provided: API-KEY. You can find your API key at https://platform.openai.com/account/api-keys.",
	        "type": "invalid_request_error",
	        "param": null,
	    },
	"code": "invalid_api_key"
}
```

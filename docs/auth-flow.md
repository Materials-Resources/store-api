### Handling Authorization Scenario 1
```mermaid
sequenceDiagram
	participant WA as Customer Web App
	participant API as Customer API
	participant Z as Zidatel
	
	WA ->>Z:Intiate Login Session
	Z->>WA:Returns `access_token` for session
	WA ->>API: Sends `access_token` on request
	API->>Z: Checks validity of `access_token` and requests user info
	Z->>API: Returns `customer_id`
	API->>WA: Returns response with authorization logic using `customer_id`	
```

**Things to make this work**
- can `customer_id` be cached for that specific access token? 
- can `customer_id` be stored in the `access_token` and use the open id client to check and validate the signature?

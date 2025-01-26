# havs-service

## Description
This is a service based on the openAPI specification included in the repository.
It provides endpoints to manage exposures.

## Installation
To run the service with its mongo db database, you can use the docker-compose file provided in the repository.
```bash
docker-compose up
```

## Usage
The service is running on port 8080. You can access the endpoints using the following base url:
```
http://localhost:8080
```

Endpoints:
- /exposures
    - GET: Get all exposures
    - POST: Create a new exposure
- /exposures/{exposureId}
    - GET: Get an exposure by id
- /users/{userId}/exposure-summary
    - GET: Get the exposure summary for a user

## Testing
To run the tests, you can use the following command:
```bash
go test ./...
```

## TODO
- Add more tests
- Add update and delete endpoints for exposures

## Events
If this service was part of a bigger system, it would be useful to have events for the following actions:
- Exposure created
  - When an exposure is created this could be added to a kafka topic to be consumed by this service
  - This would ensure that the exposure is always recorded in the system. And won't fail due to the service being down temporarily.
- Exposure Summaries
  - If the exposure summary is taken at regular intervals, say to calculate the exposure for a user every day, this could be added to a kafka topic to be consumed by another service.
  - This could then be used to calculate the exposure for a user over time, and provide insights into the user's exposure.
  - This could also be used to alert the user if their exposure is too high.
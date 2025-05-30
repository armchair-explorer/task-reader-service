Steps to run:

After cloning this repo run:

go mod tidy

then inside the directory run -> go run cmd/reader-service/main.go

In the writer service:
Then run docker-compose up --build to start the db

sample curls:

Request:
curl "http://localhost:8070/tasks?page=1&limit=5&status=Pending"

Response:
[{"id":7,"title":"Write integration tests","description":"Write tests for the reader service","status":"Pending","created_at":"2025-05-30T19:38:20.081686Z","updated_at":"2025-05-30T19:38:20.081686Z"},
{"id":6,"title":"Write integration tests","description":"Write tests for the reader service","status":"Pending","created_at":"2025-05-30T19:38:19.231867Z","updated_at":"2025-05-30T19:38:19.231867Z"},
{"id":5,"title":"Write integration tests","description":"Write tests for the reader service","status":"Pending","created_at":"2025-05-30T19:38:18.290293Z","updated_at":"2025-05-30T19:38:18.290293Z"},
{"id":4,"title":"Write integration tests","description":"Write tests for the reader service","status":"Pending","created_at":"2025-05-30T19:38:15.522474Z","updated_at":"2025-05-30T19:38:15.522474Z"},
{"id":3,"title":"Write integration tests","description":"Write tests for the reader service","status":"Pending","created_at":"2025-05-30T19:38:13.571704Z","updated_at":"2025-05-30T19:38:13.571704Z"}]

Request:
curl "http://localhost:8070/tasks/7"                    
 
Response: 
{"id":7,"title":"Write integration tests","description":"Write tests for the reader service","status":"Pending","created_at":"2025-05-30T19:38:20.081686Z","updated_at":"2025-05-30T19:38:20.081686Z"}

Design Decisions:

1.I have choosen to separate the write service as the read service as the scaling factor can be different read service can have much higher throughput than the write service.

2.I have choosen postgres for the relational nature of the data.

3.I have indexed the table on status because the of the query by status. It will make the read much faster.

4.I have used soft deletion for delete. It gives us the ability to rollback.

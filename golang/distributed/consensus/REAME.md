$ curl "http://localhost:3333/?next=second"
second

$ curl "http://localhost:3333/?next=third"
third

$ curl "http://localhost:3333/?next=second"
invalid transition

$ curl "http://localhost:3333/?next=first"
first

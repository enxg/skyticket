# Skyticket
My submission for the YTU SKY LAB WEBLAB R&D Team Acceptance Task.

---

## Features
- Create, update, delete, and view events.
- Create, update, delete, and view tickets.
- Make reservations for tickets.
- OpenAPI documentation available at `/docs`. Powered by Scalar.

## Quick Start (Docker)
1. Copy `.env.example` to `.env` and fill in the values (explained below).
2. Run:
   ```sh
   docker compose up
   ```
   
## Environment Variables
- `MONGODB_URI` - MongoDB connection string (MongoDB Atlas is recommended as transactions are only supported on replica sets or sharded clusters).
- `OPENAPI_SCHEME` - The scheme to use in the OpenAPI spec (http or https).
- `OPENAPI_HOST` - The host to use in the OpenAPI spec (e.g. skyticket.enesgenc.dev).

## License
MIT

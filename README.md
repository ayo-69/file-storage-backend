# File Storage Backend

This is a simple backend server built with Express.js.

## Prerequisites

- [Node.js](https://nodejs.org/) (v16 or later)
- [Bun](https://bun.sh) (optional, for faster dependency management)

## Installation

To install dependencies, run:

```bash
bun install
```

If you prefer using npm, run:

```bash
npm install
```

## Running the Server

To start the server, use:

```bash
bun run start
```

Or with Node.js:

```bash
node index.js
```

The server will run on `http://localhost:3030` by default.

## Endpoints

### GET /

- **Description**: Returns a simple "Hello world" message.
- **Response**:
  ```json
  {
    "message": "Hello world"
  }
  ```

## Project Structure

```
.
├── .gitignore
├── index.js
├── package.json
├── README.md
└── routers/ (future expansion for route handlers)
```

## Notes

This project was initialized using `bun init` with Bun v1.1.43. [Learn more about Bun](https://bun.sh).

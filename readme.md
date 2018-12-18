Comment

It's a service for calculating Tribonacci sequence (0, 0, 1, 1, 2, 4, 7, 13, 24, 44...). The entry parameter (n) is an ordinal number. For example, for n = 10, the result is 44.

Instruction

Docker version should be 17.05 or higher

Run next commands from project's directory:

1. docker build -t tribonacci .
2. docker run -d -p 8000:8000 tribonacci

Available environment variables for configuration:

- TRIBONACCI_ADDR (default: ":8000") for server address
- TRIBONACCI_GRACE_TIMEOUT (default: 10) for grace timeout in seconds

Usage

GET /tribonacci/{n} (where n - positive integer number)

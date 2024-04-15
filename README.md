# Advanced Routing in Go 1.22

This repository contains code examples demonstrating the advanced routing capabilities introduced in Go 1.22. Focusing on the `net/http` package enhancements, these examples illustrate how to leverage new features such as method-based routing, path parameters, hostname-based routing, sub-routing, and middleware implementation for more efficient web development.

## Inspiration

The project was inspired by the video titled "The standard library now has all you need for advanced routing in Go" by Dreams of Code. The video explores the significant updates to Go's `net/http` package, providing a deep dive into how these advancements can be utilized to simplify and enhance routing in web applications.

Watch the video here: [The standard library now has all you need for advanced routing in Go](https://www.youtube.com/watch?v=H7tbjKFSg58&t)

## Features Explored

- **Path Parameters**: Using dynamic segments in the URL for more expressive route patterns.
- **Method-Based Routing**: Defining routes specific to HTTP methods (GET, POST, DELETE) directly within the route declaration.
- **Hostname-Based Routing**: Handling requests differently based on the requested hostname, ideal for multi-domain applications.
- **Sub-Routing**: Organizing routes using sub-routers for cleaner code and better structure.
- **Middleware**: Implementing middleware for cross-cutting concerns like logging, authentication, and more.

## Getting Started

To explore these examples, ensure you have Go 1.22 or later installed on your machine. Clone this repository and navigate into the project directory:

```bash
git clone https://github.com/dmelim/go-std-routing.git
cd yourrepository
```

And to run the code:

```bash
go run .
```

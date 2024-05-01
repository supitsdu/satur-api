# Satur API

Satur API is a project aimed at building a robust and efficient API using the Go programming language. It serves as a personal web development mockup project with the primary goal of learning and implementing best practices for API development.

> [!WARNING]
> **Experimental Project**: Please note that Satur API is an experimental project created with a learn-by-doing approach. While efforts have been made to ensure its functionality and stability, it may contain bugs or unforeseen issues. Use it with caution.

## Dependencies

- Go >= 1.22
- Make >= 4.4.1

## Getting Started

To get started with the Satur API, follow these steps:

1. **Clone the Repository**: Clone the Satur API repository to your local machine:

    ```sh
    git clone https://github.com/supitsdu/satur-api.git
    ```

2. **Install Dependencies**: Navigate to the project directory and run the following command to fetch or manage dependencies:

    ```sh
    make deps
    ```

3. **Create Environment Variables**: Before running the API, create a `.env` file in the project root directory with the following variables:

    ```sh
    MONGODB_URI=<your MongoDB connection string>
    MONGODB_ID=<your MongoDB database ID>
    SERVER_ADDRESS=localhost:8080
    ```

    Replace `<your MongoDB connection string>` and `<your MongoDB database ID>` with your MongoDB connection details.

4. **Run the API**: Once the dependencies are installed and the `.env` file is configured, execute the following command to run the API:

    ```sh
    make run
    ```

   This will start the API server, allowing you to interact with it locally.

## Contributing

Contributions to the Satur API project are welcome! If you find any bugs, have feature requests, or want to contribute code, please open an issue or submit a pull request on the [GitHub repository](https://github.com/supitsdu/satur-api).

## License

Satur API is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code for your own projects.

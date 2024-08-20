# Salesforce Data Generator

This project is a fictional data generator for Salesforce, using the Go Gin framework for the web server and the `faker` package to generate fake data.

**Note:** This project is currently under development and I am new to Go. I am actively working on adding new features, including additional Salesforce objects. I am also open to mentorship to improve my Go skills and general development knowledge. Any help, contribution, or suggestions are welcome!

## Features

- Generate fictional Salesforce accounts.
- Generate fictional Salesforce contacts.
- Download data in CSV format.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.18 or higher)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/julien5988/salesforce-data-generator.git
    cd salesforce-data-generator
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

## Usage

1. **Start the server:**

    ```sh
    go run main.go
    ```

2. **Access the application in your browser:**

    Open [http://localhost:8080](http://localhost:8080)

3. **Generate data:**

    - To generate accounts, use the form under the "Generate Accounts" section.
    - To generate contacts, use the form under the "Generate Contacts" section.

    CSV files will be available for download directly from the browser.

## Project Structure

- **`main.go`**: Entry point of the application, sets up routes and starts the server.
- **`routes/routes.go`**: Route configuration and request handling.
- **`routes/account_routes.go`**: Route handling for account generation.
- **`routes/contact_routes.go`**: Route handling for contact generation.
- **`generator/`**: Contains functions for generating fake data (accounts and contacts).
- **`utils/`**: Utility functions, such as writing data to CSV files.
- **`static/`**: Contains static files such as HTML, CSS, and JavaScript files.

## Contributing

Contributions are welcome! To suggest improvements or report issues, please follow these steps:

1. Fork the repository.
2. Create a branch for your feature (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -am 'Add new feature'`).
4. Push your branch (`git push origin feature/your-feature`).
5. Open a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

If you have any questions or suggestions, feel free to reach out to me:

- Email: [julien.thoumy@gmail.com](mailto:julien.thoumy@gmail.com)
- GitHub: [https://github.com/julien5988](https://github.com/julien5988)

---

Thank you for using this Salesforce data generator!

# GoVerifyDomain

## Introduction

GoVerifyDomain is a simple domain verification tool written in Go (Golang). This tool allows users to interactively verify various attributes of a domain, including MX records, SPF records, and DMARC records. The project aims to provide a hands-on understanding of DNS queries, email authentication, and security mechanisms.

## Features

- Interactive command-line tool for domain verification.
- Checks the presence of MX, SPF, and DMARC records.
- Displays the content of SPF and DMARC records if found.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/GoVerifyDomain.git
   ```

2. Navigate to the project directory:

   ```bash
   cd GoVerifyDomain
   ```

3. Build the executable:

   ```bash
   go build
   ```

4. Run the program:

   ```bash
   ./GoVerifyDomain
   ```

## Usage

1. Run the program and follow the prompts to enter a domain name for verification.
2. The tool will perform checks for MX, SPF, and DMARC records.
3. The results will be displayed, indicating whether each attribute is present.
4. If SPF or DMARC records are found, their content will be displayed.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
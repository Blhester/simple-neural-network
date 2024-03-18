# Simple Neural Network
> 
> This project is a simple implementation of a neural network in Go.

## Getting Started
> 
> These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
>
> - Go (version 1.16 or later)
> - Python (version 3.8 or later)

### Installing
>
> 1. Clone the repository
> ```bash
> git clone https://github.com/Blhester/simple-neural-network.git
> ```
> 2. Install the required packages
> ```bash
> go mod download
> ```
> 3. Install Datasets with Python (_Optional if you want to use the MNIST dataset_)
> ```bash
> ./prepare_mnist_dataset.sh
> ```

### Running the tests
> To run all tests in the current directory and all subdirectories, use the following command:
> ```bash
> go test ./...
> ```
> To run specific tests, use the -run flag followed by a regular expression. For example to run all Unit tests, use the following command:
> ```bash
> go test -run Unit ./...
> ```
> To test against the MNIST dataset, use the following command:
>
> _**Note**: This will only work if you have installed the MNIST dataset using the `./prepare_mnist_dataset` script._
> ```bash
> go test -run TestPlayground ./...
> ```

### Contributing
> Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### License
> This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

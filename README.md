# Dota 2 Advanced Hero Picker API

The Dota 2 Advanced Hero Picker API is a Golang-based backend web service that suggests a random Dota 2 hero to users based on their specified characteristics such as primary attribute and roles. The API is built using the Gin web framework and interacts with a PostgreSQL database to store and retrieve hero data.

## Requirements

- Golang 1.19 or later
- PostgreSQL database
- Kaggle API key (to fetch initial hero data)

## Installation and Setup

1. Clone the repository:

```bash
git clone https://github.com/jizambrana5/dota2-hero-picker
cd dota2-hero-picker
```

2. Install dependencies:

```bash
go mod download
```

3. Set up PostgreSQL:

    - Install PostgreSQL on your local machine or use a cloud-hosted PostgreSQL instance.
    - Create a new database and update the database connection configuration in `db.go` with the appropriate credentials.

4. Fetch hero data from Kaggle:

    - Obtain your Kaggle API key and set it in the `main.go` file (`const apiKey = "your_kaggle_api_key"`).
    - Run the script to fetch and insert data from Kaggle:

```bash
go run main.go
```

5. Build and run the API server:

```bash
go build
./dota2-hero-picker-api
```

The API server will start on `http://localhost:8080`.

## API Endpoints

- `GET /api/heroes`: Returns the list of all heroes with their details (Hero Index, Primary Attribute, Name in Game, and Role).
- `POST /api/hero-picker`: Accepts user preferences as input (in JSON format) and returns a random hero suggestion that matches the specified characteristics.

## Sample Usage

1. Fetch all heroes:

```bash
curl http://localhost:8080/api/heroes
```

2. Get a random hero suggestion based on user preferences:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"primary_attribute": "agi", "roles": ["Carry", "Disabler"]}' http://localhost:8080/api/hero-picker
```

## Concurrency

The API utilizes Goroutines for concurrent filtering and hero selection, ensuring improved performance and responsiveness.

## Dockerization and Cloud Deployment

To deploy the API using Docker and CI/CD, follow the steps outlined in the project documentation.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project uses data from Kaggle: [Kaggle Dota 2 Dataset](https://www.kaggle.com/some/dataset/url)

Feel free to explore the API and use it to discover new Dota 2 heroes based on your preferences!

For more information and updates, check out the GitHub repository: [Dota 2 Hero Picker API](https://github.com/your-username/dota2-hero-picker-api).
```

Please customize the placeholders (`your-username`, `your_kaggle_api_key`, etc.) and update the sections accordingly based on your actual project details.

The README file serves as the primary documentation for your project and helps users and potential employers understand how to use and set up your API. It includes information about installation, API endpoints, sample usage, concurrency, and deployment. Additionally, it provides license details and acknowledgments.

Feel free to enhance the README with additional sections, badges, or more detailed instructions as needed. A well-written README is crucial for project visibility and a positive user experience.
# RankEdge
RankEdge is a real-time leaderboard system built with Go and powered by Redis. It supports instant ranking updates, efficient queries, and scalability for millions of users. Ideal for gaming and competitions, RankEdge ensures low-latency and high availability for seamless user experiences.

## Features

- **Leaderboard Management:** Create, update, and retrieve real-time user rankings.
- **User Management:** Secure user registration and login.
- **Real-Time Scoring:** Instantly update and display user scores.
- **Redis Integration:** Leveraging Redis Sorted Sets for fast, scalable rank management.
- **API Documentation:** Fully documented RESTful API with Swagger.

## Getting Started

Follow these instructions to set up and run the RankEdge project on your local machine for development and testing purposes.

### Prerequisites

Ensure that you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.23 or later)
- Docker
- Docker compose

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/nubufi/RankEdge.git
   cd RankEdge
   ```
2. Build and deploy the app:
   ```sh
   docker compose up --build
   ```

## Configuration

The application can be configured using environment variables. Refer to the `.env` file for available configuration options such as database credentials, API keys, and more.

## Usage

The API documentation can be accessed at `http://localhost:8080/docs/index.html`.

## Contributing

We welcome contributions from the community! To contribute to RankEdge:

1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

For any inquiries or feedback, feel free to reach out to us:

- **Website:** [www.nubufi.com](https://www.nubufi.com)
- **Email:** numanburakfidan@yandex.com

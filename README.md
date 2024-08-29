# Cat API Project

This project is built using Go and the Beego framework. It leverages the [TheCatAPI](https://thecatapi.com/) services to provide various cat-related functionalities, including voting on cat images (like, dislike, love), viewing cat breeds' descriptions, and managing a favorites section to see previously added favorites. The project also utilizes Go channels to handle asynchronous API calls, ensuring efficient handling of concurrent operations.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)
- [License](#license)


## Features

- **Voting**: Users can vote on cat images with options such as like, dislike, or love.
- **Breed Descriptions**: Display information about different cat breeds.
- **Favorites Section**: Users can add cat images to their favorites and view them later.

## Installation

To run this project locally, follow these steps:
1. **Install Go and Beego**

   - For go follow - https://go.dev/doc/install
   - For Beego follow - https://beego.w3engineers.com/beego.vip/index.html

2. **Clone the repository (into src folder of go)**:

   ```bash
   git clone https://github.com/w3-software-intern-Riad/GO-BEEGO-CATAPI.git
   cd GO-BEEGO-CATAPI
   ```

3. **Configuration**:
    - **API Key**: This project requires an API key from [TheCatAPI](https://thecatapi.com/). You can get a free API key by signing up on their website.

    - **Setup Configuration File**: Create a `conf/app.conf` file and add the following configuration settings or follow the `conf/app.conf.sample`:

     ```ini
       appname = catProject
       httpport = 8080
       runmode = dev

       baseUrl = https://api.thecatapi.com
       apiKey = your_api_key_here
   ```

   Replace `your_api_key_here` with the actual API key you obtained from TheCatAPI.


4. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

5. **Run the application**:
   ```bash
   bee run
   ```
6. **Access the frontend**:
   ```bash 
    http://localhost:8080
   ```




## Usage

After running the application, you can interact with the following features through the frontend:

- **Vote on Cat Images**: Like, dislike, or love cat images displayed on the site.
- **View Cat Breed Descriptions**: Browse descriptions of different cat breeds.
- **Manage Favorites**: Add cat images to your favorites and view them in a dedicated favorites section.


## Technologies Used

- **Go**: Programming language used for building the backend.
- **Beego**: Web framework used to develop the application.
- **TheCatAPI**: External API service for fetching cat-related data.


## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. **Fork the repository**: Click the "Fork" button at the top right of this repository to create your own copy of the project.

2. **Clone the forked repository**:
    ```bash
    git clone https://github.com/w3-software-intern-Riad/GO-BEEGO-CATAPI.git
    ```

3. **Create a new branch** for your feature or bug fix:
    ```bash
    git checkout -b feature-or-bugfix-name
    ```

4. **Make your changes**: Implement your feature or fix the bug.

5. **Commit your changes** with a descriptive commit message:
    ```bash
    git commit -m "Add feature or fix description"
    ```

6. **Push to the branch**:
    ```bash
    git push origin feature-or-bugfix-name
    ```

7. **Create a Pull Request**: Go to the original repository on GitHub and create a pull request from your forked repository's branch.

8. **Review**: Your pull request will be reviewed, and suggestions may be provided.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
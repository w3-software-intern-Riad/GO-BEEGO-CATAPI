<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Cat Voting App</title>
    <link rel="stylesheet" href="/static/css/style.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
    <div class="container">
        <div class="header">
            <button class="nav-btn" onclick="window.location.href='/';">Voting</button>
            <button class="nav-btn" onclick="window.location.href='/get-breeds';">Breeds</button>
            <button class="nav-btn" onclick="window.location.href='/get-favorite';">Favs</button>
        </div>

        <div class="breeds-section">
            <label for="breeds-dropdown">Select a Cat Breed:</label>
            <select id="breeds-dropdown" class="dropdown">
                <option value="">--Select a Breed--</option>
            </select>

            <div class="breed-info">
                <div class="image-slider" id="image-slider">
                    <button class="slider-btn prev-btn" onclick="prevImage()">&#10094;</button>
                    <button class="slider-btn next-btn" onclick="nextImage()">&#10095;</button>
                    <div class="images-container" id="images-container">
                        <!-- Images will be dynamically added here -->
                    </div>
                </div>
                <div class="slider-navigation" id="slider-navigation">
                    <!-- Navigation dots will be dynamically added here -->
                </div>
                <h2></h2><h3></h3><h4></h4>
                <p id="breed-description"></p>
                <a id="wikipedia-link" class="wiki-btn" href="#" target="_blank">Learn More on Wikipedia</a>
            </div>
        </div>
    </div>

    <script src="/static/js/getBreeds.js"></script>
</body>
</html>

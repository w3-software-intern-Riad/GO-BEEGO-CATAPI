document.addEventListener("DOMContentLoaded", () => {
    async function showFavorites() {
        try {
            const response = await fetch("/get-favorite-ctl"); // Replace with your actual endpoint
            if (!response.ok) {
                throw new Error("Failed to fetch favorite images");
            }
            const favoriteImages = await response.json();
            
            const favoritesContainer = document.getElementById("favorites-container");
            favoritesContainer.innerHTML = ""; // Clear any existing content

            favoriteImages.forEach(image => {
                // Check if the image URL is not empty
                if (image.image && image.image.url) {
                    console.log("image url :",image.image.url)
                    const imgElement = document.createElement("img");
                    imgElement.src = image.image.url; // Use the correct URL
                    imgElement.alt = "Favorite Cat Image";
                    imgElement.className = "favorite-cat-image"; // Add a class for styling
                    favoritesContainer.appendChild(imgElement);
                }
            });
        } catch (error) {
            console.error("Error fetching favorite images:", error);
        }
    }

    showFavorites();
});

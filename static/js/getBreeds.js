document.addEventListener("DOMContentLoaded", function () {
    const breedsInput = document.getElementById("breeds-input");
    const breedsListDiv = document.getElementById("breeds-list");
    const breedDescription = document.getElementById("breed-description");
    const name = document.getElementsByTagName("h2")[0];
    const origin = document.getElementsByTagName("h3")[0];
    const id = document.getElementsByTagName("h4")[0];
    const wikipediaLink = document.getElementById("wikipedia-link");
    const imagesContainer = document.getElementById("images-container");
    const sliderNavigation = document.getElementById("slider-navigation");

    let breedsList = [];
    let currentSlide = 0;

    // Fetch and populate the breeds list
    fetch("/get-breeds-ctl")
        .then((response) => response.json())
        .then((breeds) => {
            breedsList = breeds;
            showFilteredBreeds(breedsList);
            if (breedsList.length > 0) {
                selectBreed(breedsList[0].id); // Automatically select the first breed
            }
        });

    // Show the dropdown list when clicking on the input field
    breedsInput.addEventListener("click", function (event) {
        event.stopPropagation(); // Prevents the click event from bubbling up and closing the dropdown
        breedsListDiv.style.display = breedsListDiv.style.display === "block" ? "none" : "block";
    });

    // Event listener for typing in the input field
    breedsInput.addEventListener("input", function () {
        const searchTerm = breedsInput.value.toLowerCase();
        const filteredBreeds = breedsList.filter((breed) =>
            breed.name.toLowerCase().includes(searchTerm)
        );
        showFilteredBreeds(filteredBreeds);
    });

    // Display filtered breeds in the dropdown
    function showFilteredBreeds(breeds) {
        breedsListDiv.innerHTML = "";
        breedsListDiv.style.display = breeds.length > 0 ? "block" : "none";
        breeds.forEach((breed) => {
            const breedDiv = document.createElement("div");
            breedDiv.textContent = breed.name;
            breedDiv.onclick = () => {
                selectBreed(breed.id);
                breedsListDiv.style.display = "none"; // Hide dropdown after selection
            };
            breedsListDiv.appendChild(breedDiv);
        });
    }

    // Handle breed selection
    function selectBreed(breedId) {
        const selectedBreed = breedsList.find((breed) => breed.id === breedId);

        if (selectedBreed) {
            breedDescription.textContent = selectedBreed.description;
            name.innerHTML = selectedBreed.name;
            id.innerHTML = selectedBreed.id;
            origin.innerHTML = selectedBreed.origin;
            wikipediaLink.href = selectedBreed.wikipedia_url;
            wikipediaLink.textContent = `Learn more about ${selectedBreed.name}`;
            wikipediaLink.style.display = "inline-block";

            imagesContainer.innerHTML = "";
            sliderNavigation.innerHTML = "";

            fetch(`/cat-images/${selectedBreed.id}`)
                .then((response) => response.json())
                .then((images) => {
                    images.forEach((imgUrl, index) => {
                        const img = document.createElement("img");
                        img.src = imgUrl;
                        if (index === 0) img.classList.add("active");
                        imagesContainer.appendChild(img);

                        const dot = document.createElement("div");
                        dot.classList.add("slider-dot");
                        if (index === 0) dot.classList.add("active");
                        dot.addEventListener("click", () => showSlide(index));
                        sliderNavigation.appendChild(dot);
                    });

                    currentSlide = 0; // Reset to the first slide
                });

            breedsInput.value = selectedBreed.name;
        }
    }

    // Show the next/previous image
    function showSlide(index) {
        const images = imagesContainer.getElementsByTagName("img");
        const dots = sliderNavigation.getElementsByClassName("slider-dot");

        if (images.length > 0) {
            images[currentSlide].classList.remove("active");
            dots[currentSlide].classList.remove("active");
            currentSlide = index;
            if (currentSlide >= images.length) currentSlide = 0;
            if (currentSlide < 0) currentSlide = images.length - 1;
            images[currentSlide].classList.add("active");
            dots[currentSlide].classList.add("active");
        }
    }

    // Event listeners for slider buttons
    document.querySelector(".prev-btn").addEventListener("click", function () {
        showSlide((currentSlide - 1 + imagesContainer.getElementsByTagName("img").length) % imagesContainer.getElementsByTagName("img").length);
    });

    document.querySelector(".next-btn").addEventListener("click", function () {
        showSlide((currentSlide + 1) % imagesContainer.getElementsByTagName("img").length);
    });

    // Hide the dropdown when clicking outside
    document.addEventListener("click", function (event) {
        if (!breedsInput.contains(event.target) && !breedsListDiv.contains(event.target)) {
            breedsListDiv.style.display = "none";
        }
    });
});

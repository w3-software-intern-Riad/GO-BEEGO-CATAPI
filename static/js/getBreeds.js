document.addEventListener("DOMContentLoaded", function () {
    const breedsDropdown = document.getElementById("breeds-dropdown");
    const breedDescription = document.getElementById("breed-description");
    const name = document.getElementsByTagName("h2")[0];
    const origin = document.getElementsByTagName("h3")[0];
    const id = document.getElementsByTagName("h4")[0];
    const wikipediaLink = document.getElementById("wikipedia-link");
    const imagesContainer = document.getElementById("images-container");
    const sliderNavigation = document.getElementById("slider-navigation");
  
    let breedsList = [];
    let currentSlide = 0;
  
    // Fetch and populate the breeds dropdown
    fetch("/get-breeds-ctl")
      .then((response) => response.json())
      .then((breeds) => {
        breedsList = breeds;
        breeds.forEach((breed) => {
          const option = document.createElement("option");
          option.value = breed.id;
          option.textContent = breed.name;
          breedsDropdown.appendChild(option);
        });
  
        // Select the first breed by default and trigger the change event
        if (breedsList.length > 0) {
          breedsDropdown.value = breedsList[0].id;
          breedsDropdown.dispatchEvent(new Event("change"));
        }
      });
  
    breedsDropdown.addEventListener("change", function () {
      const selectedBreedId = this.value;
  
      if (selectedBreedId) {
        const selectedBreed = breedsList.find(
          (breed) => breed.id === selectedBreedId
        );
        
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
  
          fetch(`/cat-images/${selectedBreedId}`)
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
        }
      } else {
        breedDescription.textContent = "";
        wikipediaLink.style.display = "none";
        imagesContainer.innerHTML = "";
        sliderNavigation.innerHTML = "";
      }
    });
  
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
  
    window.prevImage = function () {
      const images = imagesContainer.getElementsByTagName("img");
      if (images.length > 0) {
        showSlide((currentSlide - 1 + images.length) % images.length);
      }
    };
  
    window.nextImage = function () {
      const images = imagesContainer.getElementsByTagName("img");
      if (images.length > 0) {
        showSlide((currentSlide + 1) % images.length);
      }
    };
  });
  
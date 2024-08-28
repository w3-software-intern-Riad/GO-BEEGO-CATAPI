document.addEventListener("DOMContentLoaded", () => {
  const likeButton = document.querySelector(".like-btn");
  const dislikeButton = document.querySelector(".dislike-btn");
  const loveButton = document.querySelector(".love-btn");
  const catImageElement = document.getElementById("cat-image");

  // Function to fetch cat image data from the API
  async function fetchCatImage() {
    try {
      const response = await fetch("/cat"); // Your API endpoint to get cat image
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.json();
      console.log("current Image data :", data);

      // Check if the data has the expected structure
      if (data && data.url) {
        catImageElement.src = data.url; // Update the image source
      } else {
        console.error("Invalid data format");
      }
    } catch (error) {
      console.error("Fetch error:", error);
    }
  }

  async function sendVote(value) {
    try {
      const imageId = catImageElement.src.split("/").pop().split(".")[0];
      console.log("imageId: ", imageId);
      const voteData = {
        image_id: imageId,
        sub_id: "my-user-1234", // Replace with actual user ID
        value: value,
      };
      console.log("vote data : ", voteData);
      const response = await fetch("/vote", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(voteData),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Network response was not ok");
      }
      const data = await response.json();
      if (data.message === "SUCCESS") {
        alert("Vote is done");
      }
    } catch (error) {
      console.error("Error sending vote:", error);
    }
  }

  async function addFavorite() {
    try {
      const imageId = catImageElement.src.split("/").pop().split(".")[0];
      const favData = {
        image_id: imageId,
        sub_id: "my-user-1234", // Replace with actual user ID
      };
      const response = await fetch("/favorite", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(favData),
      });
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "Network response was not ok");
      }
      const data = await response.json();
      if (data.message === "SUCCESS") {
        alert("Add to favorite done ");
      }
    } catch (error) {
      console.error("Error add to favorite :", error);
    }
  }

  // Event listener for like button
  likeButton.addEventListener("click", async (event) => {
    event.preventDefault();
    await sendVote(1); // Send an upvote
    await fetchCatImage(); // Fetch a new cat image when the like button is clicked
  });

  // Event listener for dislike button
  dislikeButton.addEventListener("click", async (event) => {
    event.preventDefault();
    await sendVote(-1); // Send a downvote
    await fetchCatImage(); // Fetch a new cat image when the dislike button is clicked
  });

  loveButton.addEventListener("click", async (event) => {
    event.preventDefault();
    await addFavorite();
    await fetchCatImage(); // Fetch a new cat image when the like button is clicked
  });

  // Fetch the initial cat image when the page loads
  fetchCatImage();
});

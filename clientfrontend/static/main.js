

// Handle form submission for creating a new post
const createPostForm = document.getElementById("create-post-form");
createPostForm.addEventListener("submit", function (event) {
  event.preventDefault();

  // Get the form data
  const formData = new FormData(createPostForm);
  const postData = Object.fromEntries(formData.entries());

  // Send the post data to the server
  fetch("/create-post", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(postData),
  })
    .then((response) => response.json())
    .then((data) => {
      // Handle the response from the server
      console.log(data);
      // Refresh the page or update the UI as needed
    })
    .catch((error) => {
      console.error("Error:", error);
    });
});

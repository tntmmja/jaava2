// Get posts and display them
function getPosts() {
  fetch('/loggedin')
    .then(response => response.json())
    .then(posts => {
      // Handle the response data (posts)
      console.log(posts);
      // Display posts on the page
      const postContainer = document.getElementById('post-container');
      postContainer.innerHTML = '';

      posts.forEach(post => {
        const postElement = document.createElement('div');
        postElement.innerHTML = `
          <h3>${post.title}</h3>
          <p>${post.content}</p>
          <p>Created at: ${post.created_at}</p>
        `;
        postContainer.appendChild(postElement);
      });
    })
    .catch(error => {
      console.error('Error:', error);
    });
}

// Call the getPosts function when the page is loaded
window.addEventListener('load', getPosts);

// Handle the creation of a new post
document.getElementById("create-post-form").addEventListener("submit", async (e) => {
  e.preventDefault();

  // Collect form data
  const formData = new FormData(e.target);
  const title = formData.get("title");
  const content = formData.get("content");

  // Send the post data to the server
  const response = await fetch("/create-post", {
    method: "POST",
    body: JSON.stringify({
      title,
      content
    }),
    headers: {
      "Content-Type": "application/json"
    }
  });

  if (response.ok) {
    // Post creation successful, display a success message
    document.getElementById("create-post-message").textContent = "Post created successfully";
    // Clear the form fields
    e.target.reset();
    // Refresh the posts to display the newly created post
    getPosts();
  } else {
    // Post creation failed, display an error message
    document.getElementById("create-post-message").textContent = "Failed to create post";
  }
});

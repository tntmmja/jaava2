// Fetch posts when the feed page is loaded
fetchPosts();



// Function to toggle the visibility of the create post form
function toggleCreatePostForm() {
    var createPostForm = document.getElementById("create-post-form");
    createPostForm.style.display = createPostForm.style.display === "none" ? "block" : "none";
}

// Add event listener to the "Create Post" button
document.getElementById("create-post-button").addEventListener("click", toggleCreatePostForm);

// Function to handle the form submission for creating a new post
function handleCreatePost(event) {
    event.preventDefault();

    // Get the form data
    var form = document.getElementById("create-post-form");
    var data = {
        title: form.title.value,
        text: form.text.value
    };

    // Send the form data to the server
    fetch("/create-post", {///feed
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then(function (response) {
            if (response.ok) {
                // Post created successfully, fetch and render the updated posts
                fetchPosts();
            } else {
                throw new Error("Failed to create post");
            }
        })
        .catch(function (error) {
            console.error(error);
            // Handle error while creating the post
        });
}

// Function to handle the form submission for creating a new post
document.getElementById("create-post-form").addEventListener("submit", handleCreatePost);

// Fetch posts from the server and render them on the feed page
function fetchPosts() {
    fetch("/feed")
        .then(function (response) {
            if (response.ok) {
                return response.json();
            } else {
                throw new Error("Failed to fetch posts");
            }
        })
        .then(function (posts) {
            // Render the posts on the feed page
            var postsContainer = document.getElementById("posts");
            postsContainer.innerHTML = ""; // Clear previous posts

            posts.forEach(function (post) {
                var postElement = createPostElement(post);
                postsContainer.appendChild(postElement);
            });
        })
        .catch(function (error) {
            console.error(error);
            // Handle error while fetching posts
        });
}

// Create HTML element for a single post
function createPostElement(post) {
    var postElement = document.createElement("div");
    postElement.classList.add("post");

    var titleElement = document.createElement("h2");
    titleElement.textContent = post.title;
    postElement.appendChild(titleElement);

    var textElement = document.createElement("p");
    textElement.textContent = post.text;
    postElement.appendChild(textElement);

    // You can add more elements or customize the post rendering as per your requirements

    return postElement;
}


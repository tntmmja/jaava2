// Fetch posts when the feed page is loaded
fetchPosts();

// Function to handle the form submission
function createPost(event) {
    event.preventDefault();
    
    // Get the form data
    var form = document.getElementById("create-post-form");
    var data = {
        title: form.title.value,
        text: form.text.value,
        user_id: form.user_id.value // Add the user ID field
    };
    
    // Send the form data to the server
    fetch("/create-post", {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json"
        }
    })
    .then(function(response) {
        if (response.ok) {
            console.log("Post created successfully");
            // Refresh the feed to display the new post
            fetchPosts();
            // Reset the form fields
            form.title.value = "";
            form.text.value = "";
        } else {
            console.error("Post creation failed");
            // Handle post creation failure
        }
    })
    .catch(function(error) {
        console.error("Post creation failed", error);
        // Handle post creation failure
    });
}

// Attach the form submission event listener
document.getElementById("create-post-form").addEventListener("submit", createPost);

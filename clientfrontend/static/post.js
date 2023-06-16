function getPosts() {
    fetch('/api/posts')
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
            <p>${post.text}</p>
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
  
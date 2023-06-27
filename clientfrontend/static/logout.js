// logout.js

function logoutUser() {
    fetch('/logout', {
      method: 'GET',
    })
      .then(response => {
        // Clear any cached data or session on the frontend
        // Redirect to the homepage
        window.location.href = "/";
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }
  
  // Call the logoutUser function when the logout button is clicked
  const logoutButton = document.getElementById('logoutButton');
  logoutButton.addEventListener('click', logoutUser);
  
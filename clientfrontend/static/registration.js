// Perform user registration
function registerUser() {
    // Get form input values
    const nickname = document.getElementById('nickname').value;
    const age = document.getElementById('age').value;
    const gender = document.getElementById('gender').value;
    const firstName = document.getElementById('firstName').value;
    const lastName = document.getElementById('lastName').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
  
    // Create user object
    const user = {
      nickname: nickname,
      age: age,
      gender: gender,
      firstName: firstName,
      lastName: lastName,
      email: email,
      password: password
    };
  
    // Send a POST request to the backend API for user registration
    fetch('/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(user)
    })
      .then(response => response.json())
      .then(data => {
        // Handle the response from the backend
        if (data.message === 'Registration successful') {
          // Registration successful, redirect to the login page
          window.location.href = '/login.html';
        } else {
          // Display the error message returned from the backend
          const errorMessage = document.getElementById('errorMessage');
          errorMessage.innerText = data.message;
        }
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }
  
  // Add an event listener to the registration form submit button
  const registerForm = document.getElementById('registerForm');
  registerForm.addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent the default form submission
    registerUser(); // Call the registerUser function
  });
  
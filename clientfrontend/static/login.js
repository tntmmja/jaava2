function loginUser(event) {
  event.preventDefault(); // Prevent the default form submission

  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value;

  const data = {
    username: username,
    password: password
  };

  fetch('/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
  })
  .then(response => {
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return response.json();
   // return response.text(); // Return the response as text
  })
    .then(data => {
      // Handle the response data
      if (data.success) {
        // Login successful, redirect to the main page
        window.location.href = "/loggedin";
      } else {
        // Login failed, display an error message
        const loginMessage = document.getElementById('login-message');
        loginMessage.textContent = "Login failed";
      }
    // .then(responseText => {
    //   console.log('Response:', responseText); // Log the response as text
    //   const data = JSON.parse(responseText); // Parse the response as JSON
    //   // Handle the response data
    //   if (data.success) {
    //     // Login successful, redirect to the main page
    //     window.location.href = "/loggedin";
    //   } else {
    //     // Login failed, display an error message
    //     const loginMessage = document.getElementById('login-message');
    //     loginMessage.textContent = "Login failed";
    //   }


    })
    .catch(error => {
      console.error('Error:', error);
      const loginMessage = document.getElementById('login-message');
      loginMessage.textContent = "An error occurred during login.";


    });
}





// //some other option of the code
// // login.js
// // Handle the login form submission
// document.getElementById("login-form").addEventListener("submit", async (e) => {
//   e.preventDefault();

//   // Collect form data
//   const formData = new FormData(e.target);
//   const email = formData.get("email");
//   const password = formData.get("password");

//   // Send the login data to the server
//   const response = await fetch("/login", {
//     method: "POST",
//     body: JSON.stringify({
//       email,
//       password
//     }),
//     headers: {
//       "Content-Type": "application/json"
//     }
//   });

//   if (response.ok) {
//     // Login successful, redirect to the main page
//     window.location.href = "/";
//   } else {
//     // Login failed, display an error message
//     document.getElementById("login-message").textContent = "Login failed";
//   }
// });


// Home component
class Home {
  render() {
    const div = document.createElement('div');
    div.innerHTML = '<h1>Welcome to the Home</h1>';
    return div;
  }
}

// // Register component
// class Register {
//   render() {
//     const div = document.createElement('div');
//     div.innerHTML = `
//       <h1>Register Page</h1>
//       <form id="registerForm">
//         <label for="username">Username:</label>
//         <input type="text" id="username" name="username" required>
//         <br>
//         <label for="password">Password:</label>
//         <input type="password" id="password" name="password" required>
//         <br>
//         <label for="email">Email:</label>
//         <input type="email" id="email" name="email" required>
//         <br>
//         <button type="submit">Register</button>
//       </form>
//     `;

//     return div;
//   }

//   componentDidMount() {
//     const registerForm = document.getElementById('registerForm');
//     if (registerForm) {
//       registerForm.addEventListener('submit', (event) => {
//         event.preventDefault();
//         registerUser();
//       });
//     }
//   }
// }

// Login component
class Login {
  render() {
    const div = document.createElement('div');
    div.innerHTML = `
      <h1>Login Page</h1>
      <form id="loginForm">
        <label for="username">Username:</label>
        <input type="text" id="username" name="username" required>
        <br>
        <label for="password">Password:</label>
        <input type="password" id="password" name="password" required>
        <br>
        <button type="submit">Login</button>
      </form>
    `;
    return div;
  }
}

// Define the routes
const routes = [
  { path: '/', component: Home },
 // { path: '/register', component: Register },
  { path: '/login', component: Login },
];

// Function to render a component based on the current route
function renderComponent() {
  const path = window.location.pathname;
  const route = routes.find(r => r.path === path);

  if (route && route.component) {
    const contentDiv = document.getElementById('content');
    contentDiv.innerHTML = ''; // Clear the content div
    const component = new route.component();
    contentDiv.appendChild(component.render());
    if (component.componentDidMount) {
      component.componentDidMount();
    }
  }
}


// Function to handle navigation
function navigateTo(url) {
  window.history.pushState(null, null, url);
  renderComponent();
}

// Event listener for popstate event (browser back/forward buttons)
window.addEventListener('popstate', () => {
  renderComponent();
});

document.addEventListener('DOMContentLoaded', () => {
  renderComponent();

  const registerButton = document.createElement('button');
  registerButton.textContent = 'Register';
  registerButton.addEventListener('click', () => {
    navigateTo('/register');
  });

  const redirectToRegisterButton = document.createElement('button');
  redirectToRegisterButton.textContent = 'Go to Register';
  redirectToRegisterButton.addEventListener('click', () => {
    navigateTo('/register');
  });

  const appDiv = document.getElementById('app');
  appDiv.appendChild(registerButton);
  appDiv.appendChild(redirectToRegisterButton);

  const loginForm = document.getElementById('loginForm');
  loginForm.addEventListener('submit', (event) => {
    event.preventDefault();
    loginUser(event);
  });
});
// // this function is in login.go
// function loginUser() {
//   const username = document.getElementById('username').value;
//   const password = document.getElementById('password').value;

//   const data = {
//     username: username,
//     password: password
//   };

//   fetch('/login', {
//     method: 'POST',
//     headers: {
//       'Content-Type': 'application/json'
//     },
//     body: JSON.stringify(data)
//   })
//     .then(response => response.json())
//     .then(data => {
//       // Handle the response data
//       if (data.success) {
//         // Login successful, redirect to the logged-in page
//         window.location.href = "/loggedin";
//       } else {
//         // Login failed, display an error message
//         const loginMessage = document.getElementById('login-message');
//         loginMessage.textContent = "Login failed";
//       }
//     })
//     .catch(error => {
//       console.error('Error:', error);
//     });
// }


// WebSocket connection
const socket = new WebSocket("ws://localhost:8082/loggedin");

// Event handler for WebSocket connection open
socket.onopen = () => {
  console.log("WebSocket connection established");
  // Additional initialization or actions after connection is open
};

// Event handler for WebSocket messages
socket.onmessage = (event) => {
  const message = event.data;
  console.log("Received message:", message);
  // Process the received message, update UI, etc.
};

// Event handler for WebSocket connection close
socket.onclose = () => {
  console.log("WebSocket connection closed");
  // Handle connection close, perform cleanup, etc.
};

// Function to send a message through WebSocket
function sendMessage(message) {
  socket.send(JSON.stringify(message));
  console.log("Sent message:", message);
  // Additional actions after sending the message
}

// Usage: sending a message
sendMessage({ text: "Hello, server!" });

// Usage: handling user input and sending messages
const inputElement = document.getElementById("messageInput");
const sendButton = document.getElementById("sendButton");

sendButton.addEventListener("click", () => {
  const message = inputElement.value;
  sendMessage({ text: message });
  inputElement.value = "";
});

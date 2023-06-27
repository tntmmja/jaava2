// // Create custom headers
// const headers = {
//   "Custom-Header-1": "Value1",
//   "Custom-Header-2": "Value2"
// };

// // Convert the headers object to a string
// const headersString = Object.keys(headers)
//   .map(key => `${key}: ${headers[key]}`)
//   .join("\n");






// Create a WebSocket connection
const socket = new WebSocket("ws://localhost:8082/loggedin");

// Event handler for WebSocket connection open
socket.onopen = () => {
  console.log("WebSocket connection established");
  console.log("Connection headers:", socket.extensions, socket.protocol); // Log the connection headers
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
  socket.send(message);
  console.log("Sent message:", message);
  // Additional actions after sending the message
}

// Example message
sendMessage("Hello, server!");

// Example handling user input and sending messages
const inputElement = document.getElementById("messageInput");
const sendButton = document.getElementById("sendButton");

sendButton.addEventListener("click", () => {
  const message = inputElement.value;
  sendMessage(message);
  inputElement.value = "";
});

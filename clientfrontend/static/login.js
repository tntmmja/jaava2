function login() {
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
      .then(response => response.json())
      .then(data => {
        // Handle the response data
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }
  
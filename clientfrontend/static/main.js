
document.addEventListener('DOMContentLoaded', function() {
  const appDiv = document.getElementById('app');
  
  const registerButton = document.createElement('button');
  registerButton.classList.add('button');
  registerButton.textContent = 'Register';
  registerButton.addEventListener('click', function() {
      navigateTo('/register');
  });
  appDiv.appendChild(registerButton);
  
  const loginButton = document.createElement('button');
  loginButton.classList.add('button');
  loginButton.textContent = 'Login';
  loginButton.addEventListener('click', function() {
      navigateTo('/login');
  });
  appDiv.appendChild(loginButton);
});

function navigateTo(url) {
  window.location.href = url;
}

// Login Modal
const loginBtn = document.querySelector('.login-btn');
const loginModal = document.getElementById('loginModal');
const closeBtn = document.querySelector('.close');

loginBtn.addEventListener('click', () => {
    loginModal.style.display = 'block';
});

closeBtn.addEventListener('click', () => {
    loginModal.style.display = 'none';
});

window.addEventListener('click', (event) => {
    if (event.target === loginModal) {
        loginModal.style.display = 'none';
    }
});

// Chat Interface
const chatMessages = document.getElementById('chatMessages');
const userInput = document.getElementById('userInput');

function sendMessage() {
    const message = userInput.value.trim();
    if (message) {
        addMessage('user', message);
        userInput.value = '';
        // Simulate AI response (replace with actual API call)
        setTimeout(() => {
            const aiResponse = "Thank you for your question. I'm processing your request...";
            addMessage('ai', aiResponse);
        }, 1000);
    }
}

function addMessage(sender, message) {
    const messageElement = document.createElement('div');
    messageElement.classList.add('chat-message', `${sender}-message`);
    messageElement.textContent = message;
    chatMessages.appendChild(messageElement);
    chatMessages.scrollTop = chatMessages.scrollHeight;
}

// Feedback Submission
const feedbackForm = document.getElementById('feedbackForm');
const feedbackText = document.getElementById('feedbackText');

feedbackForm.addEventListener('submit', (e) => {
    e.preventDefault();
    const feedback = feedbackText.value.trim();
    if (feedback) {
        // Send feedback to server (replace with actual API call)
        console.log('Feedback submitted:', feedback);
        alert('Thank you for your feedback!');
        feedbackText.value = '';
    }
});

// Add event listener for Enter key in chat input
userInput.addEventListener('keypress', (e) => {
    if (e.key === 'Enter') {
        sendMessage();
    }
});

// Login form submission (replace with actual authentication logic)
const loginForm = document.getElementById('loginForm');
loginForm.addEventListener('submit', (e) => {
    e.preventDefault();
    const username = loginForm.querySelector('input[type="text"]').value;
    const password = loginForm.querySelector('input[type="password"]').value;
    
    // Replace this with actual authentication logic
    console.log('Login attempt:', { username, password });
    alert('Login functionality not implemented yet.');
    
    loginModal.style.display = 'none';
    loginForm.reset();
});
const switcher = document.getElementById('switcher');
const container = document.querySelector('.container');
const registerForm = document.getElementById('registerForm');
const loginForm = document.getElementById('loginForm');
const heading = document.getElementById('heading');


let registerPage = false;

function togglePanel() {
    const switcherValue = switcher.checked;
    if (switcherValue) {
        registerForm.style.display = 'block';
        loginForm.style.display = 'none';
        container.classList.add('flipped');
        heading.textContent = "User Register";
        registerPage = true;
        clearFormFields(loginForm);
    } else {
        registerForm.style.display = 'none';
        loginForm.style.display = 'block';
        container.classList.remove('flipped');
        heading.textContent = "User Login";
        registerPage = false;
        clearFormFields(registerForm);
    }
}

async function submitForm(event) {
    event.preventDefault();
     let formData = new FormData();

    if (registerPage) {
        const newUsername = document.getElementById('newUsername').value;
        const newPassword = document.getElementById('newPassword').value;
        const repeatPassword = document.getElementById('repeatPassword').value;
        const studentID = document.getElementById('studentID').value;
        if (!newUsername) {
            alert('Please enter a value for New Username.');
            return;
        }

        if (!newPassword) {
            alert('Please enter a value for New Password.');
            return;
        }

        if (!repeatPassword) {
            alert('Please enter a value for Repeat Password.');
            return;
        }

        if (!studentID) {
            alert('Please enter a value for Student ID.');
            return;
        }

        if (newPassword !== repeatPassword) {
            alert('Please enter same password for both password input')
            return
        }

         formData.append('newUsername', newUsername);
         formData.append('newPassword', newPassword);
         formData.append('repeatPassword', repeatPassword);
         formData.append('studentID', studentID);

        fetch('http://localhost:8080/create-user', {
            method: 'POST',
            body: formData
        })
            .then(response => {
                if (response != null) {
                    return response.json(); // Extract the response body as JSON
                } else {
                    throw new Error('Failed to receive response from Go server');
                }
            })
            .then(data => {
                const message = data.message; // Extract the message text from the response JSON
                alert(message); // Display the message text in an alert dialog
            })
            .catch(error => {
                console.error('An error occurred while sending form data:', error);
            });

        clearFormFields(registerForm)

    } else {
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        if (!username) {
            alert('Please enter a value for Username.');
            return;
        }

        if (!password) {
            alert('Please enter a value for Password.');
            return;
        }

        formData.append('username', username);
        formData.append('password', password);

        fetch('http://localhost:8080/login', {
            method: 'POST',
            body: formData
        })
            .then(response => {
                if (response != null) {
                    return response.json(); // Extract the response body as JSON
                } else {
                    throw new Error('Failed to receive response from Go server');
                }
            })
            .then(data => {
                const message = data.message; // Extract the message text from the response JSON
                if (message !== 'login successful') {
                    alert(message)
                } else {
                    window.location.href = '../dashboard/index.html'
                    localStorage["data"] = data.username;
                    clearFormFields(loginForm);
                }
            })
            .catch(error => {
                console.error('An error occurred while sending form data:', error);
            });
    }
}

function clearFormFields(form) {
    const inputs = form.querySelectorAll('input[type="text"], input[type="password"]');
    inputs.forEach((input) => {
        input.value = '';
    });
}

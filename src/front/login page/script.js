const switcher = document.getElementById('switcher');
const container = document.querySelector('.container');
const registerForm = document.getElementById('registerForm');
const loginForm = document.getElementById('loginForm');
const heading = document.getElementById('heading');
/*--------------------------------------------------------------------------------------*/
const allData = {}
const usernames = []
/*--------------------------------------------------------------------------------------*/

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

         formData.append('newUsername', newUsername);
         formData.append('newPassword', newPassword);
         formData.append('repeatPassword', repeatPassword);
         formData.append('studentID', studentID);
        

        // TODO : fix these commented codes. then delete every thing between commented line /*---*/
     	fetch('http://localhost:8080/create-user', {
        	 method: 'POST',
       	  body: formData
     	})
         .then(response => {
             if (response.ok) {
                 return response.json(); // Parse response body as JSON
             } else {
                 throw new Error('Failed to receive response from Go server');
             }
         })
         .then(data => {
             // Handle the response data
             console.log(data.message); // Access the message field of the response
             return null
             // Perform additional actions based on the response
             // ...
         })
         .catch(error => {
             console.error('An error occurred while sending form data:', error);
             return null
         });

        /*--------------------------------------------------------------------------------------*/
        /*    if (usernames.includes(newUsername)) {
                alert('This username has already been taken! Please pick another one.');
                return;
            }
            if (newPassword !== repeatPassword) {
                alert('Two password fields are not the same.');
                return;
            }
            allData[newUsername] = [studentID, newPassword]
            usernames.push(newUsername)
        */
        /*--------------------------------------------------------------------------------------*/
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

        // formData.append('username', username);
        // formData.append('password', password);

        /*--------------------------------------------------------------------------------------*/
        /*
            if (!usernames.includes(username) || allData[username][0] !== password) {
                alert('Username or Password is not correct!')
                return;
            }
        */
        /*--------------------------------------------------------------------------------------*/
        window.location.href = '../dashboard/index.html'
        clearFormFields(loginForm)
    }


    /*--------------------------------------------------------------------------------------*/

    /*--------------------------------------------------------------------------------------*/
}

function clearFormFields(form) {
    const inputs = form.querySelectorAll('input[type="text"], input[type="password"]');
    inputs.forEach((input) => {
        input.value = '';
    });
}

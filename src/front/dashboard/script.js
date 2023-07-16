let roleTextElement = document.getElementById("roleText");
const usernameID = localStorage["data"];
let userData = {}

const profileNameText = document.getElementById("profile-name-text")

const usernameDefaultText = document.getElementById("username")

const studentIDDefaultText = document.getElementById("studentID")

let showProfile = false;
const profileForm = document.getElementById("profile-form");


function addToGroup() {
    let formData = new FormData()
    formData.append(userData.role)
    fetch('http://localhost:8080/add-groups', {
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
            if (userData.role === 0) {
                
            } else {

            }
        })
        .catch(_ =>{
            alert('An error occurred while sending form data:');
        });
}

function createCourse() {
    // Code to handle creating a course (only for admin)
}

function createCourseRequest() {
    // Code to handle creating a course request
}

function changeProfile() {
    showProfile = !showProfile;
    if (showProfile) {
        profileForm.style.display = 'block'
    } else {
        profileForm.style.display = 'none'
    }
    clearFormFields(profileForm)
}

function requestRole() {
    // Code to handle requesting a role change
}

function expandBox(boxNumber) {
    const box = document.querySelectorAll('.blue-box')[boxNumber - 1];
    box.classList.toggle('expanded');
}

function clearFormFields(form) {
    const inputs = form.querySelectorAll('input[type="text"], input[type="password"]');
    inputs.forEach((input) => {
        input.value = '';
    });
}

window.addEventListener('load', function() {
    // Call your initialization function here
    initializePage();
});

function initializePage() {
    let formData = new FormData()
    formData.append('username', usernameID);
    fetch('http://localhost:8080/get-user', {
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
            userData = data;
            profileNameText.textContent = userData.username;
            usernameDefaultText.placeholder = userData.username;
            studentIDDefaultText.placeholder = userData.number;
            switch (userData.role) {
                case 0:
                    roleTextElement.textContent = "Admin";
                    break;
                case 1:
                    roleTextElement.textContent = "G.O";
                    break;
                case 2:
                    roleTextElement.textContent = "S.G.O";
                    break;
                default:
                    roleTextElement.textContent = "Guest";
                    break
            }
        })
        .catch(_ => {
            alert('An error occurred while sending form data:');
        });
}

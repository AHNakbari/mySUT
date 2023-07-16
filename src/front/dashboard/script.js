const role = "Admin";
const roleTextElement = document.getElementById("roleText");
const usernameID = localStorage["data"];
roleTextElement.textContent = role;
let userData = {
    "username": "",
    "role": "",
    "groups": [],
    "courses": [],
    "field": "",
    "year": ""
};

let profileNameText = document.getElementById("profile-name-text")

const usernameDefault = "Amirhossein"
const usernameDefaultText = document.getElementById("username")
usernameDefaultText.value = usernameDefault

const studentIDDefault = "400104737"
const studentIDDefaultText = document.getElementById("studentID")
studentIDDefaultText.value = studentIDDefault

let showProfile = false;
const profileForm = document.getElementById("profile-form");


function createGroup() {
    // Code to handle creating a group (only for admin)
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
            userData.role = data.role;
            // part 1
            userData.username = data.username;
            userData.year = data.year;
            userData.field = data.field;
            userData.courses = data.courses;
            userData.groups = data.groups;
        })
        .catch(error => {
            alert('An error occurred while sending form data:');
        });
    // part 2
    alert(userData.role)
    profileNameText.textContent = userData.username;
}

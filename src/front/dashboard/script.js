const role = "Admin";
const roleTextElement = document.getElementById("roleText");
roleTextElement.textContent = role;

const profileName = "Amirhossein";
const profileNameText = document.getElementById("profile-name-text")
profileNameText.textContent = profileName;

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
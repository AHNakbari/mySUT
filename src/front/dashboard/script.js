const role = "Admin";
const roleTextElement = document.getElementById("roleText");
roleTextElement.textContent = role;

const profileName = "Amirhossein";
const profileNameText = document.getElementById("profile-name-text")
profileNameText.textContent = profileName;

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
    // Code to handle changing the profile
}

function requestRole() {
    // Code to handle requesting a role change
}

function expandBox(boxNumber) {
    const box = document.querySelectorAll('.blue-box')[boxNumber - 1];
    box.classList.toggle('expanded');
}
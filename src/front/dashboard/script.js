let roleTextElement = document.getElementById("roleText");
const usernameID = localStorage["data"];
let userData = {}

const profileNameText = document.getElementById("profile-name-text")
const userGroupList = document.getElementById('userGroupList');
const userCourseList = document.getElementById('userCourseList');
const usernameDefaultText = document.getElementById("username")

const allGroups = document.getElementById('all-groups-list')
const studentIDDefaultText = document.getElementById("studentID")

let showProfile = false;
const profileForm = document.getElementById("profile-form");
const creteGroup = document.getElementById("create-group");
const group = document.getElementById("group");
const owner = document.getElementById("owner");


function addToGroup() {
    profileForm.style.display = 'none'
    let formData = new FormData()
    fetch('http://localhost:8080/get-groups', {
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
            allGroups.innerHTML = '';
            console.log(data.groups)
            data.groups[1] = "Computer Science"
            data.groups.forEach(function (string) {
                const li = document.createElement('li');
                li.textContent = string;
                allGroups.appendChild(li);
            });

            allGroups.style.display = 'block';
            if (userData.role === 0) {
                creteGroup.style.display = 'block'
            }
        })
        .catch(_ => {
            alert('An error occurred while sending form data:');
        });
}

function createGroups() {
    let formData = new FormData()
    formData.append("name", group.value)
    formData.append("owner", owner.value)
    fetch('http://localhost:8080/create-group', {
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
        })
        .catch(_ => {
            alert('An error occurred while sending form data:');
        });
}

function createCourseRequest() {
    // Code to handle creating a course request
}

function changeProfile() {
    showProfile = !showProfile;
    if (showProfile) {
        allGroups.style.display = 'none';
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
    if (boxNumber === 3) {
        userGroupList.innerHTML = '';
        box.classList.toggle('expanded');

        // Create list items and append them to the ul element
        userData.groups.forEach(function (string) {
            const li = document.createElement('li');
            li.textContent = string;
            userGroupList.appendChild(li);
        });

        userGroupList.style.display = 'block';
    } else if (boxNumber === 4) {
        userCourseList.innerHTML = '';
        box.classList.toggle('expanded');
        console.log(userData.courses)
        // Create list items and append them to the ul element
        userData.courses.forEach(function (string) {
            const li = document.createElement('li');
            li.textContent = string;
            userCourseList.appendChild(li);
        });

        userCourseList.style.display = 'block';
    } else if (boxNumber === 1) {
        box.classList.toggle('expanded');
    }

}

function clearFormFields(form) {
    const inputs = form.querySelectorAll('input[type="text"], input[type="password"]');
    inputs.forEach((input) => {
        input.value = '';
    });
}

window.addEventListener('load', function () {
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

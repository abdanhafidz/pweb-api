// Configuration
const API_BASE_URL = 'http://api-pweb.abdanhafidz.com';
const TOKEN_COOKIE_NAME = 'auth_token';

// Utility Functions
function setCookie(name, value, days) {
    let expires = '';
    if (days) {
        const date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        expires = '; expires=' + date.toUTCString();
    }
    document.cookie = name + '=' + (value || '') + expires + '; path=/';
}

function getCookie(name) {
    const nameEQ = name + '=';
    const ca = document.cookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') c = c.substring(1, c.length);
        if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length);
    }
    return null;
}

function eraseCookie(name) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

function showLoading() {
    $('#loadingIndicator').show();
}

function hideLoading() {
    $('#loadingIndicator').hide();
}

function showAlert(elementId, type, message) {
    const alertElement = $(`#${elementId}`);
    alertElement.attr('class', `alert alert-${type}`);
    alertElement.html(message);
    alertElement.show();
    
    // Auto hide after 5 seconds
    setTimeout(() => {
        alertElement.hide();
    }, 5000);
}

function isLoggedIn() {
    return !!getCookie(TOKEN_COOKIE_NAME);
}

function updateNavLinks() {
    const navLinks = $('#navLinks');
    navLinks.empty();

    if (isLoggedIn()) {
        navLinks.append('<a href="profile.html" id="navProfile">Profile</a>');
        navLinks.append('<a href="#" id="navLogout">Logout</a>');
    } else {
        navLinks.append('<a href="login.html" id="navLogin">Login</a>');
        navLinks.append('<a href="register.html" id="navRegister">Register</a>');
    }

    // Attach event listeners to nav links
    $('#navLogout').on('click', function(e) {
        e.preventDefault();
        logout();
    });
}

function setAuthToken(token) {
    setCookie(TOKEN_COOKIE_NAME, token, 7); // Store token for 7 days
}

function getAuthToken() {
    return getCookie(TOKEN_COOKIE_NAME);
}

function getAuthHeader() {
    const token = getAuthToken();
    return token ? { 'Authorization': `Bearer ${token}` } : {};
}

// Auth Functions
function login(email, password) {
    showLoading();
    
    $.ajax({
        url: `${API_BASE_URL}/auth/login`,
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({
            email: email,
            password: password
        }),
        success: function(response) {
            if (response.status === 'success') {
                setAuthToken(response.data.token);
                showAlert('loginAlert', 'success', 'Login successful! Redirecting to your profile...');
                setTimeout(() => {
                    window.location.href = 'profile.html';
                }, 1000);
            } else {
                showAlert('loginAlert', 'danger', 'Login failed. Please check your credentials.');
            }
        },
        error: function(xhr) {
            const message = xhr.responseJSON ? xhr.responseJSON.message : 'An error occurred during login.';
            showAlert('loginAlert', 'danger', message);
        },
        complete: function() {
            hideLoading();
        }
    });
}

function register(email, password) {
    showLoading();
    
    $.ajax({
        url: `${API_BASE_URL}/auth/register`,
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({
            email: email,
            password: password
        }),
        success: function(response) {
            if (response.status === 'success') {
                showAlert('registerAlert', 'success', 'Registration successful! Redirecting to login page...');
                setTimeout(() => {
                    window.location.href = 'login.html';
                }, 2000);
            } else {
                showAlert('registerAlert', 'danger', 'Registration failed. Please try again.');
            }
        },
        error: function(xhr) {
            const message = xhr.responseJSON ? xhr.responseJSON.message : 'An error occurred during registration.';
            showAlert('registerAlert', 'danger', message);
        },
        complete: function() {
            hideLoading();
        }
    });
}

function logout() {
    eraseCookie(TOKEN_COOKIE_NAME);
    window.location.href = 'login.html';
}

function fetchUserProfile() {
    showLoading();
    
    $.ajax({
        url: `${API_BASE_URL}/user/me`,
        type: 'GET',
        headers: getAuthHeader(),
        success: function(response) {
            if (response.status === 'success') {
                const userData = response.data;
                
                // Update profile email
                $('#profileEmail').text(userData.account.email);
                
                // Update form fields if details exist
                if (userData.details) {
                    $('#profileFullName').val(userData.details.full_name || '');
                    $('#profileInitialName').val(userData.details.initial_name || '');
                    $('#profileUniversity').val(userData.details.university || '');
                    $('#profilePhone').val(userData.details.phone_number || '');
                    
                    // Update profile name and initials
                    if (userData.details.full_name) {
                        $('#profileName').text(userData.details.full_name);
                    } else {
                        $('#profileName').text('User Profile');
                    }
                    
                    if (userData.details.initial_name) {
                        $('#profileInitials').text(userData.details.initial_name.substring(0, 2).toUpperCase());
                    } else if (userData.details.full_name) {
                        const nameParts = userData.details.full_name.split(' ');
                        if (nameParts.length > 1) {
                            $('#profileInitials').text((nameParts[0][0] + nameParts[1][0]).toUpperCase());
                        } else {
                            $('#profileInitials').text(nameParts[0][0].toUpperCase());
                        }
                    } else {
                        $('#profileInitials').text(userData.account.email[0].toUpperCase());
                    }
                } else {
                    $('#profileInitials').text(userData.account.email[0].toUpperCase());
                }
            } else {
                showAlert('profileAlert', 'danger', 'Failed to load profile data.');
            }
        },
        error: function(xhr) {
            if (xhr.status === 401) {
                eraseCookie(TOKEN_COOKIE_NAME);
                window.location.href = 'login.html';
                showAlert('loginAlert', 'danger', 'Session expired. Please login again.');
            } else {
                const message = xhr.responseJSON ? xhr.responseJSON.message : 'An error occurred while fetching profile data.';
                showAlert('profileAlert', 'danger', message);
            }
        },
        complete: function() {
            hideLoading();
        }
    });
}

function updateUserProfile(profileData) {
    showLoading();
    
    $.ajax({
        url: `${API_BASE_URL}/user/me`,
        type: 'PUT',
        headers: getAuthHeader(),
        contentType: 'application/json',
        data: JSON.stringify({
            full_name: profileData.fullName,
            initial_name: profileData.initialName,
            university: profileData.university,
            phone_number: profileData.phoneNumber
        }),
        success: function(response) {
            if (response.status === 'success') {
                showAlert('profileAlert', 'success', 'Profile updated successfully!');
                fetchUserProfile(); // Refresh profile data
            } else {
                showAlert('profileAlert', 'danger', 'Failed to update profile data.');
            }
        },
        error: function(xhr) {
            if (xhr.status === 401) {
                eraseCookie(TOKEN_COOKIE_NAME);
                window.location.href = 'login.html';
                showAlert('loginAlert', 'danger', 'Session expired. Please login again.');
            } else {
                const message = xhr.responseJSON ? xhr.responseJSON.message : 'An error occurred while updating profile data.';
                showAlert('profileAlert', 'danger', message);
            }
        },
        complete: function() {
            hideLoading();
        }
    });
}

// Check Authentication State
$(document).ready(function() {
    updateNavLinks();
    
    // Handle logo click
    $('.logo').on('click', function(e) {
        e.preventDefault();
        if (isLoggedIn()) {
            window.location.href = 'profile.html';
        } else {
            window.location.href = 'login.html';
        }
    });
});
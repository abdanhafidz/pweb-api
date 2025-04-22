
# AJAX Implementation in UserPortal Web Application Documentation
## BY : ABDAN HAFIDZ (5054231 021)
# You Can See The result Below 
- [Front-End/view Link HERE](https://abdanhafidz.com/user-portal/)
- [Back-End/Documentation AUTHENTICATION HERE](https://documenter.getpostman.com/view/13117366/2sB2iwHFsn)
- [Back-End/Documentation USER HERE](https://documenter.getpostman.com/view/13117366/2sB2ixiD8t)
## Project Overview
![OVERVIEW](https://github.com/user-attachments/assets/9311ea91-8c3c-40ff-b285-a9eda887e298)

UserPortal is a web application built with frontend technologies (HTML, CSS, jQuery) that communicates with a backend API developed in Golang. The application features user authentication, registration, and profile management functionalities.

## Architecture Diagram

```mermaid
graph TD
    subgraph "Frontend"
        UI[HTML/CSS/jQuery] -->|AJAX| API
    end
    
    subgraph "Backend Infrastructure"
        API[Golang API] -->|SQL Queries| DB[(PostgreSQL)]
        GH[GitHub] -->|CI/CD Pipeline| HF[HuggingFace]
        HF -->|Deploys| API
    end
    
    subgraph "Database Infrastructure"
        DB -->|Connection Pool| SB[Supabase Pooler]
        SB -->|Managed Service| AWS[AWS RDS]
    end
    
    Users[Users] -->|Interacts| UI
```


```mermaid
sequenceDiagram
    participant User
    participant Frontend
    participant API as Golang API
    participant DB as PostgreSQL
    
    User->>Frontend: Enter credentials
    Frontend->>API: POST /auth/login
    API->>DB: Verify credentials
    DB->>API: Credentials valid
    API->>API: Generate JWT token
    API->>Frontend: Return token in response
    Frontend->>Frontend: Store token in cookies
    Frontend->>User: Redirect to profile page
    
    User->>Frontend: Access profile page
    Frontend->>Frontend: Check for token
    Frontend->>API: GET /user/me with Bearer token
    API->>API: Validate JWT token
    API->>DB: Fetch user data
    DB->>API: Return user data
    API->>Frontend: Return user profile
    Frontend->>User: Display profile data
```


```mermaid
graph TD
    A[Developer] -->|Push Code| B[GitHub Repository]
    B -->|Trigger Workflow| C[GitHub Actions]
    
    subgraph "CI Pipeline"
        C -->|Checkout Code| D[Setup Go Environment]
        D -->|Install Dependencies| E[Run Tests]
        E -->|Build Binary| F[Create Docker Image]
    end
    
    subgraph "CD Pipeline"
        F -->|Login to Registry| G[Push to Container Registry]
        G -->|Deploy to| H[HuggingFace Spaces]
        H -->|Update| I[Running Service]
    end
    
    subgraph "Database Infrastructure"
        J[Supabase Connection Pooler] -->|Manages Connections| K[(AWS PostgreSQL Database)]
        I -->|Connect via| J
    end
    
    style A fill:#f96,stroke:#333,stroke-width:1px
    style B fill:#f96,stroke:#333,stroke-width:1px
    style C fill:#69f,stroke:#333,stroke-width:1px
    style F fill:#69f,stroke:#333,stroke-width:1px
    style H fill:#9f6,stroke:#333,stroke-width:1px
    style I fill:#9f6,stroke:#333,stroke-width:1px
    style K fill:#f69,stroke:#333,stroke-width:1px
```

## Technology Stack

### Frontend
- **HTML5**: Structure of the web application
- **CSS3**: Styling and responsive design
- **jQuery**: DOM manipulation and AJAX requests
- **JavaScript**: Client-side logic and form validations

### Backend


- **Golang**: Server-side API development
- **PostgreSQL**: Database for storing user data
- **Supabase**: Database connection pooling 
- **AWS RDS**: Managed database service
- **HuggingFace**: Deployment platform
- **GitHub Actions**: CI/CD pipeline

## Frontend Implementation

The frontend is built as a multi-page application with separate HTML files for each main function, linked by a common JavaScript file for shared functionality.

### Directory Structure

```
├── index.html
├── login.html
├── register.html
├── profile.html
├── styles/ styles.css
└── scripts/ scripts.js
```

### JavaScript Implementation

The main JavaScript file (`scripts.js`) handles all client-side logic including API communication, authentication, and user interface updates.

#### Core Functions

##### Authentication Management

```javascript
// Set JWT token in cookies
function setAuthToken(token) {
    setCookie(TOKEN_COOKIE_NAME, token, 7); // Store token for 7 days
}

// Get JWT token from cookies
function getAuthToken() {
    return getCookie(TOKEN_COOKIE_NAME);
}

// Add authorization header to API requests
function getAuthHeader() {
    const token = getAuthToken();
    return token ? { 'Authorization': `Bearer ${token}` } : {};
}

// Check if user is authenticated
function isLoggedIn() {
    return !!getCookie(TOKEN_COOKIE_NAME);
}
```

##### API Communication

The application uses jQuery AJAX to communicate with the backend API:

```javascript
// Login function
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
                // Redirect after successful login
                window.location.href = 'profile.html';
            } else {
                showAlert('loginAlert', 'danger', 'Login failed.');
            }
        },
        error: function(xhr) {
            // Error handling
        },
        complete: function() {
            hideLoading();
        }
    });
}
```

##### User Profile Management

```javascript
// Fetch user profile data
function fetchUserProfile() {
    showLoading();
    
    $.ajax({
        url: `${API_BASE_URL}/user/me`,
        type: 'GET',
        headers: getAuthHeader(),
        success: function(response) {
            if (response.status === 'success') {
                // Update UI with user data
                updateProfileUI(response.data);
            } else {
                showAlert('profileAlert', 'danger', 'Failed to load profile.');
            }
        },
        error: function(xhr) {
            // Handle unauthorized access
            if (xhr.status === 401) {
                eraseCookie(TOKEN_COOKIE_NAME);
                window.location.href = 'login.html';
            }
        },
        complete: function() {
            hideLoading();
        }
    });
}

// Update user profile data
function updateUserProfile(profileData) {
    showLoading();
    
    $.ajax({
        url: `${API_BASE_URL}/user/me`,
        type: 'PUT',
        headers: getAuthHeader(),
        contentType: 'application/json',
        data: JSON.stringify(profileData),
        success: function(response) {
            if (response.status === 'success') {
                showAlert('profileAlert', 'success', 'Profile updated!');
                fetchUserProfile(); // Refresh data
            }
        },
        error: function(xhr) {
            // Error handling
        },
        complete: function() {
            hideLoading();
        }
    });
}
```

##### Navigation and UI Management

```javascript
// Update navigation based on auth state
function updateNavLinks() {
    const navLinks = $('#navLinks');
    navLinks.empty();

    if (isLoggedIn()) {
        navLinks.append('<a href="profile.html">Profile</a>');
        navLinks.append('<a href="#" id="navLogout">Logout</a>');
        
        // Attach logout handler
        $('#navLogout').on('click', function(e) {
            e.preventDefault();
            logout();
        });
    } else {
        navLinks.append('<a href="login.html">Login</a>');
        navLinks.append('<a href="register.html">Register</a>');
    }
}
```

## Backend Implementation

The backend is built with Golang and deployed on HuggingFace using GitHub Actions for continuous integration and delivery.

### API Endpoints
I"ve Deployed the Backend at (Hugging Face's Space) [https://lifedebugger-pweb-api.hf.space/api/v1]
| Endpoint | Method | Description | Authentication Required |
|----------|--------|-------------|------------------------|
| `/auth/login` | POST | User login | No |
| `/auth/register` | POST | User registration | No |
| `/user/me` | GET | Get user profile | Yes |
| `/user/me` | PUT | Update user profile | Yes |

### Authentication Flow

1. User submits login credentials
2. Backend validates credentials against database
3. Backend generates JWT token
4. Frontend stores token in cookies
5. Frontend includes token in subsequent API requests

### Database Schema

```mermaid
erDiagram
    ACCOUNT {
        uuid UUID PK
        int id
        string email
        string password
        boolean is_email_verified
        boolean is_detail_completed
        timestamp created_at
        timestamp deleted_at
    }
    
    USER_DETAILS {
        int id PK
        int account_id FK
        string initial_name
        string full_name
        string university
        string phone_number
    }
    
    ACCOUNT ||--o| USER_DETAILS : "has"
```

## Deployment Architecture

### CI/CD Pipeline

```mermaid
flowchart LR
    A[Developer] -->|Push Code| B[GitHub Repository]
    B -->|Trigger| C[GitHub Actions]
    C -->|Run Tests| D{Tests Pass?}
    D -->|Yes| E[Build Go Binary]
    E -->|Deploy| F[HuggingFace]
    D -->|No| G[Notify Developer]
```

### Database Infrastructure

- **PostgreSQL**: Primary database running on AWS RDS
- **Supabase**: Provides connection pooling for efficient database management
- **Connection Flow**: Application → Supabase Pooler → AWS RDS PostgreSQL

## Security Implementation

1. **Authentication**: JWT-based authentication with token expiration
2. **Password Security**: Passwords stored as secure hashes, not plaintext
3. **CORS Protection**: Properly configured for API access
4. **Input Validation**: Both client-side and server-side validation

## Future Enhancements

1. Implement email verification
2. Add two-factor authentication
3. Expand user profile with more details
4. Add user avatar upload functionality
5. Implement password reset functionality

---

## Appendix: API Response Formats

### Login Response

```json
{
    "status": "success",
    "message": "Data retrieved successfully!",
    "data": {
        "account": {
            "id": 61,
            "uuid": "63983d79-386a-44ed-9c60-1fbc443992f8",
            "email": "abdan.hafidz21@gmail.com",
            "password": "SECRET",
            "is_email_verified": false,
            "is_detail_completed": false,
            "created_at": "2025-04-22T16:06:29.912177Z",
            "deleted_at": null
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJxb2JpbHR1LmlkIiwiZXhwIjoxNzQ1NDI0NDEwLCJpYXQiOjE3NDUzMzgwMTAsImlkIjo2MX0.BpbYCAjENP2zBAb8t_cZA5HIeA6bdZI0VKjE95ki-Qk"
    },
    "meta_data": {}
}
```

### User Profile Response

```json
{
    "status": "success",
    "message": "Data retrieved successfully!",
    "data": {
        "account": {
            "id": 35,
            "uuid": "1c47068a-f65d-43df-85c9-8436af4d9dd7",
            "email": "lolmeng267@gmail.com",
            "password": "SECRET",
            "is_email_verified": true,
            "is_detail_completed": false,
            "created_at": "2025-04-20T10:46:20.182564+07:00",
            "deleted_at": null
        },
        "details": {
            "id": 16,
            "account_id": 35,
            "initial_name": "",
            "full_name": null,
            "university": null,
            "phone_number": null
        }
    },
    "meta_data": {}
}
```

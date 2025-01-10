# Project: Connect+ - Inclusive Dating App - Current Status: Backend Development - User Creation

**Description:**

Connect+ is an inclusive dating app designed for users of all sexual orientations and gender identities. It prioritizes meaningful connections, safety, and an engaging user experience.

**Key Features:**

*   **Inclusive Onboarding:** Detailed signup process capturing sexual orientation, gender identity, and relationship preferences.
*   **Advanced Matching Algorithm:** Prioritizes compatibility based on personality, interests, and relationship goals.
*   **Swipe Functionality:** Familiar swipe left (not interested) / swipe right (interested) interaction for browsing profiles.
*   **Robust Verification:** Multi-factor authentication and photo verification to ensure profile authenticity.
*   **Tiered Subscription Model:** Offers free and paid tiers with varying features and benefits.
*   **Safety & Moderation:** Strong safety features, community guidelines, and active moderation to create a respectful environment.
*   **Community Features:** Groups, forums, and events for users to connect based on shared interests and identities.

**Technical Specifications:**

*   **Frontend:** Flutter with Dart (for cross-platform compatibility - iOS, Android, Web)
*   **Backend:** Go (Golang) (for API development)
*   **Database:** PostgreSQL (for data storage)
*   **Cloud Hosting:** AWS (for scalability and reliability)
*   **Authentication:** JWT (JSON Web Tokens) for secure user authentication.

**Rules File:** rules.md

**Development Stages:**

1.  **Backend Development:**
    * Build the API for user management, matching, messaging, etc., using Go.  
    * A basic Go server is set up with a root handler (`/`) and a user handler (`/user`).
    * The `user.go` file has been removed as it had duplicate information in the `main.go`.   
    * We have created a User struct in `main.go` with ID, Username, and Email fields.
    * The rootHandler now prints "Welcome to Connect+! The API is running!" instead of "Welcome to Connect+".
    * The userHandler now prints "User endpoint hit!" instead of "User Path".
    * A createUserHandler function has been added to handle POST requests to /user and it prints "User created!".
    * The main() function now routes POST requests to /user to the createUserHandler.

2.  **Frontend Development:** Create the user interface with Flutter, integrating swipe functionality and other features.
3.  **Database Design:** Structure the database to efficiently store user data, profiles, and interactions.
4.  **Authentication & Security:** Implement secure authentication and authorization mechanisms.
5.  **Matching Algorithm:** Develop and refine the algorithm for accurate and personalized matching.
6.  **Testing & Deployment:** Thorough testing and deployment to app stores and web.

**Success Metrics:**

*   User acquisition and retention rates
*   User engagement (messages sent, profiles viewed, etc.)
*   Subscription conversion rates
*   User satisfaction (ratings and reviews)
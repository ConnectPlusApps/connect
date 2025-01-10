# Connect+ - Cline Rules

**General Rules:**

* **Adhere to Project Scope:**  Focus on the features and functionalities defined in `project.md`.
* **Prioritize User Privacy:**  Implement data protection measures and comply with privacy regulations.
* **Maintain Code Quality:**  Write clean, well-documented, and efficient code.
* **Follow Best Practices:**  Use industry-standard coding conventions and best practices for Go API development and Flutter frontend development.

**Specific Rules:**

* **Authentication:**  Implement JWT authentication for all API endpoints requiring user authorization.
* **Database Interactions:**  Use parameterized queries to prevent SQL injection vulnerabilities.
* **Error Handling:**  Implement robust error handling and logging mechanisms.
* **API Documentation:**  Generate clear and concise API documentation using Swagger or a similar tool.
* **Frontend Development:**  Use Flutter best practices for performance optimization and cross-platform compatibility.
* **Concurrency (Go):**  Leverage Go's concurrency features (goroutines, channels) for efficient handling of concurrent requests.

**Forbidden Actions:**

* **Do not introduce features not listed in `project.md` without explicit confirmation.**
* **Do not use third-party libraries or frameworks without approval.**
* **Do not generate code that violates user privacy or data security.**
* **Do not deviate from the specified tech stack (Flutter, Go, PostgreSQL).**

**Code Style Guide:**

* **Go:** Follow the Go Code Review Comments style guide (https://github.com/golang/go/wiki/CodeReviewComments)
* **Flutter (Dart):** Follow the Dart Style Guide (https://dart.dev/guides/language/effective-dart/style)

**Reporting:**

* If Cline encounters any issues or requires clarification, it should clearly state the problem and request further instructions.
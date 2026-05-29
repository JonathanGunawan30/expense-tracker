# Expense Tracker RESTful API

A simple RESTful API for managing personal expenses and categories, built with Go following Clean Architecture principles.

## Features
- **User Management**: Register and list users.
- **Category Management**: Organize expenses into customizable categories with full CRUD operations.
- **Expense Tracking**: Detailed expense logging including amount, title, and description with full CRUD operations.
- **Security**:
    - **API Key Authentication**: Secure access via `X-API-Key` header.
    - **Ownership Protection**: Users can only access, modify, or delete their own categories and expenses (enforced via `X-User-ID` header).
- **Input Validation**: Robust request validation using `go-playground/validator`.
- **Clean Architecture**: Decoupled layers (Delivery, Usecase, Domain, Infrastructure) for maintainability and testability.
- **Configuration Management**: Flexible environment-based configuration using `viper`.

## Tech Stack
- **Language**: Go 1.26+
- **Router**: `httprouter`
- **Database**: MySQL with GORM
- **Logging**: `logrus`
- **Configuration**: `viper`
- **Validation**: `go-playground/validator`
- **Testing**: JetBrains HTTP Client (`.http` files)

## Prerequisites
- Go 1.26 or later
- MySQL Database

## Setup
1. **Clone the repository**:
   ```bash
   git clone https://github.com/JonathanGunawan30/expense-tracker
   cd expense-tracker
   ```
2. **Environment Configuration**:
   Create a `.env` file based on `.env.example`:
   ```bash
   cp .env.example .env
   ```
   Update the database credentials and `X_API_KEY` in the `.env` file.

3. **Database Setup**:
    - Create a database named `expense_tracker` in MySQL.
    - Run the migrations found in `database/migrations/` (Up SQL files).

4. **Run the Application**:
   ```bash
   go run cmd/main.go
   ```
   The server will start on port `3000` (default).

## Testing
Manual testing can be performed using the JetBrains HTTP Client.
- Open `test/manual-test.http`.
- The script automatically handles `user_id`, `category_id`, and `expense_id` variables after successful creation requests.
- Ensure your `@api_key` in the `.http` file matches your `.env` configuration.

## API Documentation
The complete API specification is available in `docs/api.json` (OpenAPI format).

### Key Endpoints:
- **Users**:
    - `POST /api/users` - Register a new user
    - `GET /api/users` - List all users
- **Categories**:
    - `POST /api/categories` - Create a category
    - `GET /api/categories` - List all user's categories
    - `GET /api/categories/:categoryID` - Get category details
    - `PUT /api/categories/:categoryID` - Update a category
    - `DELETE /api/categories/:categoryID` - Delete a category
- **Expenses**:
    - `POST /api/expenses` - Create an expense
    - `GET /api/expenses` - List all user's expenses
    - `GET /api/expenses/:expenseID` - Get expense details
    - `PUT /api/expenses/:expenseID` - Update an expense
    - `DELETE /api/expenses/:expenseID` - Delete an expense

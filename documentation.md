A **cart**, often referred to as a **shopping cart** or **basket**, is a fundamental feature in e-commerce platforms that allows users to select and temporarily store products they intend to purchase. It serves as a virtual container where users can accumulate items while they continue browsing the online store. Once they are ready to make a purchase, they proceed to checkout, where they can review, modify, and finalize their order.

### Key Features of a Shopping Cart

1. **Product Listing**: The cart displays all the items that a user has added, typically showing details like product name, quantity, price, and any options like size or color.

2. **Quantity Management**: Users can adjust the quantity of each item, potentially removing items from the cart or changing the number of units they wish to purchase.

3. **Price Calculation**: The cart calculates the total cost of all items, including any applicable taxes, discounts, or shipping fees.

4. **Persistence**: The cart is often persistent, meaning that the items remain in the cart even if the user leaves the website or logs out, allowing them to return later to complete their purchase.

5. **Checkout Process**: The cart links to the checkout process, where users can enter shipping information, select payment methods, and finalize their order.

6. **User Authentication**: While some carts are accessible to guest users, others may require users to log in to save their cart for future sessions or access additional features.

7. **Promotional Codes**: Carts may include the ability to apply discount codes or promotions, which can alter the total price.

### Implementation Considerations

- **User Experience**: A good cart design is crucial for a positive user experience, providing clear feedback, easy navigation, and a seamless transition to checkout.

- **Backend Management**: On the backend, carts can be managed using sessions for guest users or databases for logged-in users, allowing carts to persist across sessions.

- **Security**: It's important to secure the cart to prevent issues like price manipulation or unauthorized access.

In an e-commerce backend project, implementing a shopping cart involves creating APIs to handle adding items, updating quantities, and removing items, as well as calculating totals and managing discounts. It's a central component of the purchasing process and can significantly impact the user's shopping experience and the store's conversion rate.

# Has One vs Belongs To

In GORM (Go Object-Relational Mapping), "has one" and "belongs to" associations describe the relationships between models. Here’s a clear breakdown with illustrations to highlight the differences:

### 1. **Has One Association**

A "has one" association means that a model has exactly one associated model. This is a one-to-one relationship where one record in the parent model is associated with one record in the child model.

#### Illustration

Consider the following example where a `User` has one `Profile`.

```go
// Profile model
type Profile struct {
    gorm.Model
    UserID int    `json:"user_id" gorm:"unique;not null"`
    Bio    string `json:"bio"`
    User   User   `json:"user" gorm:"foreignKey:UserID"` // "has one" relationship
}

// User model
type User struct {
    gorm.Model
    Name    string   `json:"name" gorm:"not null"`
    Email   string   `json:"email" gorm:"unique;not null"`
    Profile Profile  `json:"profile" gorm:"constraint:OnDelete:CASCADE;"`
}
```

#### Explanation

- **User Model**: Has a `Profile` field which signifies that each `User` has one associated `Profile`.
- **Profile Model**: Has a `UserID` field that references the `User` model. The `foreignKey:UserID` tag indicates that `Profile` belongs to `User`.

In the database, this translates to:

- `profiles` table will have a `user_id` column which is a foreign key referring to the `users` table.
- Deleting a `User` will automatically delete the associated `Profile` due to `OnDelete:CASCADE`.

### 2. **Belongs To Association**

A "belongs to" association indicates that a model belongs to another model. This is also a one-to-one relationship but focuses on the child model referencing the parent model.

#### Illustration

Consider the following example where an `Order` belongs to a `Customer`.

```go
// Order model
type Order struct {
    gorm.Model
    OrderNumber string   `json:"order_number" gorm:"not null"`
    CustomerID  int      `json:"customer_id" gorm:"not null"`
    Customer    Customer `json:"customer" gorm:"foreignKey:CustomerID"` // "belongs to" relationship
}

// Customer model
type Customer struct {
    gorm.Model
    Name    string   `json:"name" gorm:"not null"`
    Email   string   `json:"email" gorm:"unique;not null"`
    Orders  []Order  `json:"orders" gorm:"foreignKey:CustomerID"`
}
```

#### Explanation

- **Order Model**: Has a `CustomerID` field that refers to the `Customer` model. The `foreignKey:CustomerID` tag indicates that `Order` belongs to `Customer`.
- **Customer Model**: Has an `Orders` field which represents the collection of orders belonging to that customer.

In the database, this translates to:

- `orders` table will have a `customer_id` column which is a foreign key referring to the `customers` table.
- `Customer` can have multiple `Orders`.

### Summary

- **Has One**: Describes a one-to-one relationship where the parent model has a single child model. For example, a `User` has one `Profile`.
- **Belongs To**: Describes a one-to-one relationship where the child model belongs to a single parent model. For example, an `Order` belongs to a `Customer`.

In both cases, the foreign key is placed in the child model. The difference lies in which side is considered the "owner" of the relationship: the parent in a "has one" or the child in a "belongs to."

### Entering into a database

Use `psql -d ecommerce -U postgres`. The `-d` flag stands for the database name and the `-U` flag for the database user's username. [Visit for more](https://www.freecodecamp.org/news/manage-postgresql-with-psql/)

### Common Commands

`\l` -Lists the names of databases within the postgres interactive shell.
`\q` -Moves out of the postgres interactive shell to the terminal.
`\d` -Displays the database tables and their relations.

Generated the secret key using `openssl rand -base64 32`

### On JWT

A JWT (JSON Web Token) is composed of three parts: the header, the payload, and the signature. These parts are Base64URL encoded and separated by dots (`.`). To view the JSON format of the token, you need to decode the header and payload from Base64URL.

Here is how you can decode the provided JWT:

### Token Structure

- **Header**: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`
- **Payload**: `eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ`
- **Signature**: `SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`

### Decoding Header

The header part `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9` decodes to:

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

### Decoding Payload

The payload part `eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ` decodes to:

```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```

### Signature

The signature part `SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c` is created by signing the header and payload with the secret key. It is not typically decoded but used to verify the token's authenticity.

### Complete JSON Format

So, the complete JSON representation of the JWT is:

```json
{
  "header": {
    "alg": "HS256",
    "typ": "JWT"
  },
  "payload": {
    "sub": "1234567890",
    "name": "John Doe",
    "iat": 1516239022
  },
  "signature": "SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

### Explanation of Fields

- **Header**
  - `alg`: The algorithm used to sign the token, in this case, HMAC SHA-256 (`HS256`).
  - `typ`: The type of token, which is JWT.

- **Payload**
  - `sub`: Subject identifier (usually a user ID).
  - `name`: Name of the user.
  - `iat`: Issued at time, in Unix timestamp format.

- **Signature**
  - A hash created using the header, payload, and a secret key, ensuring the token hasn't been tampered with.

You can decode JWT tokens using various online tools or libraries in your programming language of choice to see their contents in a human-readable format.

### On setting RegisteredClaims fileds

When working with JWTs (JSON Web Tokens), setting or not setting fields in `RegisteredClaims` can have significant implications for security, token validation, and overall functionality. Here’s a detailed look at the implications of not setting these fields versus the benefits of setting them:

### Implications of Not Setting Fields

1. **`ExpiresAt` (`exp`)**:
   - **Implication**: If not set, the token will not have an expiration time, and it might be accepted indefinitely. This can pose a security risk because stolen tokens would remain valid indefinitely.
   - **Benefit of Setting**: Setting an expiration time ensures that tokens become invalid after a certain period, reducing the risk if a token is compromised. It enforces a "token lifespan" which is crucial for security.

2. **`Not Before` (`nbf`)**:
   - **Implication**: If not set, the token might be accepted immediately, even if it’s meant to be used only after a certain time. This could be problematic in scenarios where the token should not be valid before a specific time.
   - **Benefit of Setting**: Setting a "not before" time ensures that the token is only accepted after a specific point in time, which is useful for delayed access or when initializing tokens.

3. **`Issued At` (`iat`)**:
   - **Implication**: If not set, it’s unclear when the token was issued, which can complicate debugging and auditing. Some systems may use this to determine the age of the token.
   - **Benefit of Setting**: Provides a timestamp of when the token was issued, which can be useful for auditing, debugging, or implementing custom token age policies.

4. **`Issuer` (`iss`)**:
   - **Implication**: If not set, the token does not specify which system issued it, which might lead to issues in environments with multiple systems or services.
   - **Benefit of Setting**: Helps identify the issuer of the token. This is useful for validating that the token was issued by a trusted source and for handling tokens from multiple issuers.

5. **`Subject` (`sub`)**:
   - **Implication**: If not set, there is no clear identification of the principal (user or entity) the token is representing. This might affect systems that need to identify or handle specific users.
   - **Benefit of Setting**: Provides a unique identifier for the subject of the token, making it easier to handle user-specific logic and access control.

6. **`Audience` (`aud`)**:
   - **Implication**: If not set, the token may be accepted by any audience. This can be a security risk if you want to restrict the token to specific recipients.
   - **Benefit of Setting**: Restricts the token to specific audiences or applications. This ensures that the token is only valid for the intended recipient.

7. **`JWT ID` (`jti`)**:
   - **Implication**: If not set, the token does not have a unique identifier, which can make it harder to track or manage individual tokens, particularly in scenarios where token revocation or tracking is needed.
   - **Benefit of Setting**: Provides a unique identifier for the token, which is useful for implementing token revocation or managing tokens uniquely.

### On SSO Implementation

Integrating Single Sign-On (SSO) into your authentication method involves using a third-party authentication provider, such as Google, Facebook, or a custom OpenID Connect (OIDC) provider. Here, we'll focus on integrating Google SSO with your JWT-based authentication system in a Go application.

### Steps to Integrate Google SSO

#### 1. Set Up Google OAuth 2.0 Credentials

1. Go to the [Google Developers Console](https://console.developers.google.com/).
2. Create a new project or select an existing one.
3. Navigate to "Credentials" and create OAuth 2.0 Client IDs.
4. Set the authorized redirect URI to a route in your application that will handle the OAuth callback, e.g., `http://localhost:8000/auth/google/callback`.

#### 2. Install Required Packages

Install the necessary packages for handling OAuth 2.0.

```bash
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
```

#### 3. Configure OAuth 2.0

Create a file `config/oauth.go` to hold your OAuth 2.0 configuration.

```go
package config

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
    RedirectURL:  "http://localhost:8000/auth/google/callback",
    ClientID:     "your-google-client-id",
    ClientSecret: "your-google-client-secret",
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:     google.Endpoint,
}
```

#### 4. Implement Google SSO Handlers

Create a file `controllers/sso.go` to handle Google SSO.

```go
package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "yourprojectname/config"
    "yourprojectname/models"
    "yourprojectname/utils"
    "gorm.io/gorm"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/idtoken"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
    db = database
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
    url := config.GoogleOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
        return
    }

    idToken, ok := token.Extra("id_token").(string)
    if !ok {
        http.Error(w, "Failed to get ID token", http.StatusInternalServerError)
        return
    }

    payload, err := idtoken.Validate(context.Background(), idToken, config.GoogleOAuthConfig.ClientID)
    if err != nil {
        http.Error(w, "Failed to validate ID token", http.StatusInternalServerError)
        return
    }

    email, _ := payload.Claims["email"].(string)
    if email == "" {
        http.Error(w, "Email not found", http.StatusInternalServerError)
        return
    }

    var user models.User
    if err := db.Where("email = ?", email).First(&user).Error; err != nil {
        // User not found, create a new user
        user = models.User{
            Email: email,
            Role:  "customer", // Default role
        }
        if err := db.Create(&user).Error; err != nil {
            http.Error(w, "Failed to create user", http.StatusInternalServerError)
            return
        }
    }

    // Generate JWT token
    jwtToken, err := utils.GenerateToken(user.Email, user.Role)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Set the token in the response
    json.NewEncoder(w).Encode(map[string]string{"token": jwtToken})
}
```

#### 5. Update Routes to Include SSO

Update your `main.go` to include the new SSO routes.

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "yourprojectname/config"
    "yourprojectname/controllers"
    "yourprojectname/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&models.User{})
    controllers.InitDB(db)

    router := mux.NewRouter()

    router.HandleFunc("/auth/google/login", controllers.GoogleLogin).Methods("GET")
    router.HandleFunc("/auth/google/callback", controllers.GoogleCallback).Methods("GET")

    // Other routes...

    fmt.Println("Server is running on port 8000")
    http.ListenAndServe(":8000", router)
}
```

### Summary

- **Google OAuth Setup**: Configure Google OAuth 2.0 credentials and set up your project.
- **OAuth 2.0 Configuration**: Create and configure the OAuth 2.0 settings in your project.
- **SSO Handlers**: Implement handlers for initiating login and handling the callback.
- **Routes**: Update your routes to include the new SSO endpoints.

By following these steps, you can integrate Google SSO into your JWT-based authentication system. This can be adapted for other OAuth providers by changing the specific OAuth configurations and endpoints.



## On Pagination and Sorting

### Pagination

**Pagination** involves breaking a large set of data into smaller chunks or pages, so clients can request and view one page of data at a time. This helps improve performance and user experience by not overwhelming clients with too much data at once.

#### Key Concepts

1. **Page Number**: Specifies which page of results to return.
2. **Page Size (Limit)**: Defines how many results per page.

#### Implementing Pagination

1. **Extract Query Parameters**: Extract `page` and `limit` from the request URL.
2. **Calculate Offset**: Calculate the starting point for the query based on the page number and page size.
3. **Apply Pagination to Query**: Use GORM's `Offset` and `Limit` methods to paginate results.

**Example Code**

```go
func GetProducts(w http.ResponseWriter, r *http.Request) {
    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")

    page := 1
    limit := 10

    if pageStr != "" {
        var err error
        page, err = strconv.Atoi(pageStr)
        if err != nil {
            http.Error(w, "Invalid page number", http.StatusBadRequest)
            return
        }
    }
    if limitStr != "" {
        var err error
        limit, err = strconv.Atoi(limitStr)
        if err != nil {
            http.Error(w, "Invalid limit number", http.StatusBadRequest)
            return
        }
    }

    var products []models.Product
    offset := (page - 1) * limit
    if err := db.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}
```

### Sorting

**Sorting** allows clients to organize data in a specific order, based on one or more fields, and in ascending or descending order.

#### Key Concepts

1. **Sort Field**: The field by which to sort the data (e.g., `name`, `price`).
2. **Sort Order**: The direction of the sort (e.g., `asc` for ascending, `desc` for descending).

#### Implementing Sorting

1. **Extract Sort Parameters**: Extract `sort_by` and `order` from the request URL.
2. **Apply Sorting to Query**: Use GORM's `Order` method to sort results based on the specified field and order.

**Example Code**

```go
func GetProducts(w http.ResponseWriter, r *http.Request) {
    //pagination parameters
    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")

    //sorting parameters
    sortBy := r.URL.Query().Get("sort_by")
    order := r.URL.Query().Get("order")

    page := 1
    limit := 10
    if pageStr != "" {
        var err error
        page, err = strconv.Atoi(pageStr)
        if err != nil {
            http.Error(w, "Invalid page number", http.StatusBadRequest)
            return
        }
    }
    if limitStr != "" {
        var err error
        limit, err = strconv.Atoi(limitStr)
        if err != nil {
            http.Error(w, "Invalid limit number", http.StatusBadRequest)
            return
        }
    }

    if sortBy == "" {
        sortBy = "name"
    }
    if order == "" {
        order = "asc"
    }

    var products []models.Product
    offset := (page - 1) * limit
    query := db.Offset(offset).Limit(limit)
    if strings.ToLower(order) == "desc" {
        query = query.Order(sortBy + " desc")
    } else {
        query = query.Order(sortBy + " asc")
    }

    if err := query.Find(&products).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}
```

### Summary

1. **Pagination**:
   - **Page**: Which page to display.
   - **Limit**: Number of items per page.
   - **Calculate Offset**: `(page - 1) * limit`.
   - **Apply to Query**: `db.Offset(offset).Limit(limit)`.

2. **Sorting**:
   - **Sort Field**: Field to sort by.
   - **Sort Order**: Ascending or descending.
   - **Apply to Query**: `db.Order(sortBy + " asc")` or `db.Order(sortBy + " desc")`.


## Database sorting capability

The sorting feature in your API is typically implemented using the sorting capabilities provided by your database system rather than a specific sorting algorithm in your Go code.

### Database Sorting

When you use GORM with SQL databases like PostgreSQL or MySQL, the actual sorting is performed by the database engine. The database uses its internal sorting algorithms, which are optimized for performance. Common sorting algorithms used by databases include:

1. **Quick Sort**: A highly efficient, divide-and-conquer algorithm that is commonly used for its average-case performance.
2. **Merge Sort**: A stable sorting algorithm used by some databases for its predictable performance and stability.
3. **Heap Sort**: Sometimes used in databases for its space efficiency and good worst-case performance.

### How Sorting is Implemented

In your Go code using GORM, sorting is specified as part of the query:

```go
if strings.ToLower(order) == "desc" {
    query = query.Order(sortBy + " desc")
} else {
    query = query.Order(sortBy + " asc")
}
```

- **`query.Order(sortBy + " desc")`**: Specifies that the results should be sorted in descending order based on the `sortBy` field.
- **`query.Order(sortBy + " asc")`**: Specifies ascending order.

### Key Points

- **Database Engine Handles Sorting**: The sorting is handled by the database engine using its internal algorithms, not explicitly in your Go code.
- **Optimization**: Modern databases are optimized for sorting operations and often use advanced algorithms and indexing techniques to perform these operations efficiently.

### Summary

In your Go application, you specify how to sort data (by which field and in what order) using GORM. The actual sorting is performed by the database system, which uses efficient sorting algorithms internally. This offloads the sorting workload from your application code to the database engine, leveraging its optimized performance capabilities.

## The `Order` Method

In the provided Go code with GORM, the sorting is handled by the `Order` method of the GORM query builder. Here's the relevant part of the code that performs sorting:

```go
if strings.ToLower(order) == "desc" {
    query = query.Order(sortBy + " desc")
} else {
    query = query.Order(sortBy + " asc")
}
```

### Breakdown of the Sorting Process

#### 1. **Order Method**
```go
query = query.Order(sortBy + " desc")
```
or
```go
query = query.Order(sortBy + " asc")
```
- **Purpose**: The `Order` method is used to specify the sort order for the query results.
- **Usage**: `query.Order(sortBy + " desc")` means that the results should be sorted in descending order based on the field specified by `sortBy`. Similarly, `query.Order(sortBy + " asc")` sorts the results in ascending order.
- **Parameter**: The parameter passed to `Order` is a string that specifies the column to sort by and the direction (`asc` or `desc`).

#### 2. **Integration into Query**
```go
var products []models.Product
query := db.Model(&models.Product{})
query = query.Offset((page - 1) * limit).Limit(limit)
if strings.ToLower(order) == "desc" {
    query = query.Order(sortBy + " desc")
} else {
    query = query.Order(sortBy + " asc")
}
if err := query.Find(&products).Error; err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
```
- **Setting Up the Query**: `db.Model(&models.Product{})` initializes the query for the `Product` model.
- **Pagination**: `query.Offset((page - 1) * limit).Limit(limit)` handles pagination.
- **Sorting**: The `Order` method is applied to specify how the results should be sorted.
- **Executing the Query**: `query.Find(&products)` executes the query and retrieves the sorted and paginated results.

### Summary

The `Order` method is responsible for specifying the sorting of the query results. By passing it a string that includes the column name and the sorting direction (`asc` or `desc`), you instruct GORM to generate the appropriate SQL `ORDER BY` clause. This method is directly tied to the sorting feature provided by the underlying database engine, which uses optimized algorithms to sort the data efficiently.


## Explaining the `Offset` function in pagination

The line `query = query.Offset((page - 1) * limit).Limit(limit)` is used to implement pagination in your database queries. Here’s a breakdown of what each part does:

### `Offset`

- **Purpose**: Specifies how many records to skip before starting to return results.
- **Usage**: `Offset(n)` tells the database to skip the first `n` records.
- **Calculation**: In the context of pagination:
  - `page` is the current page number (e.g., page 1, page 2).
  - `limit` is the number of records per page (e.g., 10 records per page).

  To calculate the offset:
  - For page 1, you want to skip 0 records, so the offset is `(1 - 1) * limit = 0`.
  - For page 2, you want to skip the first `limit` records, so the offset is `(2 - 1) * limit = limit`.
  - For page 3, you want to skip the first `2 * limit` records, so the offset is `(3 - 1) * limit = 2 * limit`.

### `Limit`

- **Purpose**: Specifies the maximum number of records to return.
- **Usage**: `Limit(n)` tells the database to return at most `n` records.

### Combined Pagination Calculation

- **Calculation**: The expression `(page - 1) * limit` calculates the total number of records to skip. This is determined by multiplying the number of records per page (`limit`) by the number of pages to skip (`page - 1`).

  Here’s how it works:
  - For page 1 (`page = 1`): `(1 - 1) * limit = 0`. Skip 0 records and return the first `limit` records.
  - For page 2 (`page = 2`): `(2 - 1) * limit = limit`. Skip the first `limit` records and return the next `limit` records.
  - For page 3 (`page = 3`): `(3 - 1) * limit = 2 * limit`. Skip the first `2 * limit` records and return the next `limit` records.

### Example

Assuming `limit = 10`:

- **Page 1**: `Offset((1 - 1) * 10).Limit(10)` returns records 0 to 9.
- **Page 2**: `Offset((2 - 1) * 10).Limit(10)` returns records 10 to 19.
- **Page 3**: `Offset((3 - 1) * 10).Limit(10)` returns records 20 to 29.

This setup allows you to retrieve records in chunks based on the page number and the number of records per page, which is crucial for handling large datasets efficiently.


## Filtering Feature and Real World Explanation

### Explanation of Code

This part of the code is responsible for parsing the minimum and maximum price filters from the query parameters in a HTTP request. Let's break it down:

```go
var minPrice, maxPrice float64
var err error

if minPriceStr != "" {
    minPrice, err = strconv.ParseFloat(minPriceStr, 64)
    if err != nil {
        http.Error(w, "Invalid minimum price", http.StatusBadRequest)
        return
    }
}

if maxPriceStr != "" {
    maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
    if err != nil {
        http.Error(w, "Invalid maximum price", http.StatusBadRequest)
        return
    }
}
```

### Line-by-Line Explanation

1. **Variable Declaration**:
   ```go
   var minPrice, maxPrice float64
   var err error
   ```
   - Declare `minPrice` and `maxPrice` as `float64` to hold the parsed price values.
   - Declare `err` to capture any errors that occur during parsing.

2. **Parsing `minPriceStr`**:
   ```go
   if minPriceStr != "" {
       minPrice, err = strconv.ParseFloat(minPriceStr, 64)
       if err != nil {
           http.Error(w, "Invalid minimum price", http.StatusBadRequest)
           return
       }
   }
   ```
   - Check if `minPriceStr` is not an empty string (i.e., the user provided a minimum price).
   - Attempt to parse `minPriceStr` to a `float64`.
   - If parsing fails (`err` is not `nil`), respond with a "400 Bad Request" status and an error message.

3. **Parsing `maxPriceStr`**:
   ```go
   if maxPriceStr != "" {
       maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
       if err != nil {
           http.Error(w, "Invalid maximum price", http.StatusBadRequest)
           return
       }
   }
   ```
   - Similar to `minPriceStr`, check if `maxPriceStr` is not empty.
   - Attempt to parse `maxPriceStr` to a `float64`.
   - If parsing fails, respond with a "400 Bad Request" status and an error message.

### Application in a Real-World E-commerce Website

In an e-commerce website, customers often want to filter products based on price. This feature enables users to specify a price range and see products within that range. Here’s how this applies in a real-world scenario:

1. **User Interaction**:
   - A user visits the product listing page and sees filter options.
   - The user sets the minimum and maximum price range, e.g., min: $50, max: $200.

2. **Client Request**:
   - The client's browser sends an HTTP GET request with these query parameters:
     ```
     GET /products?min_price=50&max_price=200
     ```

3. **Backend Handling**:
   - The server receives the request and extracts `min_price` and `max_price` from the query parameters.
   - The code parses these values to ensure they are valid numbers.
   - The server uses these parsed values to filter products from the database that fall within the specified price range.

4. **Database Query**:
   - The query might look something like this:
     ```go
     if minPriceStr != "" {
         query = query.Where("price >= ?", minPrice)
     }
     if maxPriceStr != "" {
         query = query.Where("price <= ?", maxPrice)
     }
     ```

5. **Response to Client**:
   - The server retrieves the filtered products from the database.
   - The server sends a JSON response back to the client with the filtered product list.

   **How a query with search, filtering, pagination, and sorting would look like:**

   `GET /products?page=2&limit=5&sort_by=price&order=desc&category=electronics&min_price=100&max_price=1000&search=phone`
   This request retrieves the second page of products (5 per page), sorted by price in descending order, filtered by category "electronics" and price range 100 to 1000, with a search term "phone".


## ON SQL `ILIKE` and Wildcard Character

   `ILIKE` is a SQL keyword used in PostgreSQL to perform a case-insensitive pattern match. It is similar to the `LIKE` operator, but `ILIKE` ignores case when matching text.

### How `ILIKE` Works

When you use `ILIKE` in a SQL query, it matches the pattern regardless of whether the characters are uppercase or lowercase. This is particularly useful when you want to perform a case-insensitive search.

### Example Usage

Suppose you have a table called `products` with a column `name`. You want to search for products whose names contain the word "phone" regardless of whether "phone" is written as "Phone", "PHONE", "pHoNe", etc.

```sql
SELECT * FROM products WHERE name ILIKE '%phone%';
```

### Explanation

- `ILIKE`: The keyword indicating a case-insensitive search.
- `'phone'`: The pattern you are searching for.
- `%`: Wildcard characters that match any sequence of characters (including no characters).

### Real-World Application in E-commerce

In an e-commerce website, users might search for products using different cases. For example, a user might search for "Phone", "phone", "PHONE", etc. By using `ILIKE`, the backend can handle these searches case-insensitively, providing a more user-friendly search experience.

### Integrating `ILIKE` into Go Code

Let's modify the `GetProducts` handler to use `ILIKE` for case-insensitive search:

**Code**:
```go
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Retrieve query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")
	category := r.URL.Query().Get("category")
	minPriceStr := r.URL.Query().Get("min_price")
	maxPriceStr := r.URL.Query().Get("max_price")
	search := r.URL.Query().Get("search")

	// Set default values for pagination
	page := 1
	limit := 10
	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit number", http.StatusBadRequest)
			return
		}
	}

	// Set default sorting
	if sortBy == "" {
		sortBy = "name"
	}
	if order == "" {
		order = "asc"
	}

	// Parse price filters
	var minPrice, maxPrice float64
	var err error
	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			http.Error(w, "Invalid minimum price", http.StatusBadRequest)
			return
		}
	}
	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			http.Error(w, "Invalid maximum price", http.StatusBadRequest)
			return
		}
	}

	// Query the database with pagination, sorting, filtering, and search
	var products []models.Product
	query := db.Model(&models.Product{})
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if minPriceStr != "" {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPriceStr != "" {
		query = query.Where("price <= ?", maxPrice)
	}
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	query = query.Offset((page - 1) * limit).Limit(limit)
	if strings.ToLower(order) == "desc" {
		query = query.Order(sortBy + " desc")
	} else {
		query = query.Order(sortBy + " asc")
	}
	if err := query.Find(&products).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
```

### Explanation of the Search Part

```go
if search != "" {
	query = query.Where("name ILIKE ?", "%"+search+"%")
}
```
- `query.Where("name ILIKE ?", "%"+search+"%")`: This adds a condition to the query to search for products where the `name` matches the `search` term case-insensitively.
- `%` before and after the search term: These are wildcards that allow for partial matches.


## On Extracting Variables from URL using Mux

The line `vars := mux.Vars(r)` is used to extract variables from the URL path in a request when using the Gorilla Mux router in Go.

### How It Works

When you define routes in Gorilla Mux, you can include variables in the URL path. For example, if you have a route defined as `/cart/{id}`, `{id}` is a variable placeholder. When a request matches this route, you can extract the value of `id` using `mux.Vars(r)`.


### Line-by-Line Breakdown

1. **Extract Variables**:
   ```go
   vars := mux.Vars(r)
   ```
   This line extracts the variables from the URL path and stores them in a map. For example, if the URL is `/cart/123`, `vars` will contain `{"id": "123"}`.

2. **Convert ID to Integer**:
   ```go
   id, err := strconv.Atoi(vars["id"])
   ```
   Retrieve the `id` value from the `vars` map and convert it from a string to an integer. If the conversion fails, return a `400 Bad Request` error.



## PaymnetHandler Concurrency Logic

### Concurrency Explanation

The provided `PaymentHandler` function uses concurrency to handle payment processing efficiently. Here's how each part of the concurrency model works:

#### 1. **Context with Timeout**
```go
ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
defer cancel()
```
- **Purpose**: Creates a context with a 10-second timeout to ensure that the payment processing doesn't hang indefinitely.
- **Real-World Simulation**: If a payment request takes too long (e.g., due to network issues or delays from the payment provider), this context will cancel the operation, preventing the server from waiting forever.

#### 2. **WaitGroup**
```go
var wg sync.WaitGroup
wg.Add(1)
```
- **Purpose**: Manages synchronization by waiting for the payment processing goroutine to complete.
- **Real-World Simulation**: Ensures that the main function waits for all concurrent tasks to finish before closing channels.

#### 3. **Channels for Communication**
```go
resultChan := make(chan *stripe.Charge)
errorChan := make(chan error)
```
- **Purpose**: Channels are used to send and receive the results of the payment processing.
- **Real-World Simulation**: `resultChan` will receive the successful charge object, while `errorChan` will receive any errors that occur during payment processing.

#### 4. **Goroutines for Concurrency**
```go
go func() {
    defer wg.Done()
    // Payment processing logic
}()
```
- **Purpose**: Executes the payment processing in a separate goroutine to avoid blocking the main request-handling flow.
- **Real-World Simulation**: Allows the server to handle multiple payment requests simultaneously. Each request is processed independently and does not block other requests.

#### 5. **Select Statement**
```go
select {
case ch := <-resultChan:
    // Handle successful charge
case err := <-errorChan:
    // Handle payment processing error
case <-ctx.Done():
    // Handle request timeout
}
```
- **Purpose**: Handles different outcomes (successful payment, error, or timeout) based on the channels and context.
- **Real-World Simulation**: Ensures that the handler can react appropriately to each outcome:
  - **Success**: Updates the payment status and returns the charge details.
  - **Error**: Returns an error message to the client if the payment fails.
  - **Timeout**: Returns a timeout error if the request takes too long.

### Real-World Application

#### Scenario
In a real-world e-commerce website, users make payments for their orders. When a payment request is received, the server needs to:
1. **Process the Payment**: Handle the payment transaction with Stripe.
2. **Update Status**: Reflect the payment status in the system.
3. **Handle Multiple Requests**: Ensure the system can handle multiple payment requests concurrently without performance degradation.

#### Integration

1. **Concurrent Payment Processing**: With the use of goroutines, the server can process multiple payments at the same time, which is crucial during peak shopping periods or sales events.

2. **Timeout Management**: The timeout ensures that the server doesn't hang indefinitely on a payment request. This is essential to maintain responsiveness, especially during network or server issues.

3. **Error Handling**: Using channels and context helps in handling errors effectively, ensuring users are informed promptly if something goes wrong.

4. **Scalability**: This approach allows the server to scale by handling more payment requests concurrently, improving overall efficiency.

#### Visual Representation

Here’s a simplified visual representation of how concurrency fits into the payment handler:

```plaintext
   +------------------+
   | Payment Request  |
   +------------------+
           |
           v
   +----------------------+
   | Goroutine (Payment)  | <------------------------+
   +----------------------+                          |
           |                                       +---+---+
           v                                       | Channel|
   +----------------------+                          +-------+
   | Stripe Payment API   |                              |
   +----------------------+                              |
           |                                       +------+------+
           v                                       | Result/Error |
   +----------------------+                          +--------------+
   | Update Payment Status|                              |
   +----------------------+                              |
           |                                       +------v------+
           v                                       | Return Result|
   +----------------------+                          +--------------+
   | Send Response to     |
   | Client               |
   +----------------------+
```

In this flow:
1. **Payment Request**: Received from the user.
2. **Goroutine**: Processes the payment in the background.
3. **Stripe Payment API**: Handles the actual transaction.
4. **Channels**: Communicate results and errors.
5. **Update Payment Status**: Updates the system with payment results.
6. **Send Response**: Returns the result to the user.

This approach ensures that your payment handler is both efficient and robust, handling multiple requests and errors gracefully.


## On Idempotency

The idempotency key in the Stripe API is used to ensure that a particular operation (like creating a charge) is only executed once, even if it is sent multiple times. This is crucial for preventing duplicate charges, which can happen due to network issues, retries, or user actions.

### Placement of Idempotency Key

The idempotency key should **precede** the charge creation to ensure that Stripe uses it to prevent duplicate transactions:

1. **Define the Idempotency Key**:
   Set the `IdempotencyKey` field before making the charge request.

2. **Create the Charge**:
   Use the `chargeParams` with the `IdempotencyKey` when creating the charge.

Here's how you should refactor it:

```go
go func() {
    defer wg.Done()
    // Create charge parameters
    chargeParams := &stripe.ChargeParams{
        Amount:      stripe.Int64(amountInCents),
        Currency:    stripe.String("usd"),
        Description: stripe.String("Charge for order " + strconv.Itoa(paymentRequest.OrderID)),
    }

    // Add email for receipt
    chargeParams.ReceiptEmail = stripe.String(paymentRequest.Email)

    // Add metadata to charge parameters
    chargeParams.AddMetadata("order_id", strconv.Itoa(paymentRequest.OrderID))
    chargeParams.AddMetadata("transaction_id", paymentRequest.TransactionID)
    chargeParams.AddMetadata("payment_method", paymentRequest.PaymentMethod)

    // Add shipping details to metadata (for physical goods)
    chargeParams.AddMetadata("shipping_carrier", shipping.Carrier)
    chargeParams.AddMetadata("tracking_number", shipping.TrackingNumber)
    chargeParams.AddMetadata("shipping_method", shipping.ShippingMethod)
    chargeParams.AddMetadata("shipping_cost", strconv.FormatFloat(shipping.ShippingCost, 'f', 2, 64))
    chargeParams.AddMetadata("estimated_delivery", shipping.EstimatedDelivery.String())
    chargeParams.AddMetadata("shipping_date", shipping.ShippingDate)
    chargeParams.AddMetadata("shipping_type", shipping.ShippingType)
    chargeParams.AddMetadata("shipping_address", shipping.ShippingAddress)
    chargeParams.AddMetadata("shipping_city", shipping.ShippingCity)
    chargeParams.AddMetadata("shipping_state", shipping.ShippingState)
    chargeParams.AddMetadata("shipping_zip_code", shipping.ShippingZipCode)
    chargeParams.AddMetadata("shipping_country", shipping.ShippingCountry)

    // Prevents or handles duplicate charges gracefully
    chargeParams.IdempotencyKey = stripe.String(paymentRequest.TransactionID)

    // Create the charge
    ch, err := charge.New(chargeParams)
    if err != nil {
        errorChan <- err
        return
    }

    resultChan <- ch
}()
```

### Real-World Simulation

1. **Customer Payment Submission**: A user submits a payment for an order on your e-commerce site.
2. **Server Handles Payment**: The server processes the payment in a goroutine.
3. **Idempotency Key**: The `TransactionID` is used to ensure that if the same payment request is submitted multiple times (e.g., due to network retries or user mistakes), only one charge is processed.
4. **Charge Creation**: Stripe uses the idempotency key to prevent duplicate charges.
5. **Response Handling**: Based on the charge result or error, the server updates the payment status and responds to the client.

By placing the `IdempotencyKey` correctly, you ensure that your payment system is robust against duplicate charges, enhancing reliability and user experience.



## Updated code with Retry logic

```go
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var paymentRequest models.Payment
	var shipping models.Shipping

	// Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Initialize Stripe with secret key
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Convert amount to cents
	amountInCents := int64(paymentRequest.Amount * 100)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Channel to receive the charge result
	resultChan := make(chan *stripe.Charge)
	errorChan := make(chan error)

	go func() {
		defer wg.Done()
		// Create charge parameters
		chargeParams := &stripe.ChargeParams{
			Amount:      stripe.Int64(amountInCents),
			Currency:    stripe.String("usd"),
			Description: stripe.String("Charge for order " + strconv.Itoa(paymentRequest.OrderID)),
		}

		// Add email for receipt
		chargeParams.ReceiptEmail = stripe.String(paymentRequest.Email)

		// Add metadata to charge parameters
		chargeParams.AddMetadata("order_id", strconv.Itoa(paymentRequest.OrderID))
		chargeParams.AddMetadata("transaction_id", paymentRequest.TransactionID)
		chargeParams.AddMetadata("payment_method", paymentRequest.PaymentMethod)

		// Add shipping details to metadata (for physical goods)
		chargeParams.AddMetadata("shipping_carrier", shipping.Carrier)
		chargeParams.AddMetadata("tracking_number", shipping.TrackingNumber)
		chargeParams.AddMetadata("shipping_method", shipping.ShippingMethod)
		chargeParams.AddMetadata("shipping_cost", strconv.FormatFloat(shipping.ShippingCost, 'f', 2, 64))
		chargeParams.AddMetadata("estimated_delivery", shipping.EstimatedDelivery.String())
		chargeParams.AddMetadata("shipping_date", shipping.ShippingDate)
		chargeParams.AddMetadata("shipping_type", shipping.ShippingType)
		chargeParams.AddMetadata("shipping_address", shipping.ShippingAddress)
		chargeParams.AddMetadata("shipping_city", shipping.ShippingCity)
		chargeParams.AddMetadata("shipping_state", shipping.ShippingState)
		chargeParams.AddMetadata("shipping_zip_code", shipping.ShippingZipCode)
		chargeParams.AddMetadata("shipping_country", shipping.ShippingCountry)

		// Prevents or handles duplicate charges gracefully
		chargeParams.IdempotencyKey = stripe.String(paymentRequest.TransactionID)

	 // Retry logic in case of failure
        maxRetries := 3
        var ch *stripe.Charge 
        for i := 0; i < maxRetries; i++ {
            ch, err = charge.New(chargeParams)
            if err != nil {
                if stripeErr, ok := err.(*stripe.Error); ok {
                    // Retry on certain transient errors
                    if stripeErr.Type == stripe.ErrorTypeAPIConnection || stripeErr.Code == "lock_timeout" {
                        log.Printf("Transient error: %v. Retrying...", err)
                        time.Sleep(2 * time.Second)
                        continue
                    }
                }
                // Send the error to the error channel if not retryable
                errorChan <- err
                return
            }
            // If charge is successful, send it to result channel
            resultChan <- ch
            return
        }
        // If all retries are exhausted, send the last error to the error channel
        errorChan <- err
    }()

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case ch := <-resultChan:
		// Update payment status and transaction ID
		paymentRequest.Status = ch.Status
		paymentRequest.TransactionID = ch.ID

		// Respond with the charge details
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ch)

	case err := <-errorChan:
		log.Printf("Stripe charge creation failed: %v", err)
		http.Error(w, "Payment processing failed", http.StatusInternalServerError)

	case <-ctx.Done():
		log.Printf("Request timed out")
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	}

}
```


## Stripe's Built-in charge.New Function

### Function Definition
- `charge.New(chargeParams *stripe.ChargeParams) (*stripe.Charge, error)`

### Parameters
- **`chargeParams`**: This is a pointer to a `stripe.ChargeParams` struct. This struct contains all the necessary parameters required to create a charge, such as the amount, currency, description, receipt email, and metadata.

### Returns
- **`*stripe.Charge`**: If the charge is successfully created, this function returns a pointer to a `stripe.Charge` struct. This struct contains all the details of the created charge, such as the charge ID, status, amount, currency, and metadata.
- **`error`**: If there is an error while creating the charge, this function returns an error object detailing what went wrong.

### Example Usage
Here is a simplified example of how `charge.New` is used:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/charge"
)

func main() {
	// Initialize Stripe with the secret key
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Create charge parameters
	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000), // Amount in cents (e.g., $20.00)
		Currency:    stripe.String("usd"),
		Description: stripe.String("Charge for order 123"),
		ReceiptEmail: stripe.String("customer@example.com"),
		Metadata: map[string]string{
			"order_id":       "123",
			"payment_method": "card",
		},
	}

	// Create the charge
	ch, err := charge.New(chargeParams)
	if err != nil {
		log.Fatalf("Stripe charge creation failed: %v", err)
	}

	// Print the charge details
	fmt.Printf("Charge created: %+v\n", ch)
}
```

### Real-World Application
In a real-world e-commerce application, you would use `charge.New` to process payments from customers. Here's how it fits into a typical payment flow:

1. **Customer Places Order**: The customer places an order on the website and proceeds to the checkout page.
2. **Payment Details Submission**: The customer submits their payment details (e.g., credit card information).
3. **Backend Payment Processing**: The backend server receives the payment details and prepares the charge parameters.
4. **Charge Creation**: The server calls `charge.New` with the charge parameters to create the charge on Stripe.
5. **Error Handling**: If there is an error (e.g., insufficient funds, invalid card), the server handles the error and notifies the customer.
6. **Success Handling**: If the charge is successful, the server updates the order status, sends a receipt to the customer, and may trigger additional processes like inventory updates or shipping.

This function is crucial for integrating Stripe's payment processing capabilities into your Go-based application, ensuring secure and reliable transactions for your customers.


## On Implementing Stripe Subscription Pcakage

```go
package subscription

import (
    "net/http"
    "github.com/stripe/stripe-go/v72"
)

// getC is a helper function that returns an instance of Client.
// Assume getC is defined elsewhere in your package.
func getC() Client {
    return Client{
        Key: stripe.Key, // Use the Stripe API key here.
        B:   stripe.GetBackend(stripe.APIBackend), // Use the Stripe backend here.
    }
}

// Client struct representing a Stripe client.
// Assume this struct is defined elsewhere in your package.
type Client struct {
    Key string
    B   stripe.Backend
}

// Cancel removes a subscription.
func Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
    return getC().Cancel(id, params)
}

// Cancel removes a subscription.
func (c Client) Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
    path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
    sub := &stripe.Subscription{}
    err := c.B.Call(http.MethodDelete, path, c.Key, params, sub)
    return sub, err
}
```

### Explanation:

1. **Client struct**: Represents the Stripe client, which has the method `Cancel`.
    ```go
    type Client struct {
        Key string
        B   stripe.Backend
    }
    ```

2. **getC function**: Returns an instance of the `Client`. This function provides a way to get a `Client` instance with the necessary setup (API key, backend, etc.).
    ```go
    func getC() Client {
        return Client{
            Key: stripe.Key, // Use the Stripe API key here.
            B:   stripe.GetBackend(stripe.APIBackend), // Use the Stripe backend here.
        }
    }
    ```

3. **Package-level Cancel function**: This function calls the `Cancel` method on the `Client` instance returned by `getC`.
    ```go
    func Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
        return getC().Cancel(id, params)
    }
    ```

4. **Method on Client (Cancel)**: This is the actual implementation that sends the HTTP request to cancel the subscription.
    ```go
    func (c Client) Cancel(id string, params *stripe.SubscriptionCancelParams) (*stripe.Subscription, error) {
        path := stripe.FormatURLPath("/v1/subscriptions/%s", id)
        sub := &stripe.Subscription{}
        err := c.B.Call(http.MethodDelete, path, c.Key, params, sub)
        return sub, err
    }
    ```

## On Trial Periods and Cycle Billings

If a subscription is created beyond the trial period, the behavior depends on the `TrialPeriodDays`, `BillingCycleAnchor`, and `ProrationBehavior` settings you configure in your backend logic. Here’s how each aspect affects the subscription lifecycle:

### 1. **Trial Period**

- **Definition:** The `TrialPeriodDays` is the number of days a user gets free access to the subscription before they are billed.
- **Beyond Trial Period:** If the subscription is started beyond the trial period or if the user tries to activate a subscription after the trial period has expired, they will not receive any free trial days. The user will be billed immediately or on the next billing cycle, depending on the `BillingCycleAnchor` configuration.

### 2. **Billing Cycle Anchor**

- **Definition:** The `BillingCycleAnchor` sets the start date for the billing cycle.
- **Beyond Trial Period:** If the `BillingCycleAnchor` is set to a date in the future or beyond the trial period, the subscription will start billing from that date. For example, if the anchor is set to the 1st of the next month, the user will be billed on that date regardless of when the subscription was created or if a trial period is included.

## Libraries Used:

- [Go library for the Stripe API.](https://github.com/stripe/stripe-go)
- Firebase for push notifications


## Discoveries along the way

- [Dependency Injection](https://www.jetbrains.com/guide/go/tutorials/dependency_injection_part_one/introduction/). In the `CheckoutHandler` Function (check `checkout.go`), http.HandlerFunc depends on gorm.DB, which is injected into the function. This follows the DI principle where dependencies are provided to an object what it is initialialized.

- I discovered that frontend isn't all about design as I used to imagine (at least in my HTML and CSS constrained mind), there are some logic to it too. Though while the frontend and backend have logic, the nature of that logic is different: the frontend focuses on interaction and presentation (like form validation, responsiveness, data fetching-requesting data from an API and updating the UI without reloading the page, in the case of redirecting a user to a confirmation page upon successful payment.), and the backend focuses on processing and security (like authentication, database interaction, calculations, and security enforcement).

- The definition and application of SDK:
        SDK stands for **Software Development Kit**. It's a collection of tools, libraries, documentation, and code samples that developers use to create applications for a specific platform or framework.

        ### Key Components of an SDK:
        - **Tools**: These include compilers, debuggers, and other utilities that help in the development process.
        - **Libraries/Frameworks**: Pre-written code that developers can use to perform common tasks, such as handling network requests, interacting with hardware, or managing user interfaces.
        - **Documentation**: Guides, manuals, and API references that explain how to use the SDK effectively.
        - **Code Samples**: Example code that demonstrates how to implement certain features or solve specific problems using the SDK.

        ### Purpose:
        An SDK simplifies the development process by providing everything you need to create applications for a particular platform, such as iOS, Android, or web development frameworks like Firebase or React.

    Really can't be called a discovery, but well, I ran into an issue with GitHub's **secret-catching police** (*no jokes*), it is called **"GitHub's push protection feature"** (*doesn't even sound like a real name*). Basically, what it does is that it scans for secrets (such as API keys, service account credentials, etc.) in commits before allowing them to be pushed to the repository. Apparently, I happened to have a Google Cloud Service Account Credentials in one *utils/serviceAccountKey.json* file (a file I created for Firebase credentials) registered at commit `72f4a6168ecc8f008c818a0378a96da99582c9db`, which is what got me into trouble.

    To bail myself, I had to search for a solution. Thanks to judge co-pilot who came to my rescue. Long story short, I was given a life-or-death card: to choose either to allow secrets to be pushed to my repository by following a URL sent to my terminal, or to remove the secret from my commit history. Well, as a wise kid that I am, I chose life!

    So I can either use `git filter-branch` or `git rebase` to undo the commit history. Here is the full command I used:

    `git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch utils/serviceAccountKey.json' --prune-empty --tag-name-filter cat -- --all`, to remove the file from the commit history.

    Followed by:

    `git push -f` (*yes, a forced push!*)

    But then after my bail, judge co-pilot gave me a close warnng, saying: 

    *Hey boy, you gotta be careful when using `--force` or `-f` as it rewrites commit history and can affect dem collaborators!* (well, I should look into that in the future) :)


**Here is what the code that handles the frontend redirection for payment looks like:**

```javascript
// Example of initiating a payment and receiving confirmation
function processPayment(orderId, paymentDetails) {
    // Step 1: Send payment details to the gateway
    paymentGateway.processPayment(paymentDetails)
        .then(paymentConfirmation => {
            // Step 2: On success, send confirmation to backend
            return fetch(`/order/confirm/${orderId}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    paymentId: paymentConfirmation.id,
                    status: paymentConfirmation.status
                })
            });
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === 'Completed') {
                // Step 3: Redirect to confirmation page
                window.location.href = `/order/success/${orderId}`;
            }
        })
        .catch(error => {
            console.error('Payment or confirmation failed:', error);
            alert('There was an error processing your payment. Please try again.');
        });
}
```

**Here is how a confirmation page looks like (HTML/CSS):**

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Order Confirmation</title>
    <style>
        .confirmation-page {
            text-align: center;
            padding: 50px;
        }
        .order-details {
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="confirmation-page">
        <h1>Thank You for Your Order!</h1>
        <p>Your order has been successfully processed.</p>
        <div class="order-details">
            <h3>Order Summary</h3>
            <p>Order ID: #123456</p>
            <p>Status: Completed</p>
            <p>Total: $89.99</p>
            <!-- Additional order details can be listed here -->
        </div>
        <a href="/order/history">View Order History</a>
    </div>
</body>
</html>
```



# APPLICATION LOGIC

## Design flow from Cart to Order


#### 1. **Add Items to Cart**
- Endpoint: `POST /cart`
- Input: Product ID, Quantity

- **Purpose**: Allows users to add products to their cart, which is a temporary storage for items they intend to purchase.

- **Workflow**:

  1. **User Interaction**: The user selects a product and specifies the quantity.

  2. **Backend Processing**:
     - Validate the product ID and ensure it exists.
     - Check if the specified quantity is available.
     - Add the product to the cart, or update the quantity if it already exists in the cart.

  3. **Data Storage**: The cart information is stored in the database, associated with the user's session or account.

#### 2. **View Cart**

- Endpoint: `GET /cart`
- Input: User ID

- **Purpose**: Allows users to review the items in their cart before proceeding to checkout.

- **Workflow**:

  1. **User Interaction**: The user requests to view their cart.

  2. **Backend Processing**:
     - Retrieve the list of items in the user's cart from the database.
     - Calculate the subtotal for each item and the total for the cart.

  3. **Data Presentation**: The cart details (product names, quantities, prices, and total cost) are returned to the user.

#### 3. **Checkout**

- Endpoint: `POST /checkout`
- Input: User ID

- **Purpose**: Transforms the cart into a formal order, preparing it for payment.

- **Workflow**:

  1. **User Interaction**: The user initiates the checkout process.
  2. **Backend Processing**:
     - Retrieve the cart items from the database.
     - Validate the availability of items and recalculate the total cost.
     - Create a new order record in the database with status "Pending".
     - Transfer the items from the cart to the order and clear the cart.

  3. **Data Storage**: The order is saved in the database, and the cart is emptied.

#### 4. **Payment Processing**

- Endpoint: `POST /payment`
- Input: Order ID, payment token

- **Purpose**: Handles the payment for the order using a payment gateway.

- **Workflow**:

  1. **User Interaction**: The user provides payment details (e.g., credit card information).

  2. **Backend Processing**:
     - Integrate with a payment gateway (e.g., Stripe) to process the payment.
     - Handle payment responses:
       - If successful, update the order status to "Paid" or "Completed".
       - If unsuccessful, notify the user of the failure and allow them to retry.

  3. **Data Storage**: The payment status and transaction details are stored in the database.

#### 5. **Order Confirmation**

- Endpoint: `POST /order/confirm/{orderID}`
- Input: Order ID

- **Purpose**: Finalizes the order after successful payment and sends confirmation to the user.

- **Workflow**:
  1. **User Interaction**: The user is redirected to an order confirmation page.

  2. **Backend Processing**:
     - Retrieve the order by ID.
     - Update the order status to "Completed".
     - Optionally, trigger additional actions like sending a confirmation email to the user (using Mailgun or Kafka)

  3. **Data Presentation**: Display order confirmation details to the user and provide a receipt.


The process of the frontend receiving confirmation after a successful payment and updating the order status involves several steps. Here’s how it works:

### 1. **Payment Process Flow**

#### Step 1: Initiate Payment

- **Frontend Action**: 
  - The user enters payment details on the payment page, typically hosted by or embedded from a payment gateway like Stripe.
  - The frontend submits these details to the payment gateway via a secure API call.

#### Step 2: Payment Gateway Response
- **Payment Gateway Action**:
  - The payment gateway processes the payment.
  - If successful, it returns a confirmation response to the frontend. This response typically includes a payment confirmation ID or status.

#### Step 3: Update Backend (Order Confirmation)
- **Frontend Action**:
  - Upon receiving a successful payment response, the frontend sends a `POST` request to the backend's `OrderConfirmationHandler`.
  - This request typically includes the order ID and payment confirmation details.

#### Step 4: Backend Order Status Update
- **Backend Action**:
  - The backend verifies the payment confirmation (optionally).
  - The backend updates the order status to "Completed" in the database.
  - The backend might also trigger additional actions, like sending a confirmation email to the user.

### 2. **Frontend Receipt of Confirmation and Redirection**

#### Step 5: Frontend Receives Order Status Update
- **Frontend Action**:
  - The backend responds to the frontend's confirmation request with a success message and any relevant order details (e.g., order ID, status).
  - The frontend processes this response.

#### Step 6: Redirect to Confirmation Page
- **Frontend Action**:
  - The frontend redirects the user to an order confirmation page.
  - This page displays the order summary, payment status, and a "Thank You" message.
  - The user might also see options for viewing their order history or downloading an invoice.


## On setting up Firebase  and Firebase Device Token

#### 1. **Setting Up Firebase Project**

   - **Create a Firebase Project:**
     1. Go to the [Firebase Console](https://console.firebase.google.com/).
     2. Click on "Add Project" and follow the setup steps.
   - **Add Firebase to Your Go Project:**
     1. Go to "Project Settings" in Firebase Console.
     2. Under "Your apps," select "Add App" and choose a Web app.
     3. Firebase will generate a `firebaseConfig` object that contains your API keys and project ID.

#### 2. **Install Firebase Admin SDK**

   Add the Firebase Admin SDK:
   ```bash
   go get firebase.google.com/go/v4
   go get google.golang.org/api/option
   ```
#### 3. **Write Logic to Initialize Firebase**

#### 4. **Write Push Notification Logic**


  ### Firebase Device Token

A Firebase device token (also known as an FCM token or registration token) is a long string of characters that uniquely identifies a user's device or web browser instance for Firebase Cloud Messaging (FCM). The token looks something like this:

```
djiGyrMzJRfgbFBO-U1fYo:APA91bHhZyGnQXYmEi6vQjUsIsXngUrTYE1r6jR2-8tf9W4JS1I7a3qEk3hVtkzOWXJ5dQsDCh4XFN9LZCeYAdgROQ2TdpbbzW88eHVz9twx1DpXzXXwctfSaODV6u6KdKngkErEqSMQ
```

This token is generated by the Firebase SDK and is unique for each device/browser and app combination.

### How Will the Web App Get the Device Token?

To obtain the device token in a web application, you need to integrate Firebase Cloud Messaging (FCM) into your web app. Here's a step-by-step guide:

#### 1. **Add Firebase to Your Web App**

First, ensure that you've added Firebase to your web application. This involves including the Firebase SDK and initializing your Firebase app with the configuration you get from the Firebase Console.

```html
<script src="https://www.gstatic.com/firebasejs/9.6.1/firebase-app.js"></script>
<script src="https://www.gstatic.com/firebasejs/9.6.1/firebase-messaging.js"></script>

<script>
  // Firebase configuration object
  const firebaseConfig = {
    apiKey: "your-api-key",
    authDomain: "your-auth-domain",
    projectId: "your-project-id",
    storageBucket: "your-storage-bucket",
    messagingSenderId: "your-messaging-sender-id",
    appId: "your-app-id",
  };

  // Initialize Firebase
  firebase.initializeApp(firebaseConfig);

  // Initialize Firebase Messaging
  const messaging = firebase.messaging();
</script>
```

#### 2. **Request Permission to Receive Notifications**

Before obtaining the device token, you need to request permission from the user to send notifications.

```javascript
messaging.requestPermission().then(() => {
    console.log('Notification permission granted.');
}).catch((err) => {
    console.log('Unable to get permission to notify.', err);
});
```

#### 3. **Get the Device Token**

After the user grants permission, you can retrieve the device token using the `getToken` method.

```javascript
messaging.getToken({ vapidKey: 'your-public-vapid-key' }).then((currentToken) => {
    if (currentToken) {
        console.log('Token obtained:', currentToken);
        // Send the token to your backend to store it
        sendTokenToServer(currentToken);
    } else {
        console.log('No registration token available. Request permission to generate one.');
    }
}).catch((err) => {
    console.log('An error occurred while retrieving token. ', err);
});
```

- **vapidKey**: This is the VAPID public key, which you can generate in the Firebase Console under the Cloud Messaging settings.

#### 4. **Send the Device Token to the Backend**

Once you have the token, you'll typically send it to your backend so that you can store it for later use (e.g., sending push notifications).

```javascript
function sendTokenToServer(token) {
    // Example POST request to send the token to your backend
    fetch('/api/store-device-token', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ token: token })
    }).then(response => {
        if (response.ok) {
            console.log('Token sent to server successfully.');
        } else {
            console.log('Failed to send token to server.');
        }
    }).catch(error => {
        console.log('Error sending token to server:', error);
    });
}
```

- **/api/store-device-token**: This is the endpoint on your backend where you'll handle storing the token. 

### Summary

- **Device Token**: A unique string generated by Firebase that identifies a user's device or browser for push notifications.
- **Getting the Token in a Web App**:
  1. Initialize Firebase in your web app.
  2. Request permission to send notifications.
  3. Retrieve the token using `messaging.getToken`.
  4. Send the token to your backend server for storage.

With this setup, your web app can register for notifications and provide the necessary token to your backend, enabling you to send targeted push notifications to users' devices.

*The JavaScript code is implemented in the frontend of the web application. This code runs in the user's browser, allowing the app to request permission for notifications, retrieve the Firebase device token, and then send that token to the backend server for storage.*

Here's a quick overview:

- **Frontend (JavaScript)**:
  - **Initialize Firebase**: Set up Firebase in the web app with the required Firebase project configuration.
  - **Request Permission**: Ask the user for permission to send push notifications.
  - **Get Device Token**: Retrieve the device token using Firebase Cloud Messaging (FCM).
   - **Send Token to Backend**: Send the retrieved token to the backend API for storage.

This token is then used by your backend to send push notifications specifically to the user's browser instance.

## On Firebase API Key

The API key in a Firebase setup is typically applied in the frontend of your web application when you initialize Firebase. This API key is part of the Firebase configuration object, which contains all the necessary information for your web app to communicate with Firebase services.

Here's where and how you apply the API key:

### Applying the API Key in the Firebase Configuration

1. **Find the Firebase Configuration Object**:
   When you set up a Firebase project, Firebase provides a configuration object that includes your API key, along with other important information like the `authDomain`, `projectId`, etc.

2. **Initialize Firebase with the Configuration Object**:
   In your web app's frontend JavaScript, you'll initialize Firebase using this configuration object. Here's an example:

   ```javascript
   // Your web app's Firebase configuration
   const firebaseConfig = {
       apiKey: "your-api-key",                // Your API Key
       authDomain: "your-auth-domain",
       projectId: "your-project-id",
       storageBucket: "your-storage-bucket",
       messagingSenderId: "your-messaging-sender-id",
       appId: "your-app-id",
   };

   // Initialize Firebase
   firebase.initializeApp(firebaseConfig);
   ```

   Replace `"your-api-key"` with the actual API key provided by Firebase.

### Where to Place This Code

- **In the HTML File**: You can place this JavaScript code within a `<script>` tag in your HTML file, typically at the bottom before the closing `</body>` tag.
- **In a Separate JavaScript File**: Alternatively, you can place this code in a separate JavaScript file (e.g., `app.js`) and link to it in your HTML file using a `<script src="app.js"></script>` tag.

### Example in Context

Here’s an example of how it all comes together:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My Web App</title>
    <script src="https://www.gstatic.com/firebasejs/9.6.1/firebase-app.js"></script>
    <script src="https://www.gstatic.com/firebasejs/9.6.1/firebase-messaging.js"></script>
</head>
<body>
    <!-- Your HTML content here -->

    <script>
        // Your web app's Firebase configuration
        const firebaseConfig = {
            apiKey: "your-api-key",                // Your API Key
            authDomain: "your-auth-domain",
            projectId: "your-project-id",
            storageBucket: "your-storage-bucket",
            messagingSenderId: "your-messaging-sender-id",
            appId: "your-app-id",
        };

        // Initialize Firebase
        firebase.initializeApp(firebaseConfig);
    </script>
</body>
</html>
```

This code will set up Firebase in your web application, using the API key and other configuration details to connect to your Firebase project. Once initialized, you can use Firebase services like Cloud Messaging, Authentication, Firestore, and more in your web app.


## On the use of Redis for cache

Redis is an in-memory database that runs on your server or a cloud instance. When you store data in Redis, it is kept in the RAM (Random Access Memory) of the server where Redis is running. This allows for extremely fast read and write operations because accessing data from RAM is much quicker than accessing it from *disk storage.

In an e-commerce backend, Redis will serve as a high-speed cache layer. When users browse your product listings, instead of hitting the database each time, the application will first check Redis. If the data is cached, it will return the cached data almost instantly, reducing *latency and server load. This setup is crucial for handling high traffic and ensuring a smooth user experience.

### How Redis Works:

1. **Centralized Cache**: 
   - Redis stores data in the server's RAM, meaning that the data is accessible to all users interacting with the application. 
   - For example, if multiple users request the same product information on an e-commerce platform, Redis can serve this data from its cache rather than querying the database each time. This reduces the load on the database and speeds up response times for all users.

2. **Efficiency and Scalability**:
   - By caching commonly requested data, Redis reduces the need for repeated database queries, which can be slow and resource-intensive. This is particularly important in a high-traffic environment where many users might be requesting the same or similar data.
   - Redis can also handle thousands or even millions of requests per second, making it ideal for applications that need to serve a large user base efficiently.

3. **Use Cases**:
   - **Product Catalog**: In an e-commerce backend, Redis could store a frequently accessed product catalog. When any user searches for or views products, the information is quickly served from Redis rather than querying the database each time.
   - **Session Management**: Redis can store session data, allowing multiple users to maintain their sessions across different requests. This is common in scenarios where users are frequently logging in and out or interacting with personalized content.
   - **Rate Limiting**: Redis can be used to track and enforce rate limits across multiple users, ensuring that no single user overwhelms the system.

### Redis Terminologies

- **Cache Key**: This create a unique cache key based on the parameters of the request. This ensures that different queries are cached separately.
- **Redis GET**: Before querying the database, we attempt to retrieve the data from Redis using the generated cache key.
- **Cache Miss**: If the data is not in Redis (err == redis.Nil), we proceed to query the database, then cache the result in Redis for future requests.
- **Cache Hit**: If the data is found in Redis, we return it immediately without querying the database, which reduces load and improves response times.


## On Instantiating Redis DB (plus using Goroutine to handle context)


### Function Definition
```go
func InitRedisClient() *redis.Client {
```
- **Purpose**: This function initializes and returns a Redis client that can be used to interact with a Redis database. 
- **Return Type**: The function returns a pointer to a `redis.Client` object, which represents the Redis client.

### Create Redis Client
```go
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:3003", // Redis server address
    Password: "",               // No password set
    DB:       0,                // Use default DB
})
```
- **`redis.NewClient`**: This function creates a new Redis client using the provided options.
- **Options**:
  - `Addr`: Specifies the address of the Redis server. In this case, it's `localhost:3003`, meaning it's running on the local machine on port `3003`.
  - `Password`: The password for the Redis server is set to an empty string, meaning no password is required.
  - `DB`: Specifies which Redis database to use. The default is `0`.

### Create a Context with Timeout
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```
- **`context.WithTimeout`**: Creates a context that automatically cancels after 10 seconds.
  - **`ctx`**: This is the context that controls the timeout for the Redis connection attempt.
  - **`cancel`**: This is a function that can be called to cancel the context manually, but here it's deferred.
- **`defer cancel()`**: Ensures that the `cancel` function is called once the function exits, which frees up resources.

### Create a Result Channel
```go
resultChan := make(chan error, 1)
```
- **`make(chan error, 1)`**: Creates a channel of type `error` with a buffer size of 1.
- **Purpose**: This channel is used to communicate the result of the Redis `Ping` operation (either an error or success) back to the main function.

### Start a Goroutine to Ping Redis
```go
go func() {
    _, err := rdb.Ping().Result()
    if err != nil {
        resultChan <- err
    } else {
        close(resultChan)
    }
}()
```
- **Goroutine**: A separate goroutine is started to ping the Redis server.
  - **`rdb.Ping().Result()`**: Pings the Redis server to check if it is reachable.
  - **Error Handling**:
    - If the ping fails (`err != nil`), the error is sent to the `resultChan`.
    - If the ping succeeds, the channel is closed to signal success.
- **Purpose**: This goroutine allows the Redis ping operation to run concurrently with other code and to prevent blocking the main execution thread.

### Select Statement to Handle Timeout or Result
```go
select {
case <-ctx.Done():
    log.Fatalf("Context timeout: %v", ctx.Err())
    return nil
case err := <-resultChan:
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
        return nil
    }
    return rdb
}
```
- **`select` Statement**: Waits for either the context to time out or the result of the Redis ping.
  - **`case <-ctx.Done()`**:
    - Triggered if the context's timeout is reached.
    - Logs a fatal error indicating that the operation timed out and returns `nil`.
  - **`case err := <-resultChan:`**:
    - Triggered if the Redis ping operation completes (either successfully or with an error).
    - If an error occurred, it logs a fatal error and returns `nil`.
    - If no error occurred, it returns the initialized Redis client (`rdb`).
- **Purpose**: This mechanism ensures that the function either successfully initializes a Redis client or terminates gracefully if the operation takes too long or fails.

### Summary
- **Purpose of the Code**: The function attempts to connect to a Redis server within a specified timeout period. If the connection succeeds, it returns the Redis client; otherwise, it handles errors or timeouts appropriately.
- **Why It’s Necessary**: This approach ensures that the application doesn’t hang indefinitely when trying to connect to Redis, making it more robust and reliable in production environments.


**Disk storage refers to a type of storage medium used to store data persistently on physical disks, such as hard drives (HDDs) or solid-state drives (SSDs). Unlike RAM (Random Access Memory), which is volatile and loses its data when the system is powered off, disk storage retains data even when the system is turned off.*

**Latency refers to the time delay between the moment a request is made and the moment a response is received.*


### On HTTP Request Header, Session ID and Redis Session Cache

- **HTTP Request Header**:

A request header in HTTP contains key-value pairs sent by a client to a server as part of an http request. The header provides information about the request, such as client details, request parameters, and metadata. These headers help the server understand how to process the request. Here are the common components of an HTTP request header:

1. **Request Method**: Not part of the header itself, but it defines the action to be performed (e.g., GET, POST, PUT, DELETE).

2. **Host**: Specifies the domain name or IP address of the server where the request is being sent.
   - Example: `Host: www.example.com`

3. **User-Agent**: Contains information about the client software, such as the browser or application making the request.
   - Example: `User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64)`

4. **Accept**: Specifies the content types the client can process, such as `text/html`, `application/json`, etc.
   - Example: `Accept: text/html,application/xhtml+xml,application/xml`

5. **Accept-Encoding**: Indicates the content encoding (e.g., gzip, deflate) that the client can understand.
   - Example: `Accept-Encoding: gzip, deflate`

6. **Accept-Language**: Specifies the language(s) the client prefers.
   - Example: `Accept-Language: en-US,en;q=0.9`

7. **Authorization**: Contains credentials for authenticating the request, such as a token or Basic Auth.
   - Example: `Authorization: Bearer <token>`

8. **Content-Type**: Indicates the MIME type of the body of the request (for POST, PUT requests).
   - Example: `Content-Type: application/json`

9. **Content-Length**: Specifies the length of the request body in bytes.
   - Example: `Content-Length: 348`

10. **Cookie**: Sends cookies from the client to the server.
    - Example: `Cookie: sessionId=abc123; theme=light`

11. **Referer**: Specifies the URL of the page that referred the client to the current resource.
    - Example: `Referer: https://www.example.com/page`

12. **Connection**: Controls whether the network connection stays open after the current transaction.
    - Example: `Connection: keep-alive`

13. **Cache-Control**: Directives for caching mechanisms in both requests and responses.
    - Example: `Cache-Control: no-cache`

14. **X-Requested-With**: Typically used in AJAX requests to identify the request as being made via JavaScript.
    - Example: `X-Requested-With: XMLHttpRequest`

15. **If-Modified-Since**: Allows conditional requests, only fetching the resource if it has been modified since the specified date.
    - Example: `If-Modified-Since: Wed, 21 Oct 2015 07:28:00 GMT`

#### Examples of Standard HTTP Request Headers

- **`GET /index.html HTTP/1.1`**
  ```
  Host: www.example.com
  User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36
  Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
  Accept-Language: en-US,en;q=0.5
  Accept-Encoding: gzip, deflate, br
  Connection: keep-alive
  ```

Understanding these components helps in debugging, optimizing performance, and ensuring proper communication between clients and servers.


- **Session ID**:

When you implement session management, the `Session-ID` header can be used to pass a session identifier from the client to the server, allowing the server to retrieve the session data associated with that ID (often stored in a database, cache like Redis, or in-memory).

#### Example of a Request Header with `Session-ID`:

```http
GET /api/resource HTTP/1.1
Host: www.example.com
User-Agent: Mozilla/5.0
Authorization: Bearer <token>
Session-ID: abc123xyz
```

In this example, the `Session-ID` header is included along with other headers like `User-Agent` and `Authorization`.

#### How `Session-ID` Works in Practice:

1. **User Logs In**: When a user logs in, the server creates a session and generates a `Session-ID`. This `Session-ID` is sent back to the client, usually as a cookie or in the response body.

2. **Client Stores `Session-ID`**: The client stores the `Session-ID` locally, typically in a cookie or local storage.

3. **Subsequent Requests**: For subsequent requests, the client includes the `Session-ID` in the request header. The server uses this ID to look up the session data (e.g., user info, shopping cart contents) and process the request accordingly.

4. **Session Validation**: The server validates the `Session-ID`, ensuring it corresponds to an active session. If valid, the request proceeds; if not, the server may reject the request or prompt the user to log in again.

In my code, however, JWT, rather than Redis, is handling the each user's session as seen here: 

```go
RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //session expires in 24 hours.
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ecommerce-backend",
			Subject:   "user",
			Audience:  jwt.ClaimStrings{"ecommerce-frontend"},
			ID:        "unique",
		},
```



## On Middlewares

Middleware is like a filter that processes requests before they reach the main logic of your application. In web applications, middleware functions can check if a user is logged in, if they have certain permissions, or if the request data is valid. In the context of my code, middleware is applied for RBAC. RBAC which is short for **Role-Based Access Control** restricts access based on roles assigned to users.

### The Middleware Logic from the code:

**Outer Function (RoleMiddleware):****

- Signature: `func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler`
- Role: This is a middleware generator function. It takes a variadic parameter allowedRoles which specifies the roles that are allowed to access the next handler. It returns a function that takes an http.Handler and returns another `http.Handler`.

**First Return Function:****

- Signature: `func(next http.Handler) http.Handler`
- Role: This function is the actual middleware. It takes the next handler in the chain (`next http.Handler`) as an argument and returns an `http.Handler`. This allows the middleware to wrap the next handler and perform additional processing before or after calling the next handler.

**Second Return Function:****

- Signature: `http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { ... })`
- Role: This function is the core of the middleware logic. It is an `http.HandlerFunc` that takes an `http.ResponseWriter` and an `http.Request` as arguments. Inside this function, you can implement the middleware logic, such as checking user roles, logging, or modifying the request/response. After performing the middleware logic, it can call `next.ServeHTTP(w, r)` to pass control to the next handler in the chain.

















glossary.md
architecture.md
technology.md
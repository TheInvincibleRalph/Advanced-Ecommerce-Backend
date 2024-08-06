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





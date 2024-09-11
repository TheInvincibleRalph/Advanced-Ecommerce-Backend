# Advanced E-commerce Backend

This is an advanced e-commerce backend built with the **Gorilla Mux** router and **PostgreSQL** as the database. The project integrates several modern technologies to ensure high performance, scalability, and a seamless user experience.

## Key Technologies

- **Firebase**: Used for push notifications to keep users updated on order statuses and more.
- **Mailgun**: Facilitates sending automated emails such as order confirmations and promotional emails.
- **Stripe**: Handles secure payment processing with support for various payment methods.
- **Redis**: Provides caching for faster data retrieval and reduced database load.
- **GORM**: An ORM (Object-Relational Mapper) for database interaction with PostgreSQL.
- **JWT (JSON Web Tokens)**: Used for secure authentication and authorization of users.

## Features

- **Pagination, Filtering, and Search**: Efficiently browse products using pagination, filtering by categories, and searching for specific items.

## Auth Routes

- `POST` `/api/v1/signup` (signup)
- `POST` `/api/v1/login` (user login)


## User Routes

- `POST` `/api/v1/users` (add user)
- `GET` `/api/v1/users` (get users)
- `GET` `/api/v1/users/{id}` (get user by id)
- `PUT` `/api/v1/users/{id}` (update user)
- `DELETE` `/api/v1/users/{id}` (delete user)


## Environment Variables

- `DB_HOST`
- `DB_USER`
- `DB_NAME`
- `DB_PORT`
- `DB_PASSWORD`
- `MAILGUN_DOMAIN`
- `JWT_SECRET_KEY`
- `MAILGUN_API_KEY`
- `STRIPE_SECRET_KEY`
- `MAILGUN_PUBLIC_API_KEY`

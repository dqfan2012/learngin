# Learning Gin

This is my journey into learning the gin web framework and giving it my own mini "Laravel" twist.

## Translating Laravel Commands

I use the `Makefile` to attempt to replicate Laravel console commands.

### Database Migrations

**Running Migrations**

`Laravel`

```bash
php artisan migrate
```

`Mine`

```bash
make migrate-up
```

**Rolling Migrations Back**

`Laravel`

```bash
php artisan migrate:rollback
```

`Mine`

```bash
make migrate-down
```

### Creating Models

`Laravel`

```bash
php artisan make:model <name>
```

`Mine`

My model creation isn't nearly as robust as Laravel's model creation.

```bash
make model <name>
```

### Seeding the Database

`Laravel`

```bash
php artisan db:seed
```

`Mine`

```bash
make seed
```

### Getting a basic development build going

**Running the basic server**

```bash
make server
```

You can make sure the server is running using the following command:

```bash
curl http://localhost:8080/status
```

The response should be:

```bash
{
    "status": "ok"
}
```

Be sure to pipe into `jq` to pretty print JSON to the console.

**A little more robust**

This will tear down all previous migrations. Run fresh migrations. Seed the database. Start the server.

```bash
make setup-dev
```

If you use the basic user model provided in this repository, you can see the first "dummy" user here:

```bash
curl http://localhost:8080/user | jq
```

You'll get an output similar to the following:

```bash
{
    "id": 1,
    "first_name": "Hillary",
    "last_name": "Kutch",
    "email": "vqmseDt@YjEowla.info",
    "password": "password",
    "role": "publisher",
    "remember_token": "VLxRVIaSbF",
    "created_at_utc": "2024-06-04T03:30:15.648846Z",
    "updated_at_utc": "2024-06-04T03:30:15.648846Z",
    "created_at_local": "2024-06-03T23:30:15.648846-04:00",
    "updated_at_local": "2024-06-03T23:30:15.648846-04:00",
    "email_verified_at": null
}
```

## Database

This currently features migrations for Postgresql.

## Contributors

I am not accepting contributions to this repo at this time.

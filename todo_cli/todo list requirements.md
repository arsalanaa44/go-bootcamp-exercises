
//entities

//application{technical like ports, configes, ...}

(business) entities

    categories
        properties
            title
            color
        behavior
            create a new category
            edit a category
            list categories with progress status

    task
        properties
            title
            dueDate
            category
            isDone
        behavior
            create new task
            list user today task (single responsibility from SOLID)
            list user tasks by date
            change task status(done/undone)
            edit task

    user
        properties
            id
            email
            password

        behavior
            register a user
            log in user

    userStory
        user should be registered
        can create category
        add a new task
        can see list of categories with progress status
        can see today`s task
        can see tasks by date
        make done/undone a task
        edit a task
        edit a category

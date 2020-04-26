## task-app

## schema

### Users

- display_name
- login_name
- password_digest
- avatar_url

### Project_users

- user_id(FK)
- project_id(FK)
- role

### Projects

- title
- description

### Tasks

- project_id(FK)
- name
- body(追加する)
- due_on
- status
  - wating
  - doing
  - done
  - canceled

### Notes

- message
- image_url
- project_id(FK)
- user_id(FK)

### Momes

- task_id(FK)
- body

### Project_tags

- project_id(FK)
- tag_id(FK)

### Tags

- name
- color

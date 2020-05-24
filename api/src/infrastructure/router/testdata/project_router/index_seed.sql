INSERT INTO users
  (
    id,
    display_name,
    login_name,
    password_digest,
    avatar_url,
    created_at,
    updated_at
  )
VALUES
  (
    1,
    "mike",
    "mike",
    "112233",
    "http://localhost:3000/avatar",
    "2020-04-25 15:16:08",
    "2020-04-25 15:16:08"
  );

INSERT INTO projects
  (
    id,
    title,
    description
  )
VALUES
  (
    1,
    "project1",
    "project1 description"
  ),
  (
    2,
    "project2",
    "project2 description"
  );

INSERT INTO project_users
  (
    id,
    user_id,
    project_id,
    role
  )
VALUES
  (
    1,
    1,
    1,
    "admin"
  ),
  (
    2,
    1,
    2,
    "write"
  );


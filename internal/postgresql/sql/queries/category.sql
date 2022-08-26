-- name: ListAllCategories :many
select *
from category;

-- name: CreateCategory :exec
insert into category (name, icon)
values (@name, @icon);

-- name: DeleteCategory :exec
delete
from category
where id = @id;

-- name: UpdateCategory :exec
update category
set name       = @name,
    icon       = @icon,
    updated_at = @updated_at
where id = @id;
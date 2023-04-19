create table orders (
    id varchar primary key,
    book_id varchar,
    user_id varchar,
    returned boolean not null,
    ordered_date date not null,
    returned_date date,
    foreign key (book_id) REFERENCES books (id) on delete cascade,
    foreign key (user_id) REFERENCES users (id) on delete cascade
)
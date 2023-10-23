-- create tables --

create table if not exists public.users (
    id serial not null primary key,
    name text,
    email text,
    password text,
    position text
);

create table if not exists public.timesheets (
    id serial not null primary key,
    user_id integer not null,
    date_week date,
    date_submitted date,
    data jsonb,

    constraint FK_timesheet_user foreign key(user_id)
        references public.users(id)
);

-- insert default data -- 

insert into public.users (id, name, email, password, position) values
    (1, 'Test User', 'test@email.com', 'testPassword', 'Intern');
    
insert into public.timesheets (user_id, date_week, date_submitted, data) values
    (1, current_timestamp, current_timestamp, '{"monday": 8}');
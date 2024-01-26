-- 生徒テーブルの作成
create table students (
    student_id int auto_increment primary key,
    name varchar(255) not null unique,
    age int not null
);

-- 生徒テーブルの初期データの挿入
insert into students (name, age) values ('山田太郎', 20), ('佐藤花子', 18), ('鈴木一郎', 19);
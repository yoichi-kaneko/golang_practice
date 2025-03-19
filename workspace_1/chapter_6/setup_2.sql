-- DB検証用のテーブル作成クエリ。リレーション関連
drop table if exists posts cascade;
drop table if exists comments cascade;

create table posts (
  id serial primary key,
  content text,
  author varchar(255)
);

create table comments (
  id serial primary key,
  content text,
  author varchar(255),
  post_id integer references posts(id)
);
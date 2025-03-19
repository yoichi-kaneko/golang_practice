-- DB検証用のテーブル作成クエリ
create table posts (
  id serial primary key,
  content text,
  author varchar(255)
);

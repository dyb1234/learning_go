insert into students(name, age, grade) VALUES ("张三",20,"三年级");

select * from students where age > 18;

update students set grade = "四年级" where name = "张三";

delete from students where age < 15;

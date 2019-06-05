1. 查询” 01 “课程比” 02 “课程成绩高的学生的信息及课程分数
select a.Sid, c.Sname, a.score as ascore, b.score as bscore from (select score, Sid from SC where Cid = 01) a left join
(select score, Sid from SC where Cid = 02) b on a.score > b.score left join Student c on a.Sid = c.Sid where a.Sid = b.Sid

2. 查询平均成绩大于等于 60 分的同学的学生编号和学生姓名和平均成绩
select s.Sname, b.Sid, b.avgScore from Student s right join (select Sid, avg(score) as avgScore from SC group by Sid having avg(score) > 60) b on s.Sid = b.Sid

3. 查询在 SC 表存在成绩的学生信息
select * from Student where Sid in (select distinct Sid from SC)

4. 查询所有同学的学生编号、学生姓名、选课总数、所有课程的总成绩(没成绩的显示为 null )
select s.Sid, s.Sname, b.* from Student s left join (select Sid, count(Cid), Sum(score) from SC group by Sid) b on s.Sid = b.Sid

5. 查询「李」姓老师的数量
select count(1) from Teacher where Tname like '李%'

6. 查询学过「张三」老师授课的同学的信息
select * from Student where Sid in (select Sid from SC sc, Teacher t, Course c where sc.Cid = c.Cid and t.Tid = c.Tid and t.Tname = '张三')

7. 查询没有学全所有课程的同学的信息
select * from Student where Sid in (select Sid from SC group by Sid having count(Cid) = (select Count(Cid) from Course))

https://blog.csdn.net/paul0127/article/details/82529216
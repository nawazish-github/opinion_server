create table question (
	id INT GENERATED ALWAYS AS identity  primary key,
	date date not null,
	question_prompt varchar(300) not null
);

select * from question q2 ;

create table option (
	id INT GENERATED ALWAYS AS identity primary key,
	qid int not null,
	option_prompt varchar(300) not null,
	
	CONSTRAINT fk_option_qid_question_id
      FOREIGN KEY(qid) 
	  REFERENCES question(id)
);

select * from "option" o ;

create table opinion(
	id INT GENERATED ALWAYS AS identity primary key,
	qid int not null,
	oid int references option (id) not null,
	ip_addr varchar(20),
	datetime timestamptz not null,
	
	CONSTRAINT fk_response_qid_question_id
      FOREIGN KEY(qid) 
	  REFERENCES question(id)
);

select * from opinion

select current_date;
insert into question (date, question_prompt) values('2021-09-18', 'Do you like Bangalore food?');
insert into option (qid, option_prompt) values (3, 'Chalega');
insert into option (qid, option_prompt) values (3, 'Not at all');
insert into option (qid, option_prompt) values (3, 'Absolutely!');

select * from question;
select * from option;
select * from opinion r ;
commit;

drop table opinion;
drop table option;
drop table question;

show timezone;

select q.qid, q.question_prompt from question q where qid = 2;

select q.qid, q.question_prompt, o.oid, o.option_prompt from question q inner join option o on  q.qid = o.qid where q.date = 'some date';


delete  from question 
delete from option;
delete from opinion;
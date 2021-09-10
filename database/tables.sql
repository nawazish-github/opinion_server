create table question (
	qid INT GENERATED ALWAYS AS identity  primary key,
	date varchar(20) not null,
	question_prompt varchar(300) not null
);

create table option (
	oid INT GENERATED ALWAYS AS identity primary key,
	qid int not null,
	option_prompt varchar(300) not null,
	
	CONSTRAINT fk_option_question
      FOREIGN KEY(qid) 
	  REFERENCES question(qid)
);

create table opinion(
	id INT GENERATED ALWAYS AS identity primary key,
	qid int not null,
	oid int references option (oid) not null,
	ip_addr varchar(20),
	
	CONSTRAINT fk_response_question
      FOREIGN KEY(qid) 
	  REFERENCES question(qid)

);


insert into question (date, question_prompt) values('some date', 'how is the fun?');
insert into option (qid, option_prompt) values (1, 'drizzling today');
insert into option (qid, option_prompt) values (1, 'sunny today');

select * from question;
select * from option;
select * from opinion r ;
commit;

drop table response;
drop table option;
drop table question;
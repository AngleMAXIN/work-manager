CREATE SCHEMA `work_manager` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin ;
    
drop table if exists wm_work;                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             create table wm_user
create table wm_user
(
	id int auto_increment
		primary key,
	real_name char(5) default '' not null,
	user_id bigint unsigned default 0 not null,
	u_type tinyint unsigned default 0 not null,
	create_time datetime not null,
	password varchar(255) default '' not null,
	level smallint(5) unsigned default 0 not null,
	major varchar(15) default '' not null,
	grade_id int unsigned default 0 not null,
	class_id tinyint unsigned default 0 not null,
	constraint wm_user_user_id_uindex
		unique (user_id)
);

drop table if exists wm_work;
create table wm_work (
    id int auto_increment primary key,
    creator char(5) default "" not null,
    title varchar(25) default "" not null,
    upload_time datetime null,
    creator_id int default 0 not null,
    homework_id int default 0 not null,
    comment varchar(50) default "" not null,
    score tinyint unsigned default 0 not null,
    grade_id int default 0 not null
);
                                                                                                                                                                                                   
drop table if exists wm_homework;
create table wm_homework (
    id int auto_increment primary key,
    title varchar(15) default "" not null,
    creator_id int default 0 not null,
    creator char(5) default "" not null,
    create_time datetime not null,
    start_time datetime not null,
    end_time datetime not null,
    belong_class int default 0 not null
);       

drop table if exists wm_grade;
create table wm_grade (
    id int auto_increment primary key,
    grade_id int unsigned default 0 not null,
    level smallint(5) unsigned default 0 not null,
	major varchar(15) default '' not null,
);
create table if not exists test (
    id int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    testName varchar(255),
    createdAt timestamp,
    updatedAt timestamp
);


create table if not exists test_question (
    id int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    questionName varchar(255),
    testID int,
    createdAt timestamp,
    updatedAt timestamp
);

create table if not exists  test_option (
    id int unsigned auto_increment primary key,
    optionName varchar(255),
    trueOption boolean,
    testID int,
    questionID int,
    createdAt timestamp,
    updatedAt timestamp
);

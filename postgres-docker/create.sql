CREATE USER curriculum;
CREATE DATABASE curriculum;
GRANT ALL PRIVILEGES ON DATABASE curriculum TO curriculum;
\c curriculum
CREATE TABLE subjects (
  id SERIAL PRIMARY KEY,
  subject_id VARCHAR(255),
  subject VARCHAR(255),
  key_stage INT,
  purpose_of_study VARCHAR(255)
);

CREATE TABLE subject_content (
  id SERIAL PRIMARY KEY,
  subject_id VARCHAR(255),
  subject_content VARCHAR(255),
  subject_content_id VARCHAR(255)
);

CREATE TABLE objectives (
  id SERIAL PRIMARY KEY,
  subject_content_id VARCHAR(255),
  objective_id VARCHAR(255),
  objective VARCHAR(255)
);

CREATE TABLE sub_objectives (
  id SERIAL PRIMARY KEY,
  objective_id VARCHAR(255),
  sub_objective_id VARCHAR(255),
  sub_objective VARCHAR(255)
);

COPY subjects(subject_id, subject, key_stage, purpose_of_study)
FROM '/Subjects.csv'
DELIMITER '|';

COPY subject_content(subject_id, subject_content, subject_content_id)
FROM '/SubjectContent.csv'
DELIMITER '|';

COPY objectives(subject_content_id, objective_id, objective)
FROM '/Objectives.csv'
DELIMITER '|';

COPY sub_objectives(objective_id,sub_objective_id,sub_objective)
FROM '/SubObjectives.csv'
DELIMITER '|';


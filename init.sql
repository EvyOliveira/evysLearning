CREATE TABLE Courses (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100),
  description VARCHAR(150)
);

CREATE TABLE Classes (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(50),
  resume VARCHAR(150),
  text VARCHAR(200),
  course_id BIGSERIAL,
  FOREIGN KEY (course_id) REFERENCES Courses(id),
	CONSTRAINT fk_class_course_id FOREIGN KEY (course_id) REFERENCES Courses(id)
);

CREATE TABLE Exercises (
  id BIGSERIAL PRIMARY KEY,
  question VARCHAR(70),
  answer VARCHAR(150),
  correctAnswer VARCHAR(150),
  class_id BIGSERIAL,
  course_id BIGSERIAL,
	FOREIGN KEY (class_id) REFERENCES Classes(id),
	FOREIGN KEY (course_id) REFERENCES Courses(id)
);

INSERT INTO Courses(name, description) VALUES ('Analysis and systems development', 'Focus on developing systems design and development skills.');
INSERT INTO Courses(name, description) VALUES ('Information systems', 'This course concerns the available information systems and their complexities.');
INSERT INTO Classes(title, resume, text) VALUES ('Linear Programming', 'Algorithms and methods for the Linear Programming approach.', 'Elective subject');
INSERT INTO Classes(title, resume, text) VALUES ('Software Engineering', 'Approaches, project management methodologies, systems diagramming.', 'Mandatory subject');
INSERT INTO Exercises(question, answer, correctAnswer)VALUES ('What is the best programming language?', 'Python', 'It depends on the context, developers familiarity, software applicability and other factors.');
INSERT INTO Exercises(question, answer, correctAnswer)VALUES ('What does SQL mean?', 'It is a language for manipulating data through a database that makes use of relational algebra.', 'Popular query language used in various types of applications.');

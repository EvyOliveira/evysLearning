CREATE TABLE "Courses" (
  "ID" PK,
  "Name" string,
  "Description" string
);

CREATE TABLE "Classes" (
  "ID" PK,
  "Title" string,
  "Resume" string,
  "Text" string,
  "Course" FK,
  CONSTRAINT "FK_Classes.ID"
    FOREIGN KEY ("ID")
      REFERENCES "Courses"("ID")
);

CREATE TABLE "Exercises" (
  "ID" PK,
  "Question" string,
  "Answers" string,
  "CorrectAnswer" string,
  "Class" FK,
  "Course" FK,
  CONSTRAINT "FK_Exercises.Question"
    FOREIGN KEY ("Question")
      REFERENCES "Courses"("Description")
);
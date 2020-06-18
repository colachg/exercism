class School:
    """
    class School
    """

    def __init__(self):
        self.students = {}

    def add_student(self, name, grade):
        names = []
        if grade in self.students:
            names += self.students[grade]
            names.append(name)
            self.students[grade] = names
        else:
            names.append(name)
            self.students[grade] = names

    def roster(self):
        grades = [i for i in self.students.keys()]
        students = []

        for i in sorted(grades):
            students += sorted(self.students[i])
        return students

    def grade(self, grade_number):
        if self.students:
            return sorted(self.students[grade_number])
        return []

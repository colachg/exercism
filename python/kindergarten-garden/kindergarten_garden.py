class Garden:
    origin_students = ["Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph",
                       "Kincaid", "Larry"]

    def __init__(self, diagram, students=[]):
        # parse students
        if students:
            self.students = students
        else:
            self.students = self.origin_students
        self.students.sort()
        # parse diagram
        self.rows = diagram.split("\n")

    def plants(self, name):
        plants_maps = {'R': 'Radishes', 'C': 'Clover', 'G': 'Grass', 'V': 'Violets'}
        index = self.students.index(name)
        print(self.students, name, index)
        plants = self.rows[0][2 * index:2 * index+2] + self.rows[1][2 * index :2 * index+2]
        results = []
        for f in plants:
            results.append(plants_maps[f])
        return results

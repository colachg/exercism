class Matrix:
    def __init__(self, matrix_string):
        self._rows = matrix_string.split('\n')

    def row(self, index):
        row = self._rows[index-1].split()
        return [int(i) for i in list(row)]

    def column(self, index):
        columns = []
        for i in self._rows:
            columns.append(i.split()[index-1])
        return [int(i) for i in columns]



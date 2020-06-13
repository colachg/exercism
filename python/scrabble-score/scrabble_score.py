def score(word):
    p1 = ("A", "E", "I", "O", "U", "L", "N", "R", "S", "T")
    p2 = ("D", "G")
    p3 = ("B", "C", "M", "P")
    p4 = ("F", "H", "V", "W", "Y")
    p5 = ("K",)
    p8 = ("J", "X")
    p10 = ("Q", "Z")

    points = 0
    for i in word:
        if i.upper() in p1:
            points += 1
        if i.upper() in p2:
            points += 2
        if i.upper() in p3:
            points += 3
        if i.upper() in p4:
            points += 4
        if i.upper() in p5:
            points += 5
        if i.upper() in p8:
            points += 8
        if i.upper() in p10:
            points += 10
    return points

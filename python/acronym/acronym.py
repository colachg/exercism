def abbreviate(words):
    acronym = ""
    for i in words.replace("-", " ").split():
        acronym += i.strip("_")[0].upper()
    return acronym

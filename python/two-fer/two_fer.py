def two_fer(*name):
    if not name:
        return "One for you, one for me."
    return 'One for %s, one for me.' % name

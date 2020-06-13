def is_isogram(string):
    # replace ' ' and '-'
    string = string.lower().replace(' ', '').replace('-', '')
    return len(string) == len(set(string))

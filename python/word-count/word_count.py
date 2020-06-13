from string import punctuation


def count_words(sentence):
    # need to remove ' as it is one chars we specifically want to keep inside a word
    for char in punctuation.replace("'", ""):
        sentence = sentence.lower().replace(char, " ")

    split_words = [item.strip("'") for item in sentence.split()]

    return {item: split_words.count(item) for item in split_words}
The input may be classified into three things. One is append to beginning represented with ‘[‘, another is append to end represented with ‘]’ and the other types any other character except above mentioned. The other character are appended to the next position of last inserted character.

So task is move all characters following ‘[‘ to beginning until ‘]’ is found or end of string is reached. Similar logic applies to ‘]’ character. Also it is a recursive definition, keep applying this logic for all the following third brackets.

Input:
This_is_a_[Beiju]_text
[[]][][]Happy_Birthday_to_Tsinghua_University

Output:
BeijuThis_is_a__text
Happy_Birthday_to_Tsinghua_University
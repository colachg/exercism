class Luhn:
    def __init__(self, card_num):
        card_num = card_num.replace(" ", "")
        if len(card_num) > 1 and card_num.isdecimal():
            self.card_num = [int(i) for i in card_num.replace(" ", "")]
        else:
            self.card_num = []

    def valid(self):
        if self.card_num:
            result = sum(self.card_num[-1::-2])
            for i in self.card_num[-2::-2]:
                i = i * 2 - 9 if i * 2 > 9 else i * 2
                result += i
            if result % 10 == 0:
                return True
        return False
